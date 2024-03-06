package service

import (
	"context"
	v1 "github.com/it-ku/mp-service/api/mp/v1"
)

func (s *MiniAppService) LoginByCode(ctx context.Context, req *v1.CodeLoginReq) (*v1.LoginReply, error) {
	// 参数校验
	return nil, nil
}
func (s *MiniAppService) GetPhoneNumber(ctx context.Context, req *v1.GetPhoneNumberReq) (*v1.GetPhoneNumberReply, error) {
	return nil, nil
}
