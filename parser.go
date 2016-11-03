package parser

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Parser struct
type Parser struct {
	template  *os.File
	Variables map[string]interface{}
	Delimeter string
}

// New returns a Parser
func New(t, v string) (*Parser, error) {
	// read variables first
	ymlFile, err := ioutil.ReadFile(v)
	if err != nil {
		return nil, err
	}
	var a map[string]interface{}
	if err := yaml.Unmarshal(ymlFile, &a); err != nil {
		return nil, err
	}

	// open template file
	template, err := os.Open(t)
	if err != nil {
		return nil, fmt.Errorf("Could not open template %q: %s", t, err)
	}
	return &Parser{
		template,
		a,
		"$",
	}, nil
}

// CloseTemplate closes os.File
func (p *Parser) CloseTemplate() {
	p.template.Close()
}
