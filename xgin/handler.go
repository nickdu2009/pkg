package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/nickxb/pkg/xgin/xrender"
)

type DecoratorHandlerFunc func(*gin.Context) xrender.Render

func HandlerDecorator(fn DecoratorHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := fn(ctx)
		if r != nil {
			ctx.Render(r.Code(), r)
		}
	}
}

func JSON(code int, obj interface{}) xrender.Render {
	r := xrender.JSON{}
	r.Code_ = code
	r.Data = obj
	return r
}

func String(code int, formant string, values ...interface{}) xrender.Render {
	r := xrender.String{}
	r.Code_ = code
	r.Format = formant
	r.Data = values
	return r
}
