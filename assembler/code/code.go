package code

import (
	"assembler/parser"
	"assembler/symbol"
	"fmt"
	"strconv"
)

type Code struct {
	dest   string // 3bit
	comp   string // 7bit
	jump   string // 3bit
	Binary string // Binary Code
}

func New(p *parser.Parser, t symbol.Table) *Code {
	c := &Code{}
	if p.Type == parser.C_COMMAND {
		c.genDest(p.Dest)
		c.genComp(p.Comp)
		c.genJump(p.Jump)
		c.genBinaryC()
	} else {
		// A command
		is, ok := strconv.Atoi(p.Symbol)
		if ok != nil {
			is = t.GetAddress(p.Symbol)
		}
		c.genBinaryA(is)
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

func (c *Code) genComp(comp string) {
	switch comp {
	case "0":
		c.comp = "0101010"
	case "1":
		c.comp = "0111111"
	case "-1":
		c.comp = "0111010"
	case "D":
		c.comp = "0001100"
	case "A":
		c.comp = "0110000"
	case "!D":
		c.comp = "0001101"
	case "!A":
		c.comp = "0110001"
	case "-D":
		c.comp = "0001111"
	case "-A":
		c.comp = "0110011"
	case "D+1":
		c.comp = "0011111"
	case "A+1":
		c.comp = "0110111"
	case "D-1":
		c.comp = "0001110"
	case "A-1":
		c.comp = "0110010"
	case "D+A":
		c.comp = "0000010"
	case "D-A":
		c.comp = "0010011"
	case "A-D":
		c.comp = "0000111"
	case "D&A":
		c.comp = "0000000"
	case "D|A":
		c.comp = "0010101"
	case "M":
		c.comp = "1110000"
	case "!M":
		c.comp = "1110001"
	case "-M":
		c.comp = "1110011"
	case "M+1":
		c.comp = "1110111"
	case "M-1":
		c.comp = "1110010"
	case "D+M":
		c.comp = "1000010"
	case "D-M":
		c.comp = "1010011"
	case "M-D":
		c.comp = "1000111"
	case "D&M":
		c.comp = "1000000"
	case "D|M":
		c.comp = "1010101"
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

func (c *Code) genBinaryC() {
	c.Binary = "111" + c.comp + c.dest + c.jump
}

func (c *Code) genBinaryA(is int) {
	// convert to binary number
	bn := fmt.Sprintf("%b", is)
	// number of zeros to pad to 16
	cnt := 16 - len(bn)
	pad := ""
	for i := 0; i < cnt; i++ {
		pad = pad + "0"
	}
	c.Binary = pad + bn
}
