package entity

type Config struct {
	Server *Server `yaml:"server"`
	AliConfig *AliConfig `yaml:"ali"`
	Mysql *Mysql `yaml:"mysql"`
}

type Server struct {
	Port     int    `yaml:"port"`
	HexoRoot string `yaml:"hexo_root"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AliConfig struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Bucket          string `yaml:"bucket"`
	Region          string `yaml:"region"`
	OssHost         string `yaml:"ossHost"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}