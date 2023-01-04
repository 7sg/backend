package server

import (
	repositoryV1 "backend-GuardRails/api/repository/v1"
	scanV1 "backend-GuardRails/api/scan/v1"
	"backend-GuardRails/internal/conf"
	"backend-GuardRails/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	repository *service.RepositoryService,
	scan *service.ScanService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	repositoryV1.RegisterRepositoryHTTPServer(srv, repository)
	scanV1.RegisterScanHTTPServer(srv, scan)

	return srv
}
