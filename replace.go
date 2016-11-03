package parser

import (
	"fmt"
	"strings"
)

// Replace return the variable matching a word
func (p *Parser) Replace(word string) (string, error) {
	v := strings.Replace(word, p.Delimeter, "", -1)
	k, ok := p.Variables[v]
	if ok {
		return k.(string), nil
	}
	return "", fmt.Errorf("Variable not found")
}

// FindMatch clean the placeholder in order to try to find a match and replace
// the word with defined variable
func (p *Parser) Render(word string, lineNum int) (string, error) {
	if strings.HasSuffix(word, p.Delimeter) {
		w, err := p.Replace(word)
		if err != nil {
			return "", fmt.Errorf("Could not found variable %q, on line %d please verify the variables file", word, lineNum)
		}
		return w, nil
	}
	for k, c := range word[1:] {
		if string(c) == p.Delimeter {
			reminder := word[k+2:]
			w, err := p.Replace(word[:k+1])
			if err != nil {
				return "", fmt.Errorf("Could not found variable %q, on line %d please verify the variables file", word[:k+1]+"$", lineNum)
			}
			word = w + reminder
		}
	}
	return word, nil
}
