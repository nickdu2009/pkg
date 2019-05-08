package xrender

import "github.com/gin-gonic/gin/render"

type String struct {
	HttpCode
	render.String
}
