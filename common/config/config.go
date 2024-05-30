package config

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}
