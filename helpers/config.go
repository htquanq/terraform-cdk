package helpers

import (
	. "aws-eks/structs"

	"gopkg.in/yaml.v2"
)

func Config(yamlFile []byte) AppConfig {
	var config AppConfig
	yaml.Unmarshal(yamlFile, &config)

	return config
}
