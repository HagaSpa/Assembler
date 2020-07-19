package parser

import (
	"bufio"
	"fmt"
	"strings"
)

type CommandType string

type Parser struct {
	scanner     *bufio.Scanner
	commandType CommandType
	symbol      string
	dest        string
	comp        string
	jump        string
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

func (p *Parser) hasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) advance() {
	line := p.scanner.Text()
	if line == "" || strings.HasPrefix(line, "//") {
		return
	}
	p.setCommand(line)
	fmt.Println(line)
}

/*
   A Command:
       @value
   C Command:
       dest=comp;jump
   comp is required and either dest or jump is empty. There is never both.
*/
func (p *Parser) setCommand(line string) {
	// A command
	if strings.HasPrefix(line, "@") {
		p.commandType = A_COMMAND
		p.symbol = line[1:]
		p.dest = ""
		p.comp = ""
		p.jump = ""
		return
	}

	// C command
	p.commandType = C_COMMAND
	p.symbol = ""

	// contains dest
	if strings.Contains(line, "=") {
		ei := strings.Index(line, "=")
		p.dest = line[:ei]
		p.comp = line[ei+1:]
		p.jump = ""
		return
	}

	// contains jump
	ji := strings.Index(line, ";")
	p.jump = line[ji+1:]
	p.comp = line[:ji]
	p.dest = ""
}
