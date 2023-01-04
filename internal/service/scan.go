package service

import (
	"context"
	"fmt"

	v1 "backend-GuardRails/api/scan/v1"
	"backend-GuardRails/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScanService struct {
	v1.UnimplementedScanServer
	uc *biz.ScanUsecase
	gc *biz.GitCloneUsecase
	fc *biz.FileContentUsecase
}

// NewScanService new a scan service.
func NewScanService(
	uc *biz.ScanUsecase,
	gc *biz.GitCloneUsecase,
	fc *biz.FileContentUsecase,
) *ScanService {
	return &ScanService{uc: uc, gc: gc, fc: fc}
}

func (s *ScanService) ScanRepository(
	ctx context.Context,
	scanRequest *v1.ScanRepositoryRequest,
) (*v1.ScanRepositoryResponse, error) {
	scan, err := s.uc.ScanRepository(ctx, scanRequest.RepositoryId)
	if err != nil {
		return nil, fmt.Errorf("failed to scan repository, try again later")
	}

	return &v1.ScanRepositoryResponse{
		ResultId:     scan.ID,
		ScanStatus:   v1.ScanStatus_Queued,
		EnqueuedTime: timestamppb.New(*scan.EnqueuedTime),
	}, nil
}

func (s *ScanService) GetScanRepositoryResult(
	ctx context.Context,
	getScanRequest *v1.GetScanRepositoryResultRequest,
) (*v1.GetScanRepositoryResultResponse, error) {
	scan, err := s.uc.GetScan(ctx, getScanRequest.ResultId)
	if err != nil {
		return nil, fmt.Errorf("failed to get scan result, try again later")
	}

	scanResult := &v1.GetScanRepositoryResultResponse{
		ResultId:     scan.ID,
		ScanStatus:   v1.ScanStatus(v1.ScanStatus_value[scan.Status]),
		RepositoryId: scan.RepositoryID,
		EnqueuedTime: timestamppb.New(*scan.EnqueuedTime),
		Findings:     convertToFindingsProto(scan.Findings),
	}
	if scan.StartTime != nil {
		scanResult.StartTime = timestamppb.New(*scan.StartTime)
	}
	if scan.EndTime != nil {
		scanResult.FinishTime = timestamppb.New(*scan.EndTime)
	}

	return scanResult, nil
}

func convertToFindingsProto(findings []*biz.Finding) []*v1.Finding {
	findingsProto := make([]*v1.Finding, 0)
	for _, finding := range findings {
		findingsProto = append(findingsProto, &v1.Finding{
			Type:   finding.Type,
			RuleId: finding.RuleID,
			Location: &v1.Location{
				Path: finding.Location.Path,
				Positions: &v1.Positions{
					Begin: &v1.Begin{Line: finding.Location.Positions.Begin.Line},
				},
			},
			Metadata: &v1.Metadata{
				Description: finding.Metadata.Description,
				Severity:    finding.Metadata.Severity,
			},
		})
	}

	return findingsProto
}
