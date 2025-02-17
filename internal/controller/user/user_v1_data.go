package user

import (
	"context"

	v1 "TechBE/api/user/v1"
	"TechBE/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Data(ctx context.Context, req *v1.DataReq) (res *v1.DataRes, err error) {
	res = &v1.DataRes{}
	userdata := service.User().GetData(ctx, req.Name, req.Password)
	if userdata == nil {
		g.RequestFromCtx(ctx).Response.Writeln("no userdata of ", req.Name)
		return
	}
	res.Userdata = *userdata
	// g.RequestFromCtx(ctx).Response.WriteJson(res)
	return
}
