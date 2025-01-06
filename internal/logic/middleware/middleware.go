package middleware

import (
	"TechBE/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Response(r *ghttp.Request) {
	r.Middleware.Next()
	res := r.GetHandlerResponse()
	r.Response.WriteJson(res)
}
