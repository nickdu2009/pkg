package xrender

import "github.com/gin-gonic/gin/render"

type Redirect struct {
	HttpCode
	render.Redirect
}
