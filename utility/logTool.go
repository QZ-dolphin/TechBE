package utility

// utility/logTool.go
import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

func Clog(v ...interface{}) {
	glog.Skip(1).Line(true).Info(context.TODO(), v...)
}
