package system

import (
	"context"

	v1 "TechBE/api/system/v1"
	"TechBE/internal/service"
	"TechBE/utility"
)

func (c *ControllerV1) AuthCodeGet(ctx context.Context, req *v1.AuthCodeGetReq) (res *v1.AuthCodeGetRes, err error) {
	idKeyC, base64stringC, err := service.System().GetVerifyImgString(ctx)
	if err != nil {
		utility.Clog("\n获取验证码失败", err)
		return nil, err
	}
	res = new(v1.AuthCodeGetRes)
	res.Key = idKeyC
	res.Img = base64stringC
	return
}
