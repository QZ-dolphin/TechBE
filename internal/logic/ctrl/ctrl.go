package ctrl

import (
	v1 "TechBE/api/ctrl/v1"
	"TechBE/internal/service"
	"TechBE/utility"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sCtrl struct{}

func init() {
	service.RegisterCtrl(New())
}

func New() service.ICtrl {
	return &sCtrl{}
}

func (s *sCtrl) GetUserData() (userdata []v1.BEData) {
	md := g.Model("userdata").Safe()

	md.Fields("username, password, email, client_ip").Scan(&userdata)
	for i := range userdata {
		userdata[i].Password, _ = utility.Crypt.Decrypt(userdata[i].Password)
	}
	return
}

func (s *sCtrl) UploadFile(ctx context.Context, file *ghttp.UploadFile, username string, password string, filetype string) (filepath string) {
	md := g.Model("userdata").Safe()
	if !legal_op(md, username, password) {
		return
	}
	if filetype == "avatar" {
		avatar_path := "resources/public/upload/" + username + "/" + "avatar"
		utility.ClearDir(avatar_path)
		prefx, _ := md.Fields("password").Where("username", username).Value()
		file.Filename = prefx.String()[:5] + "_" + file.Filename
		_, err := file.Save(avatar_path)
		if err != nil {
			utility.Clog("\n文件没保存成功")
		}
		filepath = avatar_path + "/" + file.Filename
		md.Where("username", username).Update(g.Map{"avatar_path": filepath})
		// g.RequestFromCtx(ctx).Response.ServeFile(filepath)
	}
	return
}

func legal_op(md *gdb.Model, username string, password string) bool {
	exits, _ := md.Where("username", username).Exist()
	if !exits {
		return false
	}
	ps, _ := md.Fields("password").Where("username", username).Value()
	x_pass, _ := utility.Crypt.Decrypt(ps.String())
	if ps.String() == password || password == x_pass {
		return true
	}
	return false
}
