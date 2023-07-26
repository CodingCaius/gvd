//core文件夹用于初始化连接

package core

import (
	"gvd_server/config"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalln("read yaml err: ", err.Error())
	}

	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		log.Fatalln("解析 yaml err: ", err.Error())
	}

	return c
}
