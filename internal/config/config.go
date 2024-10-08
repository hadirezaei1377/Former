package config

import (
	"os"
)

type Config struct {
	PublicDir  string
	PrivateDir string
}

func LoadConfig() *Config {
	return &Config{
		PublicDir:  os.Getenv("PUBLIC_DIR"),
		PrivateDir: os.Getenv("PRIVATE_DIR"),
	}
}
