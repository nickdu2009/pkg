package xrender

import "github.com/gin-gonic/gin/render"

type Reader struct {
	HttpCode
	render.Reader
}
