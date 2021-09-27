package main

import "testing"

func TestCesar_Crypt(t *testing.T) {
	type fields struct {
		Key int
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test 1",
			fields: fields{
				Key: 1,
			},
			args:   args{
				text: "abc",
			},
			want:   "bcd",
		},
		{
			name:   "test 2",
			fields: fields{
				Key: 1,
			},
			args:   args{
				text: "xyz",
			},
			want:   "yza",
		},
		{
			name:   "test 3",
			fields: fields{
				Key: 1,
			},
			args:   args{
				text: "Az9",
			},
			want:   "Ba0",
		},
		{
			name:   "test 4",
			fields: fields{
				Key: 4,
			},
			args:   args{
				text: "Ax8",
			},
			want:   "Eb2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cesar{
				Key: tt.fields.Key,
			}
			if got := c.Crypt(tt.args.text); got != tt.want {
				t.Errorf("Crypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCesar_DecryptCrypt(t *testing.T) {
	type fields struct {
		Key int
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test 1",
			fields: fields{
				Key: 1,
			},
			args:   args{
				text: "bcd",
			},
			want:   "abc",
		},
		{
			name:   "test 2",
			fields: fields{
				Key: 1,
			},
			args:   args{
				text: "yza",
			},
			want:   "xyz",
		},
		{
			name:   "test 3",
			fields: fields{
				Key: 1,
			},
			args:   args{
				text: "Ba0",
			},
			want:   "Az9",
		},
		{
			name:   "test 4",
			fields: fields{
				Key: 4,
			},
			args:   args{
				text: "Eb2",
			},
			want:   "Ax8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cesar{
				Key: tt.fields.Key,
			}
			if got := c.Decrypt(tt.args.text); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}