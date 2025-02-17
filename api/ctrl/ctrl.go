// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ctrl

import (
	"context"

	"TechBE/api/ctrl/v1"
)

type ICtrlV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
}
