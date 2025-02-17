// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "TechBE/api/ctrl/v1"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICtrl interface {
		GetUserData() (userdata []v1.BEData)
		UploadFile(ctx context.Context, file *ghttp.UploadFile, username string, password string, filetype string) (filepath string)
	}
)

var (
	localCtrl ICtrl
)

func Ctrl() ICtrl {
	if localCtrl == nil {
		panic("implement not found for interface ICtrl, forgot register?")
	}
	return localCtrl
}

func RegisterCtrl(i ICtrl) {
	localCtrl = i
}
