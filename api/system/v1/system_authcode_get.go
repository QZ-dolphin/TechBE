package v1

// system_authcode_get.go
import "github.com/gogf/gf/v2/frame/g"

type AuthCodeGetReq struct {
	g.Meta `path:"/authcode_get" method:"get"`
}
type AuthCodeGetRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Img    string `json:"img"`
}
