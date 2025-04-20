package models

type HttpServer struct {
	Addr string `yaml:"address"`
}

type MongoDB struct {
	Uri      string `yaml:"uri"`
	Username string `yaml:"mongo_user"`
	Password string `yaml:"mongo_pass"`
}

type Config struct {
	Env        string `yaml:"env"`
	HttpServer `yaml:"http_server"`
	MongoDB    `yaml:"mongodb"`
}
