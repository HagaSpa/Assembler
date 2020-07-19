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
			if got := New(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readfile(fileName string) *bufio.Scanner {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	s := bufio.NewScanner(fp)
	return s
}
