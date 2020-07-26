package main

import (
	"assembler/code"
	"assembler/parser"
	"bufio"
	"flag"
	"os"
	"path/filepath"
	"regexp"
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

	// generate .hack file
	rep := regexp.MustCompile(`.asm$`)
	name := filepath.Base(rep.ReplaceAllString(flags[0], "")) + ".hack"
	writeLine(name, b)

	defer fp.Close()
}

func writeLine(name string, b []string) {
	fp, err := os.Create(name)
	if err != nil {
		os.Exit(1)
	}
	defer fp.Close()

	for _, line := range b {
		// TODO: Unix Only?
		_, err := fp.WriteString(line + "\n")
		if err != nil {
			os.Exit(1)
		}
	}
}
