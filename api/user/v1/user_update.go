package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta `path:"/userupdate" method:"post"`
	Name   string `p:"username"`
	Data   string `p:"userdata"`
}

type UpdateRes struct{}
