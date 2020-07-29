package symbol

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Table
	}{
		{
			"test1",
			Table{
				"SP":     0x0000,
				"LCL":    0x0001,
				"ARG":    0x0002,
				"THIS":   0x0003,
				"THAT":   0x0004,
				"R0":     0x0000,
				"R1":     0x0001,
				"R2":     0x0002,
				"R3":     0x0003,
				"R4":     0x0004,
				"R5":     0x0005,
				"R6":     0x0006,
				"R7":     0x0007,
				"R8":     0x0008,
				"R9":     0x0009,
				"R10":    0x000a,
				"R11":    0x000b,
				"R12":    0x000c,
				"R13":    0x000d,
				"R14":    0x000e,
				"R15":    0x000f,
				"SCREEN": 0x4000,
				"KBD":    0x6000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddEntry(t *testing.T) {
	type args struct {
		symbol  string
		address int
	}
	tests := []struct {
		name string
		args args
		want Table
	}{
		{
			"test1",
			args{
				symbol:  "test",
				address: 1,
			},
			Table{
				"SP":     0x0000,
				"LCL":    0x0001,
				"ARG":    0x0002,
				"THIS":   0x0003,
				"THAT":   0x0004,
				"R0":     0x0000,
				"R1":     0x0001,
				"R2":     0x0002,
				"R3":     0x0003,
				"R4":     0x0004,
				"R5":     0x0005,
				"R6":     0x0006,
				"R7":     0x0007,
				"R8":     0x0008,
				"R9":     0x0009,
				"R10":    0x000a,
				"R11":    0x000b,
				"R12":    0x000c,
				"R13":    0x000d,
				"R14":    0x000e,
				"R15":    0x000f,
				"SCREEN": 0x4000,
				"KBD":    0x6000,
				"test":   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			got.AddEntry(tt.args.symbol, tt.args.address)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		symbol string
		want   bool
	}{
		{
			"test1",
			"SP",
			true,
		},
		{
			"test2",
			"NONE",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := New()
			got := table.Contains(tt.symbol)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contains = %v, want %v", got, tt.want)
			}
		})
	}
}
