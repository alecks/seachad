package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func Test(c *gin.Context) {
	_ = render.WriteMsgPack(c.Writer, map[interface{}]interface{}{
		"a": "b",
	})
}
