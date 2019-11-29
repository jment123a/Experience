package handler

import (
	"github.com/gin-gonic/gin"
)

//Video 视频信息
type Video struct {
	id   int
	name string
}

//GetVideoList 获取视频信息
func GetVideoList(c *gin.Context) {
	c.JSON(200, "视频信息序列表")
}
