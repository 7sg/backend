package service

import (
	v1 "backend-GuardRails/api/scan/v1"
	"backend-GuardRails/internal/biz"
	"context"
)

type ScanService struct {
	v1.UnimplementedScanServer

	uc *biz.ScanUsecase
}

// NewScanService new a scan service.
func NewScanService(uc *biz.ScanUsecase) *ScanService {
	return &ScanService{uc: uc}
}

func (s *ScanService) ScanRepository(context.Context, *v1.ScanRepositoryRequest) (*v1.ScanRepositoryResponse, error) {
	return nil, nil
}

func (s *ScanService) GetScanRepositoryResult(context.Context, *v1.GetScanRepositoryResultRequest) (*v1.GetScanRepositoryResultResponse, error) {
	return nil, nil
}
