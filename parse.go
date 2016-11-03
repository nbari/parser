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
		}
	}
	return nil
}
