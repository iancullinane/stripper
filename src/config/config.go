package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Input  string
	Output string
}

func New() *config {
	return &config{
		Input:  "./test_file/gauis.png",
		Output: "./output/%d.png",
	}
}

func (c *config) LoadConfig() config {
	// filename := os.Args[1]
	var config config
	source, err := ioutil.ReadFile("../config/base.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func (c *config) GetOutputDir() string {
	return c.Output
}

func (c *config) GetInputFile() string {
	return c.Input
}
