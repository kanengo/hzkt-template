package biz

import (
	"context"
	"log/slog"
	idgenv1 "solabar-server/api/idgenerator/v1"
)

type UserUsecase struct {
	log      *slog.Logger
	idGenCli idgenv1.IdGeneratorIntClient
}

func NewUserUsecase(log *slog.Logger, idGenCli idgenv1.IdGeneratorIntClient) *UserUsecase {
	return &UserUsecase{log: log, idGenCli: idGenCli}
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, userId int64) (string, error) {
	uc.log.Info("GetUserInfo", "uid", userId)
	resp, err := uc.idGenCli.GenerateId(ctx, &idgenv1.GenerateIdRequest{
		BizType: 0,
	})
	if err != nil {
		return "", err
	}
	return "leeka " + resp.GetStrId(), nil
}
