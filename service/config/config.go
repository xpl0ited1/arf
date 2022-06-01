package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mongodb",
			Host:     "127.0.0.1",
			Port:     27017,
			Username: "",
			Password: "",
			Name:     "activereconframework",
			Charset:  "utf8",
		},
	}
}
