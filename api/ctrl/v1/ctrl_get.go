package v1

import "github.com/gogf/gf/v2/frame/g"

type GetReq struct {
	g.Meta   `path:"/dqzget" method:"get"`
	Password string `p:"password"`
}

type BEData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Client_IP string `json:"client_ip"`
}

type GetRes struct {
	UserData []BEData
}
