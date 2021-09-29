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
				from: 'A',
				to:   'B',
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				from: 'B',
				to:   'D',
			},
			want: 2,
		},
		{
			name: "10",
			args: args{
				from: 'A',
				to:   'K',
			},
			want: 10,
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
