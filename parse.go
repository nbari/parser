package parser

import (
	"bufio"
	"fmt"
	"strings"
)

func (p *Parser) Parse() error {
	s := bufio.NewScanner(p.template)
	//	var buf []string
	var (
		inLoop      bool
		inLoopBody  bool
		position    int
		placeHolder string
		variable    string
		lineNum     int
	)

	for s.Scan() {
		lineNum++
		line := s.Text()
		scanner := bufio.NewScanner(strings.NewReader(line))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			word := scanner.Text()
			if strings.HasPrefix(word, p.Delimeter) {
				switch word {
				case "$for":
					inLoop = true
					continue
				case "$endfor$":
					inLoop = false
					placeHolder = ""
					variable = ""
					continue
				default:
					if !inLoop {
						if strings.HasSuffix(word, p.Delimeter) {
							variable = strings.Replace(word, "$", "", -1)
							if k, ok := p.Variables[variable]; ok {
								word = k.(string)
							} else {
								return fmt.Errorf("Could not found variable %q, on line %d please verify the variables file", variable, lineNum)
							}
						} else {
							for k, c := range word[1:] {
								if string(c) == p.Delimeter {
									variable = strings.Replace(word[:k+1], "$", "", -1)
									if k, ok := p.Variables[variable]; ok {
										// TODO
										word = k.(string) + word[7:]
									} else {
										return fmt.Errorf("Could not found variable %q, on line %d please verify the variables file", variable, lineNum)
									}
								}
							}
						}
					}
				}
			}
			if inLoop {
				if !inLoopBody {
					switch position {
					case 0:
						if strings.HasPrefix(word, p.Delimeter) && strings.HasSuffix(word, p.Delimeter) {
							return fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
						placeHolder = fmt.Sprintf("%s%s%s", p.Delimeter, word, p.Delimeter)
					case 1:
						if word != "in" {
							return fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
					case 2:
						if strings.HasPrefix(word, p.Delimeter) {
							return fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
						if !strings.HasSuffix(word, p.Delimeter) {
							return fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
						variable = strings.Replace(word, "$", "", -1)
						inLoopBody = true
					}
					position++
				} else {
					if strings.Contains(line, placeHolder) {
						// loop Body
						if k, ok := p.Variables[variable]; ok {
							for _, v := range k.([]interface{}) {
								fmt.Println(strings.Replace(line, placeHolder, v.(string), -1))
							}
						}
						inLoopBody = false
					} else {
						return fmt.Errorf("xx  Error parsing template, please verify the syntax on line %d", lineNum)
					}
				}
			}
			// No loop
			fmt.Print(word, " ")
		}
		println()
	}
	return nil
}
