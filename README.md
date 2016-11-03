[![Build Status](https://travis-ci.org/nbari/parser.svg?branch=master)](https://travis-ci.org/nbari/parser)
[![codecov](https://codecov.io/gh/nbari/parser/branch/master/graph/badge.svg)](https://codecov.io/gh/nbari/parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/nbari/parser)](https://goreportcard.com/report/github.com/nbari/parser)

# parser

Parse a template using variables in a yaml file


Usage:

    $ parser -h


## Compile from source

Setup go environment https://golang.org/doc/install

For example using $HOME/go for your workspace

    $ export GOPATH=$HOME/go

Create the directory:

    $ mkdir -p $HOME/go/src/github.com/nbari

Clone project into that directory:

    $ git clone git@github.com:nbari/parser.git $HOME/go/src/github.com/nbari/parser

Build by just typing make:

    $ cd $HOME/go/src/github.com/nbari/parser
    $ make


# Test & Coverage

Run the test using:

    $ make test

Check coverage:

    $ make cover
