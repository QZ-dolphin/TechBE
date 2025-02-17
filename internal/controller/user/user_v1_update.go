package user

import (
	"context"

	v1 "TechBE/api/user/v1"
	"TechBE/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	service.User().SaveData(ctx, req.Name, req.Password, req.Data)
	return
}
