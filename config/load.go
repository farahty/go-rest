package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func newConfig(configPath string) (*Config, error) {

	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	println("configration loaded successfully.")

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func parseFlags() (string, error) {

	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yml", "path to config file")

	flag.Parse()

	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func LoadConfig() error {
	var err error
	var str string
	str, err = parseFlags()

	if err != nil {
		return err
	}

	Keys, err = newConfig(str)

	if err != nil {
		return err
	}

	return nil
}
