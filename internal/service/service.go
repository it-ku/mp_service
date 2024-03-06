package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/it-ku/mp-service/api/mp/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMiniAppService)

type MiniAppService struct {
	v1.UnimplementedMiniAppServiceServer

	log *log.Helper
}

func NewMiniAppService(logger log.Logger) *MiniAppService {
	return &MiniAppService{
		log: log.NewHelper(log.With(logger, "module", "service/MiniAppService")),
	}
}
