package xgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/nickxb/pkg/xruntime"
	"net/http/httputil"
)



func Recovery(fn func(ctx *gin.Context, err interface{})) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				fn(ctx, err)
				stack := xruntime.Stack(2)
				httprequest, _ := httputil.DumpRequest(ctx.Request, false)
				logMsg := fmt.Sprintf("[Recovery] panic recovered:\n%s\n%s\n%s", string(httprequest), err, stack)
				logrus.Error(logMsg)
				fmt.Println(logMsg)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

