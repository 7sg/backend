package biz

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/go-kratos/kratos/v2/log"
)

type GitCloneUsecase struct {
	repo                 RepositoryRepo
	scanRepo             ScanRepo
	gitCloneKafkaRepo    GitCloneKafkaRepo
	fileContentKafkaRepo FileContentKafkaRepo
	log                  *log.Helper
}

// NewGitCloneUseCase new a GitClone usecase.
func NewGitCloneUseCase(
	repo RepositoryRepo,
	scanRepo ScanRepo,
	gitCloneKafkaRepo GitCloneKafkaRepo,
	fileContentKafkaRepo FileContentKafkaRepo,
	logger log.Logger,
) *GitCloneUsecase {
	gitCloneUseCase := &GitCloneUsecase{
		repo:                 repo,
		scanRepo:             scanRepo,
		gitCloneKafkaRepo:    gitCloneKafkaRepo,
		fileContentKafkaRepo: fileContentKafkaRepo,
		log:                  log.NewHelper(logger),
	}
	go gitCloneUseCase.ConsumeMessage()

	return gitCloneUseCase
}

func (g *GitCloneUsecase) ConsumeMessage() {
	ctx := context.Background()
	g.log.Infof("starting to consume git clone message")
	for {
		msg, err := g.gitCloneKafkaRepo.GetMessage(ctx)
		if err != nil {
			g.log.Errorf("error while getting git clone message from kafka: %v", err)

			continue
		}
		gitCloneEvent := &GitCloneEvent{}
		err = json.Unmarshal(msg.Value, gitCloneEvent)
		if err != nil {
			g.log.Errorf("error while unmarshalling git clone event: %v", err)

			continue
		}
		g.log.Infof("git clone event received: %+v", gitCloneEvent)
		repository, err := g.repo.GetRepository(ctx, gitCloneEvent.RepositoryID)
		if err != nil {
			g.log.Errorf("error while getting repository: %v", err)

			continue
		}
		startTime := time.Now()
		files, err := g.startCloning(ctx, repository.Link, gitCloneEvent.ScanID)
		if err != nil {
			g.log.Errorf("error while cloning repository: %v", err)
			err = g.scanRepo.UpdateScanStartAndEndTimeWithFailure(
				ctx,
				gitCloneEvent.ScanID,
				startTime,
				time.Now(),
			)
			if err == nil {
				_ = g.gitCloneKafkaRepo.CommitMessage(ctx, msg)
			}

			continue
		}
		if len(files) == 0 {
			g.log.Infof("no files to scan in repository: %s", repository.Link)
			err = g.scanRepo.UpdateScanStartAndEndTimeWithSuccess(
				ctx,
				gitCloneEvent.ScanID,
				startTime,
				time.Now(),
			)
			if err == nil {
				_ = g.gitCloneKafkaRepo.CommitMessage(ctx, msg)
			}

			continue
		}
		err = g.scanRepo.UpdateScanStartTimeAndTotalFiles(
			ctx,
			gitCloneEvent.ScanID,
			startTime,
			uint32(len(files)),
		)
		if err != nil {
			g.log.Errorf("error while updating scan start time: %v", err)

			continue
		}
		err = g.fileContentKafkaRepo.PublishFileContent(ctx, files)
		if err != nil {
			g.log.Errorf("error while publishing file content: %v", err)

			continue
		}
		_ = g.gitCloneKafkaRepo.CommitMessage(ctx, msg)
	}
}

func (g *GitCloneUsecase) startCloning(
	ctx context.Context,
	link string,
	scanID uint64,
) ([]*FileContent, error) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: link,
	})
	if err != nil {
		return nil, err
	}
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}
	var files []*FileContent
	_ = tree.Files().ForEach(func(f *object.File) error {
		isBinaryFile, err := f.IsBinary()

		if err == nil && !isBinaryFile {
			lines, err := f.Lines()
			if err == nil {
				fileContent := &FileContent{Path: f.Name, Lines: lines, ScanID: scanID}
				files = append(files, fileContent)
			} else {
				g.log.Errorf("error while getting lines from file: %v", err)
			}
		} else {
			g.log.Infof("file is binary: %s, err %v", f.Name, err)
		}

		return nil
	})

	return files, nil
}
