package system

import (
	"context"

	v1 "TechBE/api/system/v1"
	"TechBE/internal/service"
)

func (c *ControllerV1) EmailCodeVfy(ctx context.Context, req *v1.EmailCodeVfyReq) (res *v1.EmailCodeVfyRes, err error) {
	res = &v1.EmailCodeVfyRes{}
	res.Vfy = service.System().EmailCodeVerify(ctx, req.ToEmail, req.Code)
	return
}
