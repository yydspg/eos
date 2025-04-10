package config

import "gopkg.in/yaml.v2"

type Parser interface {
	Parse(data []byte,out any) error 
}

type YAMLParser struct {}

func (y *YAMLParser) Parse(data []byte,out any) error {
	return yaml.Unmarshal(data,out)
}