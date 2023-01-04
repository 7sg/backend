package biz

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type FileContent struct {
	ScanID uint64
	Path   string
	Lines  []string
}

type FileContentUsecase struct {
	scanRepo             ScanRepo
	fileContentKafkaRepo FileContentKafkaRepo
	log                  *log.Helper
}

func NewFileContentUseCase(
	scanRepo ScanRepo,
	fileContentKafkaRepo FileContentKafkaRepo,
	logger log.Logger,
) *FileContentUsecase {
	fileContentUsecase := &FileContentUsecase{
		scanRepo:             scanRepo,
		fileContentKafkaRepo: fileContentKafkaRepo,
		log:                  log.NewHelper(logger),
	}
	go fileContentUsecase.ConsumeMessage()

	return fileContentUsecase
}

func (f *FileContentUsecase) ConsumeMessage() {
	f.log.Info("starting to consume file content message")
	ctx := context.Background()
	for {
		message, err := f.fileContentKafkaRepo.GetMessage(ctx)
		if err != nil {
			continue
		}
		fileContent := &FileContent{}
		err = json.Unmarshal(message.Value, fileContent)
		if err != nil {
			f.log.Errorf("error while unmarshalling file content: %v", err)

			continue
		}
		secretLineNos := f.DetectSecrets(ctx, fileContent)
		findings := f.BuildFindings(fileContent.Path, secretLineNos)
		scanUpdate, err := f.scanRepo.UpdateScanFindings(ctx, fileContent.ScanID, findings)
		if err != nil {
			f.log.Errorf("error while updating scan findings: %v", err)

			continue
		}
		if scanUpdate.TotalFiles == scanUpdate.ScannedFiles {
			_ = f.scanRepo.UpdateScanWithSuccess(ctx, fileContent.ScanID, time.Now())
		}

		_ = f.fileContentKafkaRepo.CommitMessage(ctx, message)
	}
}

func (f *FileContentUsecase) DetectSecrets(ctx context.Context, content *FileContent) []uint32 {
	var linesWithPossibleSecrets []uint32
	for lineNo, line := range content.Lines {
		if strings.Contains(line, "public_key") || strings.Contains(line, "private_key") {
			linesWithPossibleSecrets = append(linesWithPossibleSecrets, uint32(lineNo+1))
		}
	}

	return linesWithPossibleSecrets
}

func (f *FileContentUsecase) BuildFindings(path string, secretLineNos []uint32) string {
	findings := make([]Finding, 0)
	for _, secretLineNo := range secretLineNos {
		findings = append(findings, buildFinding(path, secretLineNo))
	}
	bytes, err := json.Marshal(findings)
	if err != nil {
		f.log.Errorf("error while marshalling findings: %v", err)
	}

	return string(bytes)
}

func buildFinding(path string, secretLineNo uint32) Finding {
	return Finding{
		Type:   "sast",
		RuleID: "G404",
		Location: Location{
			Path: path,
			Positions: Positions{
				Begin: Begin{
					Line: secretLineNo,
				},
			},
		},
		Metadata: Metadata{
			Description: "protected secrets are found",
			Severity:    "HIGH",
		},
	}
}
