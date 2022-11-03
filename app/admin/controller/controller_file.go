package controller

import (
	"github.com/liuxiaobopro/go-api/global"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/go-lib/ecode"
	ginl "github.com/liuxiaobopro/go-lib/gin"
	"github.com/liuxiaobopro/go-lib/upload/local"
	"github.com/liuxiaobopro/go-lib/upload/qiniu"
)

type FileControllerType struct {
	ginl.Handler
}

var FileController = new(FileControllerType)

// UploadQiniu 文件上传到七牛云
func (th *FileControllerType) UploadQiniu(c *gin.Context) {
	file, fileHeader, err := th.GetFormFile(c, "file")
	if err != nil {
		th.SendError(c, err, ecode.ERROR)
		return
	}
	url, err := qiniu.Upload(file, fileHeader)
	if err != nil {
		th.SendError(c, err, ecode.ERROR_SERVER)
		return
	}
	th.SendSucc(c, url, nil)
}

// UploadLocal 文件上传到本地
func (th *FileControllerType) UploadLocal(c *gin.Context) {
	file, fileHeader, _ := th.GetFormFile(c, "file")
	filepath, filename := local.Upload(file, fileHeader, global.Conf.App.UploadFolderPath)
	var r = map[string]string{
		"filepath": filepath,
		"filename": filename,
	}
	th.SendSucc(c, r, nil)
}
