package sample

import "testing"

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FizzBuzz",
			args: args{n: 15},
			want: "FizzBuzz",
		},
		{
			name: "Fizz",
			args: args{n: 3},
			want: "Fizz",
		},
		{
			name: "Buzz",
			args: args{n: 5},
			want: "Buzz",
		},
		{
			name: "Number",
			args: args{n: 1},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.n); got != tt.want {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
