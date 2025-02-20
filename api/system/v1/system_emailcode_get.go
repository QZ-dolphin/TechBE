package v1

import "github.com/gogf/gf/v2/frame/g"

type EmailCodeGetReq struct {
	g.Meta  `path:"/emailcode_get" method:"post"`
	ToEmail string `p:"toemail"`
}

type EmailCodeGetRes struct {
	TTL   int64 `json:"ttl"` // 验证码有效时间，单位秒
	Send  bool  `json:"send"`
	Exits bool  `json:"exits"`
}
