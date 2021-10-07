package color

import "testing"

func Test_xyz2rgb(t *testing.T) {
	type args struct {
		x   float64
		y   float64
		in2 float64
	}
	tests := []struct {
		name  string
		args  args
		wantR float64
		wantG float64
		wantB float64
	}{
		{
			"normal",
			args{
				x:   0.25,
				y:   0.40,
				in2: 0.10,
			},
			0.4174,
			0.7434,
			0.2152,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := xyz2rgb(tt.args.x, tt.args.y, tt.args.in2)
			if gotR != tt.wantR {
				t.Errorf("xyz2rgb() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("xyz2rgb() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("xyz2rgb() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
