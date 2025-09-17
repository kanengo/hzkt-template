package service

import (
	"context"
	v1 "solabar-server/api/idgenerator/v1"
	biz "solabar-server/internal/idgenerator/biz"
)

type IdGeneratorIntService struct {
	v1.UnimplementedIdGeneratorIntServer
	idGenUsecase *biz.IdGeneratorUsecase
}

func NewIdGeneratorIntService(idGenUsecase *biz.IdGeneratorUsecase) *IdGeneratorIntService {
	return &IdGeneratorIntService{idGenUsecase: idGenUsecase}
}

func (s *IdGeneratorIntService) GenerateId(ctx context.Context, req *v1.GenerateIdRequest) (*v1.GenerateIdResponse, error) {
	id, err := s.idGenUsecase.GenerateStrId(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.GenerateIdResponse{
		Id: &v1.GenerateIdResponse_StrId{StrId: id},
	}, nil
}
