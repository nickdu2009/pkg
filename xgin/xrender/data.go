package xrender

import "github.com/gin-gonic/gin/render"

type Data struct {
	HttpCode
	render.Data
}
