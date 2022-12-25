package data

import (
	"backend-GuardRails/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type scanRepo struct {
	data *Data
	log  *log.Helper
}

// NewScanRepo .
func NewScanRepo(data *Data, logger log.Logger) biz.ScanRepo {
	return &scanRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
