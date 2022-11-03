package config

type App struct {
	Host             string // 服务地址
	Port             int    // 服务端口
	Runmode          string // 运行模式
	ProjectName      string // 项目名称
	Debug            bool   // 是否开启debug
	LoggerFolderPath string // 日志路径
	Sale             string // 加密盐
	UploadFolderPath string // 上传文件路径
	SqlSaveToFile    bool   // 是否将sql保存为文件
	SqlLogPath       string // sql日志路径
	IsInstall        bool   // 是否是第一次安装
	InstallFileName  string // 安装文件名
}
