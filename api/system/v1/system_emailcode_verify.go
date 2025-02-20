// system_emailcode_verify.go
package v1

import "github.com/gogf/gf/v2/frame/g"

type EmailCodeVfyReq struct {
	g.Meta  `path:"/emailcode_vfy" method:"post"`
	ToEmail string `json:"toemail" v:"required|email"`
	Code    string `json:"code" v:"required"`
}

type EmailCodeVfyRes struct {
	Vfy bool `json:"vfy"`
}
