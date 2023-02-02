package strutil

import "testing"

func TestToUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s: "Hello World",
			},
			want: "HELLO WORLD",
		},
		{
			name: "case2",
			args: args{
				s: "你好，世界",
			},
			want: "你好，世界",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.s); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
