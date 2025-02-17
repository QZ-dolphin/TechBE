package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"TechBE/internal/controller/ctrl"
	"TechBE/internal/controller/system"
	"TechBE/internal/controller/user"
	"TechBE/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(service.Middleware().CORS)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().Response)
				group.Bind(
					user.NewV1(),
					ctrl.NewV1(),
					system.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
