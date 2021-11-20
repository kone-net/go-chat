package v1

import (
	"io/ioutil"

	"chat-room/global/log"

	"github.com/gin-gonic/gin"
)

// 前端通过文件名称获取文件流，显示文件
func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	log.Info(fileName)
	data, _ := ioutil.ReadFile("static/img/" + fileName)
	c.Writer.Write(data)
}
