package config

var (
	Keys *Config
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"database"`

	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}
