package data

import (
	"backend-GuardRails/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type repositoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewRepositoryRepo(data *Data, logger log.Logger) biz.RepositoryRepo {
	return &repositoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
