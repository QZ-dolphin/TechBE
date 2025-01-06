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
