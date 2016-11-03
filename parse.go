package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// Parse
func (p *Parser) Parse() (string, error) {
	s := bufio.NewScanner(p.template)
	var (
		buffer      bytes.Buffer
		err         error
		inLoop      bool
		inLoopBody  bool
		lineBuffer  []string
		lineNum     int
		placeHolder string
		position    int
		variable    string
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
							word, err = p.Replace(word, lineNum)
							if err != nil {
								return "", err
							}
						} else {
							for k, c := range word[1:] {
								if string(c) == p.Delimeter {
									reminder := word[k+2:]
									w, err := p.Replace(word[:k+1], lineNum)
									if err != nil {
										return "", err
									}
									word = w + reminder
								}
							}
						}
						lineBuffer = append(lineBuffer, word)
					}
					continue
				}
			}
			// For block
			if inLoop {
				if !inLoopBody {
					switch position {
					case 0:
						if strings.HasPrefix(word, p.Delimeter) && strings.HasSuffix(word, p.Delimeter) {
							return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
						placeHolder = fmt.Sprintf("%s%s%s", p.Delimeter, word, p.Delimeter)
					case 1:
						if word != "in" {
							return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
					case 2:
						if strings.HasPrefix(word, p.Delimeter) {
							return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
						if !strings.HasSuffix(word, p.Delimeter) {
							return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
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
								ln := fmt.Sprintf("%s\n", strings.Replace(line, placeHolder, v.(string), -1))
								buffer.WriteString(ln)
							}
						}
						inLoopBody = false
					} else {
						return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
					}
				}
			} else {
				lineBuffer = append(lineBuffer, word)
			}
		}
		buffer.WriteString(fmt.Sprintf("%s\n", strings.Join(lineBuffer, " ")))
		lineBuffer = []string{}
	}
	return buffer.String(), nil
}
