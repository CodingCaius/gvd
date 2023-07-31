//core文件夹用于初始化连接


//该文件用于读取配置文件setting.yaml
//并将读取到的信息存储到config文件夹的Config结构体中

//逻辑上也相当于初始化Config
package core

import (
	"gvd_server/config"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const yamlPath = "setting.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		logrus.Fatalln("read yaml err: ", err.Error())
	}

	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		logrus.Fatalln("解析 yaml err: ", err.Error())
	}

	return c
}
