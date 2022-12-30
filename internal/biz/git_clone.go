package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
)

type GitCloneUseCase struct {
	repo                 RepositoryRepo
	gitCloneKafkaRepo    GitCloneKafkaRepo
	fileContentKafkaRepo FileContentKafkaRepo
	log                  *log.Helper
}

// NewGitCloneUseCase new a GitClone usecase.
func NewGitCloneUseCase(repo RepositoryRepo, gitCloneKafkaRepo GitCloneKafkaRepo, fileContentKafkaRepo FileContentKafkaRepo, logger log.Logger) *GitCloneUseCase {
	gitCloneUseCase := &GitCloneUseCase{repo: repo, gitCloneKafkaRepo: gitCloneKafkaRepo, fileContentKafkaRepo: fileContentKafkaRepo, log: log.NewHelper(logger)}
	go gitCloneUseCase.ConsumeMessage()
	return gitCloneUseCase
}

func (g *GitCloneUseCase) ConsumeMessage() {
	ctx := context.Background()
	g.log.Infof("starting to consume git clone message")
	for {
		g.log.Infof("waiting for git clone message")
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
		err = g.gitCloneKafkaRepo.CommitMessage(ctx, msg)
		if err != nil {
			g.log.Errorf("error while committing git clone message: %v", err)
			continue
		}
		g.log.Infof("git clone message committed")
	}
	g.log.Infof("git clone message consumer stopped")
}
