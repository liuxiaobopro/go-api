package qiniu

import (
	"github.com/liuxiaobopro/go-api/global"

	"github.com/liuxiaobopro/go-lib/upload/qiniu"
)

func InitQiniu() {
	// 初始化七牛云
	qiniu.QiniuConfig.AccessKey = global.Conf.Upload.Qiniu.AccessKey
	qiniu.QiniuConfig.SecretKey = global.Conf.Upload.Qiniu.SecretKey
	qiniu.QiniuConfig.Bucket = global.Conf.Upload.Qiniu.Bucket
	qiniu.QiniuConfig.Domain = global.Conf.Upload.Qiniu.Domain
	qiniu.QiniuConfig.FilePath = global.Conf.Upload.Qiniu.FilePath
	qiniu.QiniuConfig.IsDelLocal = global.Conf.Upload.Qiniu.IsDelLocal
}
