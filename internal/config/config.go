package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	HTTP     HTTPConfig `toml:"httpapi"`
	DBConfig DBConfig   `toml:"postgres"`
}

type HTTPConfig struct {
	Address string `toml:"address"`
}

type DBConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	SSLMode  string `toml:"sslmode"`
}

func New(cfgPath string) (*Config, error) {
	var cfg Config

	if _, err := toml.DecodeFile(cfgPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Database,
		c.SSLMode,
	)
}
