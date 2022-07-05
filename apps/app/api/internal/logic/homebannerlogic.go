package logic

import (
	"context"

	"wlight/lebron/apps/app/api/internal/svc"
	"wlight/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomebannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomebannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomebannerLogic {
	return &HomebannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomebannerLogic) Homebanner(req *types.RecommendRequest) (resp *types.RecommendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
