package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MiniAppRepo interface {
	ListMiniApp(ctx context.Context, current, pageSize uint32)
	CreateMiniApp()
	GetMiniApp()
	UpdateMiniApp()
	DeleteMiniApp()
}

type MiniAppUseCase struct {
	repo MiniAppRepo
	log  *log.Helper
}

func NewMiniAppUseCase(repo MiniAppRepo, logger log.Logger) *MiniAppUseCase {
	return &MiniAppUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "biz/MiniAppUseCase"))}
}
