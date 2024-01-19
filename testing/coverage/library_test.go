package main

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "two positives",
			args: args{2, 4},
			want: 6,
		},
		{
			name: "one negative with postive result",
			args: args{2, -1},
			want: 1,
		},
		{
			name: "one negative with negative result",
			args: args{2, -6},
			want: -4,
		},
		{
			name: "two negatives",
			args: args{-2, -6},
			want: -8,
		},
		{
			name: "integer overflow",
			args: args{0x7ffffffffffffffe, 2},
			want: -9223372036854775808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
