package config

// Конфиг соединения с БД
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DataBase string `json:"db"`
	User     string `json:"user"`
	Password string `json:"pswd"`
}
