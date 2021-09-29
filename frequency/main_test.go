package main

import "testing"

func Test_getKey(t *testing.T) {
	type args struct {
		from rune
		to   rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				from: 'B',
				to:   'A',
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				from: 'D',
				to:   'B',
			},
			want: 2,
		},
		{
			name: "10",
			args: args{
				from: 'K',
				to:   'A',
			},
			want: 10,
		},
		{
			name: "25",
			args: args{
				from: 'Z',
				to:   'A',
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKey(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("getKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReversedKey(t *testing.T) {
	type args struct {
		from rune
		to   rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				from: 'A',
				to:   'Z',
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReversedKey(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("getKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
