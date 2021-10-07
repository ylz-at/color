package color

import "testing"

func TestPlot(t *testing.T) {
	t.Log("generating CIE-1931-standard-colorimetric-observer-rgb.txt")
	xyz, err := load("CIE-1931-standard-colorimetric-observer-xyz.txt")
	if err != nil {
		t.Error(err)
	}
	if err := saveRGB(xyz, "CIE-1931-standard-colorimetric-observer-rgb.txt"); err != nil {
		t.Error(err)
	}
	type args struct {
		data   Data
		title  string
		format string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"CIE-1931-standard-colorimetric-observer-xyz.txt",
			args{
				title:  "CIE-1931-standard-colorimetric-observer-xyz",
				format: ".png",
			},
			false,
		},
		{
			"CIE-1931-standard-colorimetric-observer-rgb.txt",
			args{
				title:  "CIE-1931-standard-colorimetric-observer-rgb",
				format: ".png",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := load(tt.name)
			if err != nil {
				return
			}
			if err := Plot(data, tt.args.title, tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("Plot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
