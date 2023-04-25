package configs

import (
	"os"
	"project/src/pkg/utils"

	"gopkg.in/yaml.v3"
)

type Config struct {
	API apiConfig `yaml:"api"`
	DB  dbConfig  `yaml:"db"`
}

type apiConfig struct {
	Domain string `yaml:"domain"`
	Port   int    `yaml:"port"`
}

type dbConfig struct {
	Driver string `yaml:"driver"`
	Conn   string `yaml:"conn"`
}

// * Loads the config and return a struct
// * with all the avaliable configs
func LoadConfig() (*Config, error) {
	utils := utils.New()
	path, err := utils.GetFilePath(&[]string{"services", "broker", "configs", "base.yaml"})
	if err != nil {
		return nil, err
	}
	f, err := os.Open(*path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
