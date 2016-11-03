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
						word, err = p.Render(word, lineNum)
						if err != nil {
							return "", err
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
						placeHolder = fmt.Sprintf("%s%s%s", p.Delimeter, word, p.Delimeter)
					case 1:
						if word != "in" {
							return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
					case 2:
						if !strings.HasSuffix(word, p.Delimeter) {
							return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
						}
						variable = strings.Replace(word, p.Delimeter, "", -1)
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

		if lineBuffer != nil {
			buffer.WriteString(fmt.Sprintf("%s\n", strings.Join(lineBuffer, " ")))
			lineBuffer = nil
		}

	}
	return buffer.String(), nil
}
