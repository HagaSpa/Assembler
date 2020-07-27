package main

import (
	"bytes"
	"io"
	"testing"
)

type testWriteCloser struct {
	io.Writer
}

func (testWriteCloser) Close() error { return nil }
func (testWriteCloser) WriteString(s string) (n int, err error) {
	b := &bytes.Buffer{}
	b.Write(outputMap[fileName].Bytes())
	b.WriteString(s)
	outputMap[fileName] = b
	return 0, nil
}

var (
	outputMap map[string]*bytes.Buffer
	fileName  string
)

// Confirm that it will be written with a line break
func Test_writeLine(t *testing.T) {
	createFile := func(name string) (WriteCloser, error) {
		b := &bytes.Buffer{}
		fileName = name
		outputMap[name] = b
		return testWriteCloser{b}, nil
	}
	createFileIF = createFile

	tests := []struct {
		caseName string
		name     string
		b        []string
		want     map[string]string
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
			map[string]string{
				"test.hack": "0000000000000010\n1110110000010000\n0000000000000011\n1110000010010000\n0000000000000000\n1110001100001000\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			outputMap = map[string]*bytes.Buffer{}
			writeLine(tt.name, tt.b)

			if len(tt.want) != len(outputMap) {
				t.Errorf("the number of output files doesn't match (expected=%d, actual=%d)", len(tt.want), len(outputMap))
				return
			}
			for k, v := range tt.want {
				if a, ok := outputMap[k]; !ok {
					t.Errorf("file doesn't exist in output files (filename=%s)", k)
					return
				} else if a.String() != v {
					t.Errorf("file contents don't match (expected=%s, actual=%s)", v, a.String())
					return
				}
			}
		})
	}
}
