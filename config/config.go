package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func Parse(cfg interface{}, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrFileNotExist
	}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	return nil
}

func MustParse(cfg interface{}, path string) {
	if err := Parse(cfg, path); err != nil {
		panic(err)
	}
}
