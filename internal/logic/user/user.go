package user

import (
	v1 "TechBE/api/user/v1"
	"TechBE/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) GetData(ctx context.Context, username string) (userdata *v1.Userdata) {
	md := g.Model("userdata")
	md.Where("username", username).Scan(&userdata)
	// utility.Clog(userdata)
	if userdata == nil {
		// utility.Clog("\n Empty")
	}
	return
}

func (s *sUser) SaveData(ctx context.Context, username string, userdata string) {
	data := g.Map{
		"stackdata": userdata,
	}
	md := g.Model("userdata")
	md.Where("username", username).Update(data)
}

func (s *sUser) Login(ctx context.Context, username string, email string, password string) (exits bool, pass bool, dup bool) {
	md := g.Model("userdata").Safe()
	exits, _ = md.Where("username", username).Exist()
	clientIP := g.RequestFromCtx(ctx).GetClientIp()
	if email == "" {
		if !exits {
			return
		} else {
			ps, _ := md.Fields("password").Where("username", username).Value()
			if ps.String() == password {
				ip, _ := md.Fields("client_ip").Where("username", username).Value()
				if ip.String() != clientIP {
					md.Where("username", username).Update(g.Map{"client_ip": clientIP})
				}
				pass = true
				return
			}
		}
	} else {
		// ip检查账号数量，大于3拒绝
		account_nums, _ := md.Where("client_ip", clientIP).Count()
		if account_nums >= 2 {
			// dup 默认为false，用户超上限
			return
		}
		// 新用户注册
		exits = true
		pass = true
		dup = true
		data := g.Map{
			"username":  username,
			"email":     email,
			"password":  password,
			"client_ip": clientIP,
		}
		md.Insert(data)
	}
	return
}
