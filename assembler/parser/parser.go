package parser

import (
	"bufio"
	"fmt"
	"strings"
)

type Parser struct {
	scanner *bufio.Scanner
	Type    string
	Symbol  string
	Dest    string
	Comp    string
	Jump    string
}

const (
	A_COMMAND = "A"
	C_COMMAND = "C"
	L_COMMAND = "L"
)

func New(scanner *bufio.Scanner) *Parser {
	p := &Parser{
		scanner: scanner,
	}
	return p
}

func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) Advance() {
	line := p.scanner.Text()
	if line == "" || strings.HasPrefix(line, "//") {
		return
	}
	p.genParser(line)
	fmt.Println(line)
}

/*
   A Command:
       @value
   C Command:
       dest=comp;jump
   comp is required and either dest or jump is empty. There is never both.
*/
func (p *Parser) genParser(line string) {
	// A command
	if strings.HasPrefix(line, "@") {
		p.Type = A_COMMAND
		p.Symbol = line[1:]
		p.Dest = ""
		p.Comp = ""
		p.Jump = ""
		return
	}

	// L command
	if strings.HasPrefix(line, "(") {
		p.Type = L_COMMAND
		p.Symbol = strings.Trim(line, "()")
		p.Dest = ""
		p.Comp = ""
		p.Jump = ""
		return
	}

	// C command
	p.Type = C_COMMAND
	p.Symbol = ""

	// contains dest
	if strings.Contains(line, "=") {
		ei := strings.Index(line, "=")
		p.Dest = line[:ei]
		p.Comp = line[ei+1:]
		p.Jump = ""
		return
	}

	// contains jump
	ji := strings.Index(line, ";")
	p.Jump = line[ji+1:]
	p.Comp = line[:ji]
	p.Dest = ""
}
