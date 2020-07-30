package main

import (
	"assembler/code"
	"assembler/parser"
	"assembler/symbol"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// WriteCloser is interface, for testing I/O
type WriteCloser interface {
	io.WriteCloser
	WriteString(s string) (n int, err error)
}

var createFileIF func(name string) (WriteCloser, error)

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

	// generate symbol table
	tp := parser.New(s)
	t := genTable(tp)
	fmt.Println(t)

	// generate binary
	var b []string
	p := parser.New(s)
	for p.HasMoreCommands() {
		p.Advance()
		if p.Type == "" || p.Type == parser.L_COMMAND {
			continue
		}
		// TODO: TypeがA_COMMANDなら、symbol_tableへの問い合わせと、追加を行う
		c := code.New(p)
		b = append(b, c.Binary)
	}

	// generate .hack file
	rep := regexp.MustCompile(`.asm$`)
	name := filepath.Base(rep.ReplaceAllString(flags[0], "")) + ".hack"
	createFileIF = func(name string) (WriteCloser, error) {
		return os.Create(name)
	}
	writeLine(name, b)

	defer fp.Close()
}

func writeLine(name string, b []string) {
	fp, err := createFileIF(name)
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

func genTable(p *parser.Parser) symbol.Table {
	t := symbol.New()
	addr := 0
	for p.HasMoreCommands() {
		p.Advance()
		switch p.Type {
		case parser.A_COMMAND, parser.C_COMMAND:
			addr++
		case parser.L_COMMAND:
			t.AddEntry(p.Symbol, addr)
		}
	}
	return t
}
