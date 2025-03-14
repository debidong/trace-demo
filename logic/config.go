package logic

import (
	"os"

	"gopkg.in/yaml.v2"
)

var C *config

type config struct {
	Server   map[string]server   `yaml:"server"`
	Exporter map[string]exporter `yaml:"exporter"`
}

type server struct {
	Addr string   `yaml:"addr"`
	Uri  []string `yaml:"uri"`
}

type exporter struct {
	Addr string `yaml:"addr"`
}

func MustLoadConfig(path ...string) *config {
	if len(path) != 1 {
		panic("one path is required")
	}

	yamlFile, err := os.ReadFile(path[0])
	if err != nil {
		panic(err)
	}

	var config config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	C = &config
	return C
}
