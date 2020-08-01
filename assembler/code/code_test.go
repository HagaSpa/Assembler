package code

import (
	"reflect"
	"testing"
)

func TestGenDest(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *Code
	}{
		{
			name: "test null (empty)",
			args: "",
			want: &Code{
				dest: "000",
			},
		},
		{
			name: "test M",
			args: "M",
			want: &Code{
				dest: "001",
			},
		},
		{
			name: "test D",
			args: "D",
			want: &Code{
				dest: "010",
			},
		},
		{
			name: "test MD",
			args: "MD",
			want: &Code{
				dest: "011",
			},
		},
		{
			name: "test A",
			args: "A",
			want: &Code{
				dest: "100",
			},
		},
		{
			name: "test AM",
			args: "AM",
			want: &Code{
				dest: "101",
			},
		},
		{
			name: "test AD",
			args: "AD",
			want: &Code{
				dest: "110",
			},
		},
		{
			name: "test AMD",
			args: "AMD",
			want: &Code{
				dest: "111",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Code{}
			c.genDest(tt.args)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("getDest() = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestGenComp(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *Code
	}{
		{
			name: "test 0",
			args: "0",
			want: &Code{
				comp: "0101010",
			},
		},
		{
			name: "test 1",
			args: "1",
			want: &Code{
				comp: "0111111",
			},
		},
		{
			name: "test -1",
			args: "-1",
			want: &Code{
				comp: "0111010",
			},
		},
		{
			name: "test D",
			args: "D",
			want: &Code{
				comp: "0001100",
			},
		},
		{
			name: "test A",
			args: "A",
			want: &Code{
				comp: "0110000",
			},
		},
		{
			name: "test !D",
			args: "!D",
			want: &Code{
				comp: "0001101",
			},
		},
		{
			name: "test !A",
			args: "!A",
			want: &Code{
				comp: "0110001",
			},
		},
		{
			name: "test -D",
			args: "-D",
			want: &Code{
				comp: "0001111",
			},
		},
		{
			name: "test -A",
			args: "-A",
			want: &Code{
				comp: "0110011",
			},
		},
		{
			name: "test D+1",
			args: "D+1",
			want: &Code{
				comp: "0011111",
			},
		},
		{
			name: "test A+1",
			args: "A+1",
			want: &Code{
				comp: "0110111",
			},
		},
		{
			name: "test D-1",
			args: "D-1",
			want: &Code{
				comp: "0001110",
			},
		},
		{
			name: "test A-1",
			args: "A-1",
			want: &Code{
				comp: "0110010",
			},
		},
		{
			name: "test D+A",
			args: "D+A",
			want: &Code{
				comp: "0000010",
			},
		},
		{
			name: "test D-A",
			args: "D-A",
			want: &Code{
				comp: "0010011",
			},
		},
		{
			name: "test A-D",
			args: "A-D",
			want: &Code{
				comp: "0000111",
			},
		},
		{
			name: "test D&A",
			args: "D&A",
			want: &Code{
				comp: "0000000",
			},
		},
		{
			name: "test D|A",
			args: "D|A",
			want: &Code{
				comp: "0010101",
			},
		},
		{
			name: "test M",
			args: "M",
			want: &Code{
				comp: "1110000",
			},
		},
		{
			name: "test !M",
			args: "!M",
			want: &Code{
				comp: "1110001",
			},
		},
		{
			name: "test -M",
			args: "-M",
			want: &Code{
				comp: "1110011",
			},
		},
		{
			name: "test M+1",
			args: "M+1",
			want: &Code{
				comp: "1110111",
			},
		},
		{
			name: "test M-1",
			args: "M-1",
			want: &Code{
				comp: "1110010",
			},
		},
		{
			name: "test D+M",
			args: "D+M",
			want: &Code{
				comp: "1000010",
			},
		},
		{
			name: "test D-M",
			args: "D-M",
			want: &Code{
				comp: "1010011",
			},
		},
		{
			name: "test M-D",
			args: "M-D",
			want: &Code{
				comp: "1000111",
			},
		},
		{
			name: "test D&M",
			args: "D&M",
			want: &Code{
				comp: "1000000",
			},
		},
		{
			name: "test D|M",
			args: "D|M",
			want: &Code{
				comp: "1010101",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Code{}
			c.genComp(tt.args)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("getComp() = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestGenJump(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *Code
	}{
		{
			name: "test null (empty)",
			args: "",
			want: &Code{
				jump: "000",
			},
		},
		{
			name: "test JGT",
			args: "JGT",
			want: &Code{
				jump: "001",
			},
		},
		{
			name: "test JEQ",
			args: "JEQ",
			want: &Code{
				jump: "010",
			},
		},
		{
			name: "test JGE",
			args: "JGE",
			want: &Code{
				jump: "011",
			},
		},
		{
			name: "test JLT",
			args: "JLT",
			want: &Code{
				jump: "100",
			},
		},
		{
			name: "test JNE",
			args: "JNE",
			want: &Code{
				jump: "101",
			},
		},
		{
			name: "test JLE",
			args: "JLE",
			want: &Code{
				jump: "110",
			},
		},
		{
			name: "test JMP",
			args: "JMP",
			want: &Code{
				jump: "111",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Code{}
			c.genJump(tt.args)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("getJump() = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestGenBinaryC(t *testing.T) {
	tests := []struct {
		name string
		args *Code
		want *Code
	}{
		{
			name: "test D=A",
			args: &Code{
				dest: "010",
				comp: "0110000",
				jump: "000",
			},
			want: &Code{
				dest:   "010",
				comp:   "0110000",
				jump:   "000",
				Binary: "1110110000010000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.genBinaryC()
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("genBinaryC() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestGenBinaryA(t *testing.T) {
	tests := []struct {
		name string
		args int
		want *Code
	}{
		{
			name: "test 2",
			args: 2,
			want: &Code{
				Binary: "0000000000000010",
			},
		},
		{
			name: "test 3",
			args: 3,
			want: &Code{
				Binary: "0000000000000011",
			},
		},
		{
			name: "test 0",
			args: 0,
			want: &Code{
				Binary: "0000000000000000",
			},
		},
		{
			name: "test 10",
			args: 10,
			want: &Code{
				Binary: "0000000000001010",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Code{}
			c.genBinaryA(tt.args)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("genBinaryA() = %v, want %v", c, tt.want)
			}
		})
	}
}
