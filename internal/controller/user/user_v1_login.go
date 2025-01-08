package user

import (
	"context"

	v1 "TechBE/api/user/v1"
	"TechBE/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	res.Exits, res.Pass, res.Dul = service.User().Login(ctx, req.Username, req.Email, req.Password)
	return
}
