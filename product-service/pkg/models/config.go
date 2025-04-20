package models

type HttpServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env        string `yaml:"env"`
	HttpServer `yaml:"http_server"`
}
