package xrender

import "github.com/gin-gonic/gin/render"

type Render interface {
	Code() int
	render.Render
}

type HttpCode struct {
	Code_ int
}

func (hc HttpCode) Code() int {
	return hc.Code_
}
