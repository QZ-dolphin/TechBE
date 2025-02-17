package v1

import "github.com/gogf/gf/v2/frame/g"

type EmailCodeGetReq struct {
	g.Meta  `path:"/emailcode_get" method:"post"`
	ToEmail string `p:"toemail"`
}

type EmailCodeGetRes struct {
	Send bool `p:"send"`
}
