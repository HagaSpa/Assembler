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
			b1 := p.hasMoreCommands()
			if !reflect.DeepEqual(b1, true) {
				t.Errorf("hasMoreCommands() = %v, want %v", b1, true)
			}

			// Proceed to EOF
			for p.scanner.Scan() {
				// pass
			}
			// false
			b2 := p.hasMoreCommands()
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
				scanner:     s,
				commandType: A_COMMAND,
				symbol:      "2",
			},
		},
		{
			"line2",
			&Parser{
				scanner:     s,
				commandType: C_COMMAND,
				dest:        "D",
				comp:        "A",
			},
		},
		{
			"line3",
			&Parser{
				scanner:     s,
				commandType: A_COMMAND,
				symbol:      "3",
			},
		},
		{
			"line4",
			&Parser{
				scanner:     s,
				commandType: C_COMMAND,
				dest:        "D",
				comp:        "D+A",
			},
		},
		{
			"line5",
			&Parser{
				scanner:     s,
				commandType: A_COMMAND,
				symbol:      "0",
			},
		},
		{
			"line6",
			&Parser{
				scanner:     s,
				commandType: C_COMMAND,
				dest:        "M",
				comp:        "D",
			},
		},
	}

	p := New(s)
	i := 0
	for p.hasMoreCommands() {
		p.advance()
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
