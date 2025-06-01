package rules

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type RuleConfig struct {
	InputRules []Rule `yaml:"input_rules"`
	Version    string
	Name       string
}

type Rule struct {
	Name           string   `yaml:"name"`
	Category       string   `yaml:"category"`
	Pattern        string   `yaml:"pattern"`
	Type           string   `yaml:"type"`
	Actions        []Action `yaml:"actions"`
	Accomplishment bool
}

// Create creates a new rule
func Create() []Rule {

	return readRules()
}

func readRules() []Rule {
	file_path := "termevent.yml"
	meta := RuleConfig{}
	open, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err)
	}
	data, err := io.ReadAll(open)
	if err != nil {
		fmt.Println(err)
	}

	yaml.Unmarshal(data, &meta)

	return meta.InputRules
}
