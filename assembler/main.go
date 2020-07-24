package main

import (
	"assembler/code"
	"assembler/parser"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// parse args
	flag.Parse()
	flags := flag.Args()
	if flags == nil {
		// TODO: err or start repl?
		os.Exit(1)
	}

	// open assenbly
	fp, err := os.Open(flags[0])
	if err != nil {
		os.Exit(1)
	}
	s := bufio.NewScanner(fp)

	// generate binary
	var b []string
	p := parser.New(s)
	for p.HasMoreCommands() {
		p.Advance()
		if p.Type == "" {
			continue
		}
		c := code.New(p)
		b = append(b, c.Binary)
	}

	fmt.Printf("%s \n", b)

	defer fp.Close()
}
