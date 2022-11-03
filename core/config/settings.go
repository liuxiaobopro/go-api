package config

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/liuxiaobopro/go-api/config"
	"github.com/liuxiaobopro/go-api/global"

	"github.com/ghodss/yaml"
	"github.com/liuxiaobopro/go-lib/console"
)

func InitConfig(fs embed.FS) {
	parseFlag()
	loadYaml(fs)
}

// 解析命令行参数
func parseFlag() {
	// flag.StringVar(&config.Conf.App.Host, "host", "127.0.0.1", "Please enter host")
	flag.IntVar(&config.Conf.App.Port, "port", 8081, "Please enter port")
	flag.StringVar(&config.Conf.App.Runmode, "runmode", "", "product|dev|test")
	// flag.BoolVar(&config.Conf.App.Debug, "debug", false, "Please enter debug")

	flag.Parse()
}

// 加载yaml配置文件
func loadYaml(fs embed.FS) {
	runmode := config.Conf.App.Runmode
	data, err := fs.ReadFile(fmt.Sprintf("config/yaml/%s.yaml", runmode))
	if err != nil {
		console.Console.Error(fmt.Sprintf("快去 config/yaml 下创建 %s.yaml", runmode), err.Error())
	}

	//#region 映射config
	if err := yaml.Unmarshal(data, config.Conf); err != nil {
		console.Console.Error("解析配置文件失败", err.Error())
	}
	global.Conf = config.Conf
	//#endregion

	//#region 判断是否是第一次加载
	if _, err := os.Stat(config.Conf.App.InstallFileName); !os.IsNotExist(err) {
		global.Conf.App.IsInstall = true
	}
	//#endregion
}

// // 加载yaml配置文件
// func loadYaml() {
// 	runmode := config.Conf.App.Runmode
// 	var yamlFile string
// 	if runmode == "" {
// 		// 读取默认文件
// 		yamlFile = "settings.yaml"
// 	} else {
// 		// 读取指定文件
// 		yamlFile = fmt.Sprintf("settings.%s.yaml", runmode)
// 	}
// 	content, err := ioutil.ReadFile(yamlFile)
// 	if err != nil {
// 		console.Console.Error(err.Error(), fmt.Sprintf("请在项目根目录创建配置文件: %s", yamlFile))
// 	}

// 	console.Console.Info(fmt.Sprintf("读取配置文件：%s", yamlFile))

// 	err = yaml.Unmarshal(content, config.Conf)
// 	if err != nil {
// 		console.Console.Error(err.Error(), "解析配置文件失败")
// 	}
// }
