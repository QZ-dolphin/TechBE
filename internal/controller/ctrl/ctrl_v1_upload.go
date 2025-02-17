package ctrl

import (
	"context"

	v1 "TechBE/api/ctrl/v1"
	"TechBE/internal/service"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	res = &v1.UploadRes{}
	res.Filepath = service.Ctrl().UploadFile(ctx, req.File, req.Username, req.Password, req.Fileytype)
	return
}
