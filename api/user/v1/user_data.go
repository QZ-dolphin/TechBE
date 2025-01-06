package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DataReq struct {
	g.Meta `path:"/userdata" method:"get"`
	Name   string `p:"username"`
}

type DataRes struct {
	Userdata
}

type Userdata struct {
	Data  string      `json:"stackdata"`
	CTime *gtime.Time `json:"created_at"`
	UTime *gtime.Time `json:"updated_at"`
}
