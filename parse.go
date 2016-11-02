package parser

import (
	"bufio"
	"fmt"
)

func (p *Parser) Parse() error {
	s := bufio.NewScanner(p.template)
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		fmt.Println(s.Text())
		if s.Text() == "\n" {
			print("salta")
		}
	}
	return nil
}
