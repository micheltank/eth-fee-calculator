package config

type DbConfig struct {
	User     string `env:"DB_USER"`
	Port     string `env:"DB_PORT"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
}
