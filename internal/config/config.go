package config

type Config struct {
	Host     string `json:"host"`
	Post     string `json:"port"`
	DataBase string `json:"db"`
	User     string `json:"user"`
	Password string `json:"pswd"`
}
