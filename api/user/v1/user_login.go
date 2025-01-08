package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/userlogin" method:"post"`
	Username string `p:"username"`
	Password string `p:"password"`
	Email    string `p:"email"`
}

type LoginRes struct {
	Exits bool `json:"exits"`
	Pass  bool `json:"pass"`
	Dul   bool `json:"dul"`
}
