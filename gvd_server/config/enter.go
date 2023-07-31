//该文件夹用于存放解析到的配置

package config

type Config struct {
	System System `yaml:"system"`
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	Jwt Jwt `yaml:"jwt"`
}