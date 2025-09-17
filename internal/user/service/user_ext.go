package service

import (
	"context"
	v1 "solabar-server/api/user/v1"
	"solabar-server/internal/user/biz"
)

type UserExtService struct {
	v1.UnimplementedUserExtServer

	uc *biz.UserUsecase
}

func NewUserExtService(uc *biz.UserUsecase) *UserExtService {
	return &UserExtService{uc: uc}
}

func (s *UserExtService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (*v1.GetUserInfoRes, error) {
	name, err := s.uc.GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserInfoRes{
		Name: name,
	}, nil
}
