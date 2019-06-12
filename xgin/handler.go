package xgin

import (
	"io"
	"net/http"

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

func DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) xrender.Render {
	r := xrender.Reader{}
	r.Code_ = code
	r.Headers = extraHeaders
	r.ContentType = contentType
	r.ContentLength = contentLength
	r.Reader.Reader = reader

	return r
}

func Redirect(code int, location string, req *http.Request) xrender.Render {
	r := xrender.Redirect{}
	r.Code_ = -1
	r.Redirect.Code = code
	r.Location = location
	r.Request = req
	return r
}

func Data(code int, contentType string, data []byte) xrender.Render {
	r := xrender.Data{}
	r.Code_ = code
	r.ContentType = contentType
	r.Data.Data = data
	return r
}
