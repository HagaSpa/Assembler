package parser

import (
	"bufio"
	"fmt"
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
	fmt.Println(line)
}
