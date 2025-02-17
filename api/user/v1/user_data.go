package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DataReq struct {
	g.Meta   `path:"/userdata" method:"post"`
	Name     string `p:"username"`
	Password string `p:"password"`
}

type DataRes struct {
	Userdata
}

type Userdata struct {
	Stackdata   string      `json:"stackdata"`
	Avatar_path string      `json:"avatar_path"`
	Created_at  *gtime.Time `json:"created_at"`
	Updated_at  *gtime.Time `json:"updated_at"`
}
