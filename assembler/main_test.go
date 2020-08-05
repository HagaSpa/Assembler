package main

import (
	"bytes"
	"testing"
)

// Confirm that it will be written with a line break
func Test_writeLine(t *testing.T) {
	tests := []struct {
		caseName string
		name     string
		b        []string
		want     string
	}{
		{
			"test1",
			"test.hack",
			[]string{
				"0000000000000010",
				"1110110000010000",
				"0000000000000011",
				"1110000010010000",
				"0000000000000000",
				"1110001100001000",
			},
			`0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			b := bytes.NewBufferString("")
			writeLine(b, tt.b)
			if string(b.Bytes()) != tt.want {
				t.Errorf("writeLine() = %s, want %v", b, tt.want)
			}
		})
	}
}
