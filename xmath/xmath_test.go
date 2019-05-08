package xmath

import (
	"testing"
)

func TestFloatEquals(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0 == 0", args{0, 0}, true},
		{"0.00000001 == 0.000000012", args{0.00000001, 0.000000012}, true},
		{"1.21 == 1.2", args{1.21, 1.2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatEquals(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("FloatEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatIsZero(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"-0", args{-0}, true},
		{"0", args{0}, true},
		{"0.000000005", args{0.000000005}, true},
		{"0.00000001", args{0.00000001}, false},
		{"0.0000001", args{0.0000001}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatIsZero(tt.args.v); got != tt.want {
				t.Errorf("FloatIsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		input  float64
		places int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"12.3456 p1", args{12.3456, 1}, 12.3},
		{"12.3456 p2", args{12.3456, 2}, 12.35},
		{"12.3456 p3", args{12.3456, 3}, 12.346},
		{"12.3456 p4", args{12.3456, 4}, 12.3456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.input, tt.args.places); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
