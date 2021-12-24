package entity

type Config struct {
	Server *Server `yaml:"server"`
}

type Server struct {
	Port     int    `yaml:"port"`
	HexoRoot string `yaml:"hexo_root"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
