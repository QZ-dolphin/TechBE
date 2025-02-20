// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISystem interface {
		GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error)
		VerifyCaptcha(ctx context.Context, idKey string, captcha string) bool
		EmailCodeSend(ctx context.Context, to string, code string) (exits bool, send bool, ttl int64)
		EmailCodeVerify(ctx context.Context, to string, code string) bool
	}
)

var (
	localSystem ISystem
)

func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

func RegisterSystem(i ISystem) {
	localSystem = i
}
