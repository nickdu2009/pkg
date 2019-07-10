package xrender

import "github.com/gin-gonic/gin/render"

type Redirect struct {
	render.Redirect
	Code_ int
}

func (r Redirect) Code() int {
	return r.Code_
}
