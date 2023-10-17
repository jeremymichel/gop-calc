package gop_calc

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		fps    int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 int
	}{
		{name: "GOP at 25fps for 2s", args: args{
			fps:    25,
			target: 2,
		}, want: 1.92, want1: 48},
		{name: "GOP at 25fps for 4s", args: args{
			fps:    25,
			target: 4,
		}, want: 3.84, want1: 96},
		{name: "GOP at 30fps for 2s", args: args{
			fps:    30,
			target: 2,
		}, want: 1.6, want1: 48},
		{name: "GOP at 30fps for 4s", args: args{
			fps:    30,
			target: 4,
		}, want: 3.7333333333333334, want1: 112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Calculate(tt.args.fps, 48000, 1024, tt.args.target)
			if got != tt.want {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Calculate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
