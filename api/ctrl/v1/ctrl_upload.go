package v1

// ctrl_upload.go
import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta    `path:"/upload" method:"post" mime:"multipart/form-data"`
	Username  string            `p:"username"`
	Password  string            `p:"password"`
	Fileytype string            `p:"filetype"`
	File      *ghttp.UploadFile `p:"ufile" type:"file"`
}
type UploadRes struct {
	Filepath string `json:"filepath"`
}
