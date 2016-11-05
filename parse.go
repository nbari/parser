package parser

import (
	"bufio"
	"bytes"
)

// Parse parse the template
func (p *Parser) Parse() (string, error) {
	var (
		out    bytes.Buffer
		buf    bytes.Buffer
		useBuf bool
		//	err error
		inLoop bool
		//inLoopBody  bool
		//lineBuffer  []string
		lineNum int = 1
	//		placeHolder string
	//		position int
	//		variable string
	)

	scanner := bufio.NewScanner(p.template)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		c := scanner.Text()
		switch c {
		case "\n":
			lineNum++
			useBuf = false
			buf.WriteString(c)
			out.WriteString(buf.String())
			buf.Reset()
		case p.Delimeter:
			buf.WriteString(c)
			if useBuf && len(buf.String()) == 2 {
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
			} else {
				useBuf = true
			}
		default:
			if useBuf {
				buf.WriteString(c)
				if c == " " {
					switch buf.String() {
					case "$for ":
						inLoop = true
					default:
						if inLoop {
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

	return out.String(), nil
}
