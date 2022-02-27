package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config structure definition.
type Config struct {
	// store file descriptor.
	file *os.File

	// yaml decoder.
	decoder *yaml.Decoder
}

// NewConfig returns a new Config struct.
func NewConfig(configPath string) (*Config, error) {
	if err := validateConfigPath(configPath); err != nil {
		return nil, err
	}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	config := &Config{
		file:    file,
		decoder: yaml.NewDecoder(file),
	}

	return config, nil
}

// Bind binds config data to dst structure.
func (cfg *Config) Bind(dst interface{}) error {
	defer cfg.file.Close()

	if err := cfg.decoder.Decode(dst); err != nil {
		return err
	}
	return nil
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
