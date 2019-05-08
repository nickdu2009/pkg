package xrender

import "github.com/gin-gonic/gin/render"

type JSON struct {
	HttpCode
	render.JSON
}
