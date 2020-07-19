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
			for p.hasMoreCommands() {
				p.advance()
			}
		})
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
