package parser

import (
	"bufio"
	"bytes"
	"fmt"
)

// Parse parse the template
func (p *Parser) Parse() (string, error) {
	var (
		out    bytes.Buffer
		buf    bytes.Buffer
		useBuf bool
		//	err error
		//inLoop      bool
		//inLoopBody  bool
		//lineBuffer  []string
		lineNum int
	//		placeHolder string
	//		position int
	//		variable string
	)

	scanner := bufio.NewScanner(p.template)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		lineNum++
		c := scanner.Text()
		switch c {
		case "\n":
			buf.WriteString(c)
			out.WriteString(buf.String())
			lineNum++
			buf.Reset()
		case "$":
			useBuf = true
			buf.WriteString(c)
		default:
			if useBuf {
				switch c {
				case "$":
					fmt.Printf("variable = %+v\n", buf.String())
				default:
					buf.WriteString(c)
				}
			} else {
				out.WriteString(c)
			}
		}
	}

	return out.String(), nil
}
