package conf

import "github.com/jasontconnell/conf"

type Config struct {
	ConnectionString string `json:"connectionString"`
	Database         string `json:"database"`
}

func LoadConfig(filename string) (Config, error) {
	cfg := Config{}

	err := conf.LoadConfig(filename, &cfg)

	return cfg, err
}
