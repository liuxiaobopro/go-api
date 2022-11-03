package config

type Upload struct {
	Qiniu struct {
		AccessKey  string // 七牛云accessKey
		SecretKey  string // 七牛云secretKey
		Bucket     string // 七牛云存储空间
		Domain     string // 七牛云图片域名
		FilePath   string // 上传到服务器的文件路径
		IsDelLocal bool   // 上传之后是否删除本地文件
	}
}
