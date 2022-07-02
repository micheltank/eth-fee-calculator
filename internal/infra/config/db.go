package config

import "fmt"

type DbConfig struct {
	User     string `env:"DB_USER"`
	Port     string `env:"DB_PORT"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
}

func (d DbConfig) BuildURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", d.User, d.Password, d.Host, d.Port, d.Name)
}
