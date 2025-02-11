package sample

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Divide 4 by 2",
			args:    args{a: 4, b: 2},
			want:    2,
			wantErr: false,
		},
		{
			name:    "Divide 4 by 0",
			args:    args{a: 4, b: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
