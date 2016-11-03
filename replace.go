package parser

import (
	"fmt"
	"strings"
)

// Replace
func (p *Parser) Replace(word string, lineNum int) (string, error) {
	v := strings.Replace(word, p.Delimeter, "", -1)
	if k, ok := p.Variables[v]; ok {
		return k.(string), nil
	} else {
		return "", fmt.Errorf("Could not found variable %q, on line %d please verify the variables file", v, lineNum)
	}
}
