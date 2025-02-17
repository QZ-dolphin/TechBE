package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta     `path:"/userlogin" method:"post"`
	Username   string `p:"username"`
	Password   string `p:"password"`
	Email      string `p:"email"`
	VerifyKey  string `p:"verifyKey"`
	VerifyCode string `p:"verifyCode"`
}

type LoginRes struct {
	Verify bool `json:"verify"`
	Exits  bool `json:"exits"`
	Pass   bool `json:"pass"`
	Dul    bool `json:"dul"`
}
