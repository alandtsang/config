package config

import (
	"fmt"
	"testing"
)

type Data struct {
	A string `yaml:"a"`
	B B      `yaml:"b"`
}

type B struct {
	C int   `yaml:"c"`
	D []int `yaml:"d"`
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name       string
		configPath string
		wantErr    error
	}{
		{
			name:       "BindStruct",
			configPath: "test.yml",
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := NewConfig(tt.configPath)
			if err != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var data Data
			if err = cfg.Bind(&data); err != nil {
				t.Errorf("NewConfig() bind error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Printf("%+v\n", data)
		})
	}
}
