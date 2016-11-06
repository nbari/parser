package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// Parse parse the template
func (p *Parser) Parse() (string, error) {
	var (
		out    bytes.Buffer
		buf    bytes.Buffer
		useBuf bool
		//	err error
		inLoop       bool
		inLoopBody   bool
		lineNum      int = 1
		placeHolder  string
		position     int
		loopVariable string
	)

	scanner := bufio.NewScanner(p.template)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		c := scanner.Text()
		// handle for loop
		switch c {
		case "\n":
			lineNum++
			if inLoopBody {
				line := buf.String()
				buf.Reset()
				if strings.Contains(line, placeHolder) {
					// loop Body
					if k, ok := p.Variables[loopVariable]; ok {
						for _, v := range k.([]interface{}) {
							ln := fmt.Sprintf("%s\n",
								strings.Replace(line,
									placeHolder,
									fmt.Sprintf("%v", v),
									-1),
							)
							buf.WriteString(ln)
						}
					}
				}
			} else {
				buf.WriteString(c)
			}
			useBuf = false
			out.WriteString(buf.String())
			buf.Reset()
		case p.Delimeter:
			buf.WriteString(c)
			if useBuf && position == 2 {
				position = 0
				loopVariable = strings.Replace(buf.String(), p.Delimeter, "", -1)
				inLoopBody = true
				buf.Reset()
				continue
			} else if useBuf && len(buf.String()) == 2 {
				useBuf = false
				out.WriteString(buf.String())
				buf.Reset()
			} else if useBuf && !inLoop {
				useBuf = false
				variable := buf.String()
				buf.Reset()
				str, err := p.Render(variable, lineNum)
				if err != nil {
					return "", err
				}
				out.WriteString(str)
			} else if buf.String() == "$endfor$" {
				fmt.Printf("buf.String() = %+v\n", buf.String())
			} else {
				useBuf = true
			}
		default:
			if inLoopBody {
				buf.WriteString(c)
			} else if useBuf {
				buf.WriteString(c)
				if c == " " {
					switch buf.String() {
					case "$for ":
						inLoop = true
						buf.Reset()
					default:
						if inLoop {
							switch position {
							case 0:
								placeHolder = fmt.Sprintf("%s%s%s",
									p.Delimeter,
									strings.TrimSpace(buf.String()),
									p.Delimeter)
							case 1:
								if buf.String() != "in " {
									return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
								}
							case 2:
								return "", fmt.Errorf("Error parsing template, please verify the syntax on line %d", lineNum)
							}
							buf.Reset()
							position++
						} else {
							useBuf = false
							out.WriteString(buf.String())
							buf.Reset()
						}
					}
				}
			} else {
				out.WriteString(c)
			}
		}
	}
	fmt.Println(placeHolder, loopVariable)
	return out.String(), nil
}
