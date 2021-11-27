package v1

import (
	"io/ioutil"
	"net/http"
	"strings"

	"chat-room/global/log"
	"chat-room/common/response"
	"chat-room/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 前端通过文件名称获取文件流，显示文件
func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	log.Info(fileName)
	data, _ := ioutil.ReadFile("static/img/" + fileName)
	c.Writer.Write(data)
}

// 上传头像等文件
func SaveFile(c *gin.Context) {
	namePreffix := uuid.New().String()

	userUuid := c.PostForm("uuid")

	file, _ := c.FormFile("file")
	fileName := file.Filename
	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]

	newFileName := namePreffix + suffix

	log.Info("file", log.Any("file name", "static/img/"+newFileName))
	log.Info("userUuid", log.Any("userUuid name", userUuid))

	c.SaveUploadedFile(file, "static/img/"+newFileName)
	err := service.UserService.ModifyUserAvatar(newFileName, userUuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.SuccessMsg(newFileName))
}
