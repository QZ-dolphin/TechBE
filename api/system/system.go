// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"TechBE/api/system/v1"
)

type ISystemV1 interface {
	AuthCodeGet(ctx context.Context, req *v1.AuthCodeGetReq) (res *v1.AuthCodeGetRes, err error)
	EmailCodeGet(ctx context.Context, req *v1.EmailCodeGetReq) (res *v1.EmailCodeGetRes, err error)
	EmailCodeVfy(ctx context.Context, req *v1.EmailCodeVfyReq) (res *v1.EmailCodeVfyRes, err error)
}
