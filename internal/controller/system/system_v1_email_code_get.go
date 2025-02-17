package system

import (
	"context"

	v1 "TechBE/api/system/v1"
	"TechBE/internal/service"
	"TechBE/utility"
)

func (c *ControllerV1) EmailCodeGet(ctx context.Context, req *v1.EmailCodeGetReq) (res *v1.EmailCodeGetRes, err error) {
	code := utility.GenerateVerificationCode(6)
	res = new(v1.EmailCodeGetRes)
	res.Send = service.System().EmailCodeSend(ctx, req.ToEmail, code)
	return
}
