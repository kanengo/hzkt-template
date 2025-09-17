package biz

import (
	"context"
	"log/slog"
	idgenv1 "solabar-server/api/idgenerator/v1"
)

type IdGeneratorUsecase struct {
	log *slog.Logger
}

func NewIdGeneratorUsecase(log *slog.Logger) *IdGeneratorUsecase {
	return &IdGeneratorUsecase{log: log}
}

func (uc *IdGeneratorUsecase) GenerateStrId(ctx context.Context, req *idgenv1.GenerateIdRequest) (string, error) {
	uc.log.Info("GenerateStrId", "biz_type", req.BizType)

	return "167", nil
}
