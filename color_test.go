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

func Test_rgb2xyz(t *testing.T) {
	type args struct {
		r float64
		g float64
		b float64
	}
	tests := []struct {
		name  string
		args  args
		wantX float64
		wantY float64
		wantZ float64
	}{
		{
			"normal",
			args{
				r: 0.4171,
				g: 0.7434,
				b: 0.2152,
			},
			0.25,
			0.40,
			0.10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotZ := rgb2xyz(tt.args.r, tt.args.g, tt.args.b)
			if gotX != tt.wantX {
				t.Errorf("rgb2xyz() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("rgb2xyz() gotY = %v, want %v", gotY, tt.wantY)
			}
			if gotZ != tt.wantZ {
				t.Errorf("rgb2xyz() gotZ = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}
