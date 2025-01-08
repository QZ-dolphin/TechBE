// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "TechBE/api/user/v1"
	"context"
)

type (
	IUser interface {
		GetData(ctx context.Context, username string) (userdata *v1.Userdata)
		SaveData(ctx context.Context, username string, userdata string)
		Login(ctx context.Context, username string, email string, password string) (exits bool, pass bool, dup bool)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
