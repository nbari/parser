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
