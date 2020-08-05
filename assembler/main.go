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
	"strconv"
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
	defer fp.Close()
	s := bufio.NewScanner(fp)

	// generate symbol table
	tp := parser.New(s)
	t := genTable(tp)
	fmt.Println(t)

	// reset file pointer
	fp.Seek(0, 0)
	s = bufio.NewScanner(fp)

	// generate binary
	var b []string
	p := parser.New(s)
	addr := 0x0010 // start to 16
	for p.HasMoreCommands() {
		p.Advance()
		if p.Type == "" || p.Type == parser.L_COMMAND {
			continue
		}
		if p.Type == parser.A_COMMAND {
			// Only Character
			_, ok := strconv.Atoi(p.Symbol)
			if ok != nil {
				// Convert to decimal if symbol is already displayed
				if t.Contains(p.Symbol) {
					a := t.GetAddress(p.Symbol)
					p.Symbol = fmt.Sprintf("%d", a)
				} else {
					t.AddEntry(p.Symbol, addr)
					addr++
				}
			}
		}
		c := code.New(p, t)
		b = append(b, c.Binary)
	}

	// generate .hack file
	rep := regexp.MustCompile(`.asm$`)
	name := filepath.Base(rep.ReplaceAllString(flags[0], "")) + ".hack"
	f, err := os.Create(name)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	writeLine(f, b)
}

func writeLine(f io.Writer, b []string) {
	w := bufio.NewWriter(f)
	for _, line := range b {
		// TODO: Unix Only?
		_, err := w.WriteString(line + "\n")
		if err != nil {
			os.Exit(1)
		}
		w.Flush()
	}
}

func genTable(p *parser.Parser) symbol.Table {
	t := symbol.New()
	addr := 0x0000 // start to 0
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
