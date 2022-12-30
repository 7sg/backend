package service

import (
	v1 "backend-GuardRails/api/scan/v1"
	"backend-GuardRails/internal/biz"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScanService struct {
	v1.UnimplementedScanServer

	uc *biz.ScanUsecase
	gc *biz.GitCloneUseCase
}

// NewScanService new a scan service.
func NewScanService(uc *biz.ScanUsecase, gc *biz.GitCloneUseCase) *ScanService {
	return &ScanService{uc: uc, gc: gc}
}

func (s *ScanService) ScanRepository(ctx context.Context, scanRequest *v1.ScanRepositoryRequest) (*v1.ScanRepositoryResponse, error) {
	scan, err := s.uc.ScanRepository(ctx, scanRequest.RepositoryId)
	if err != nil {
		return nil, err
	}
	return &v1.ScanRepositoryResponse{ResultId: scan.Id, ScanStatus: v1.ScanStatus_Queued, EnqueuedTime: timestamppb.New(scan.EnqueuedTime)}, nil
}

func (s *ScanService) GetScanRepositoryResult(context.Context, *v1.GetScanRepositoryResultRequest) (*v1.GetScanRepositoryResultResponse, error) {
	return nil, nil
}
