package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nbari/parser"
)

func exit1(err error) {
	fmt.Println(err)
	os.Exit(1)
}

var version string

func main() {
	var (
		t  = flag.String("t", "", "`template` file to parse")
		v  = flag.String("v", "", "`variables` file to use")
		vv = flag.Bool("version", false, fmt.Sprintf("Print version: %s", version))
	)

	flag.Parse()

	if *vv {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if *t == "" {
		exit1(fmt.Errorf("Missing template, use (\"%s -h\") for help.\n", os.Args[0]))
	}

	if *v == "" {
		exit1(fmt.Errorf("Missing variables, use (\"%s -h\") for help.\n", os.Args[0]))
	}

	p, err := parser.New(*t, *v)
	if err != nil {
		exit1(err)
	}
	defer p.CloseTemplate()

	out, err := p.Parse()
	if err != nil {
		exit1(err)
	}
	fmt.Print(out)
}
