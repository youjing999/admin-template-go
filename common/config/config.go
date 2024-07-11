package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//总配置文件
type config struct {
	server        server        `yaml:"server"`
	Db            db            `yaml:"db"`
	redis         redis         `yaml:"redis"`
	log           log           `yaml:"log"`
	imageSettings imageSettings `yaml:"imageSettings"`
}

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

type redis struct {
	Address string `yaml:"address"`
	//Password string `yaml:"password"`
}

type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

// image图片上传设置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

var Config *config

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
