/*
@File   : yamlconfig.go
@Author : pan
@Time   : 2023-06-09 17:00:18
*/
package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Id       string     `yaml:"id,omitempty"`
	Info     Info       `yaml:"info"`
	Requests []Requests `yaml:"requests,omitempty"`
	Http     Http       `yaml:"Http"`
}

type Requests struct {
	Operator Operator               `yaml:",inline"`
	Raw      []string               `yaml:"raw,omitempty"`
	Payloads map[string]interface{} `yaml:"payloads,omitempty"`
}

type Info struct {
	Name   string `yaml:"name,omitempty"`
	Author string `yaml:"author,omitempty"`
}

type Operator struct {
	Matchers    string     `yaml:"matchers,omitempty"`
	MatchersCon []Matchers `yaml:"matchers-con,omitempty"`
}

type Matchers struct {
	Type      string   `yaml:"type,omitempty"`
	Condition string   `yaml:"condition,omitempty"`
	Status    []int    `yaml:"status,omitempty"`
	Part      string   `yaml:"part,omitempty"`
	Regex     []string `ymal:"regex,omitempty"`
}

type Http struct {
	Port string `yaml:"port,omitempty"`
	Host string `yaml:"host,omitempty"`
}

func ParseYaml(filepath string) (*YamlConfig, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	config := &YamlConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func WriteYaml(filename string, config *YamlConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	data, err := ParseYaml("../../config.yaml")
	if err != nil {
		fmt.Printf("parse yaml err:%v", err)
	}
	fmt.Println(data)
	err = WriteYaml("../../writeconfig.yaml", data)
	if err != nil {
		fmt.Printf("write yaml err:%v", err)
	}
}
