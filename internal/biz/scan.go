package biz

import "github.com/go-kratos/kratos/v2/log"

type ScanRepo interface {
}

type ScanUsecase struct {
	repo ScanRepo
	log  *log.Helper
}

// NewScanUsecase new a Scan usecase.
func NewScanUsecase(repo ScanRepo, logger log.Logger) *ScanUsecase {
	return &ScanUsecase{repo: repo, log: log.NewHelper(logger)}
}
