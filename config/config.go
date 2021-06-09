package config

type Config struct {
	DB   *DBConfig
	Port string
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
			Dialect:  "postgres",
			Host:     "pte-magic.cvaz8sslmgiw.ap-southeast-2.rds.amazonaws.com",
			Port:     5432,
			Username: "ptemagic",
			Password: "12345678",
			Name:     "go_tutorial",
			Charset:  "utf8",
		},
		Port: ":8080",
	}
}
