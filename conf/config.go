package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Mysql       Mysql
	Application Application
	Spark       Spark
}

type Mysql struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

type Application struct {
	Host string
	Port string
}

type Spark struct {
	HostUrl   string `yaml:"hostUrl"`
	Appid     string `yaml:"appid"`
	ApiSecret string `yaml:"apiSecret"`
	ApiKey    string `yaml:"apiKey"`
}

var Conf *Config

func init() {
	ReadConfig()
}

func ReadConfig() {
	file, err := os.ReadFile(".\\conf\\config.yaml")
	if err != nil {
		log.Println("config.ReadConfig()读取文件失败：", err)
		return
	}
	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Println("config.ReadConfig()解析yaml文件失败：", err)
		return
	}
}
