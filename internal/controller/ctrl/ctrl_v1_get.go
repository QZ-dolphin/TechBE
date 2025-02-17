package ctrl

import (
	"context"

	v1 "TechBE/api/ctrl/v1"
	"TechBE/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	res = new(v1.GetRes)
	if req.Password == "123dqz" {
		res.UserData = service.Ctrl().GetUserData()
	} else {
		g.RequestFromCtx(ctx).Response.Writeln("Wrong user login!")
	}
	return
}
