package code

import (
	"assembler/parser"
)

type Code struct {
	dest string // 3bit
	comp string // 7bit
	jump string // 3bit
	Code string // When A Command value. Return
}

func New(p *parser.Parser) *Code {
	c := &Code{}
	if p.Type == parser.C_COMMAND {
		c.genDest(p.Dest)
		c.genJump(p.Jump)
	}
	return c
}

func (c *Code) genDest(dest string) {
	switch dest {
	case "":
		c.dest = "000"
	case "M":
		c.dest = "001"
	case "D":
		c.dest = "010"
	case "MD":
		c.dest = "011"
	case "A":
		c.dest = "100"
	case "AM":
		c.dest = "101"
	case "AD":
		c.dest = "110"
	case "AMD":
		c.dest = "111"
	default:
		// err
	}
}

func (c *Code) genJump(jump string) {
	switch jump {
	case "":
		c.jump = "000"
	case "JGT":
		c.jump = "001"
	case "JEQ":
		c.jump = "010"
	case "JGE":
		c.jump = "011"
	case "JLT":
		c.jump = "100"
	case "JNE":
		c.jump = "101"
	case "JLE":
		c.jump = "110"
	case "JMP":
		c.jump = "111"
	default:
		// err
	}
}
