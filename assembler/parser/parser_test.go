package parser

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	s := readfile("../../test.asm")

	tests := []struct {
		name string
		s    *bufio.Scanner
		want *Parser
	}{
		{
			"test1",
			s,
			&Parser{
				scanner: s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// New
			p := New(tt.s)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("New() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestHasMoreCommands(t *testing.T) {
	s := readfile("../../test.asm")

	tests := []struct {
		name string
		s    *bufio.Scanner
		want *Parser
	}{
		{
			"test1",
			s,
			&Parser{
				scanner: s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.s)
			// true
			b1 := p.HasMoreCommands()
			if !reflect.DeepEqual(b1, true) {
				t.Errorf("hasMoreCommands() = %v, want %v", b1, true)
			}

			// Proceed to EOF
			for p.scanner.Scan() {
				// pass
			}
			// false
			b2 := p.HasMoreCommands()
			if !reflect.DeepEqual(b2, false) {
				t.Errorf("hasMoreCommands() = %v, want %v", b2, false)
			}
		})
	}
}

func TestAdvance(t *testing.T) {
	s := readfile("../../test2.asm")
	tests := []struct {
		name string
		want *Parser
	}{
		{
			"line1",
			&Parser{
				scanner: s,
				Type:    A_COMMAND,
				Symbol:  "2",
			},
		},
		{
			"line2",
			&Parser{
				scanner: s,
				Type:    C_COMMAND,
				Dest:    "D",
				Comp:    "A",
			},
		},
		{
			"line3",
			&Parser{
				scanner: s,
				Type:    A_COMMAND,
				Symbol:  "3",
			},
		},
		{
			"line4",
			&Parser{
				scanner: s,
				Type:    C_COMMAND,
				Dest:    "D",
				Comp:    "D+A",
			},
		},
		{
			"line5",
			&Parser{
				scanner: s,
				Type:    A_COMMAND,
				Symbol:  "0",
			},
		},
		{
			"line6",
			&Parser{
				scanner: s,
				Type:    C_COMMAND,
				Dest:    "M",
				Comp:    "D",
			},
		},
		{
			"line7",
			&Parser{
				scanner: s,
				Type:    L_COMMAND,
				Symbol:  "TEST",
			},
		},
		{
			"line8",
			&Parser{
				scanner: s,
				Type:    A_COMMAND,
				Symbol:  "R0",
			},
		},
		{
			"line9",
			&Parser{
				scanner: s,
				Type:    C_COMMAND,
				Dest:    "D",
				Comp:    "M",
			},
		},
		{
			"line10",
			&Parser{
				scanner: s,
				Type:    C_COMMAND,
				Dest:    "D",
				Comp:    "M",
			},
		},
	}

	p := New(s)
	i := 0
	for p.HasMoreCommands() {
		p.Advance()
		t.Run(tests[i].name, func(t *testing.T) {
			if !reflect.DeepEqual(p, tests[i].want) {
				t.Errorf("advance() = %v, want %v", p, tests[i].want)
			}
		})
		i++
	}
}

func readfile(fileName string) *bufio.Scanner {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(fp)
	return s
}
