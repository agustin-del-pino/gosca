package gosca

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"gosca/utils"
)

type Dir struct {
	Name  string   `json:"name" yaml:"name"`
	Files []string `json:"files" yaml:"files"`
	Dirs  []Dir    `json:"dirs" yaml:"dirs"`
}

type Scaffolding struct {
	Scaffolding Dir `json:"scaffolding" yaml:"scaffolding"`
}

func (s *Scaffolding) FromJsonBytes(b []byte) error {
	if jerr := json.Unmarshal(b, &s); jerr != nil {
		return utils.WrapError(jerr, "the Scaffolding cannot be parse from json bytes")
	}
	return nil
}

func (s *Scaffolding) FromYamlBytes(b []byte) error {
	if yerr := yaml.Unmarshal(b, &s); yerr != nil {
		return utils.WrapError(yerr, "the Scaffolding cannot be parse from yaml bytes")
	}
	return nil
}

func (s *Scaffolding) ToJson() ([]byte, error) {
	return json.MarshalIndent(&s, "", "  ")
}

func (s *Scaffolding) ToYaml() ([]byte, error) {
	return yaml.Marshal(&s)
}
