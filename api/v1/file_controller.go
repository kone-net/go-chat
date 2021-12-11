package v1

import (
	"io/ioutil"
	"net/http"
	"strings"

	"chat-room/config"
	"chat-room/internal/service"
	"chat-room/pkg/common/response"
	"chat-room/pkg/global/log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 前端通过文件名称获取文件流，显示文件
func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	log.Logger.Info(fileName)
	data, _ := ioutil.ReadFile(config.GetConfig().StaticPath.FilePath + fileName)
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

	log.Logger.Info("file", log.Any("file name", config.GetConfig().StaticPath.FilePath+newFileName))
	log.Logger.Info("userUuid", log.Any("userUuid name", userUuid))

	c.SaveUploadedFile(file, config.GetConfig().StaticPath.FilePath+newFileName)
	err := service.UserService.ModifyUserAvatar(newFileName, userUuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.SuccessMsg(newFileName))
}
