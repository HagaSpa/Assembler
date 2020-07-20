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
