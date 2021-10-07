package color

import (
	"gonum.org/v1/gonum/mat"
)

func rgb2xyz(r, g, b float64) (x, y, z float64) {
	x = (0.49000*r + 0.31000*g + 0.20000*b) / (0.66697*r + 1.13240*g + 1.20063*b)
	y = (0.17697*r + 0.81240*g + 0.01063*b) / (0.66697*r + 1.13240*g + 1.20063*b)
	z = (0.00000*r + 0.01000*g + 0.99000*b) / (0.66697*r + 1.13240*g + 1.20063*b)
	return
}

func tristimulus2Chromaticity(r1, g1, b1 float64) (r, g, b float64) {
	r = r1 / (r1 + g1 + b1)
	g = g1 / (r1 + g1 + b1)
	b = b1 / (r1 + g1 + b1)
	return
}

func chromaticity2Tristimulus(r, g, b float64) (r1, g1, b1 float64) {
	// TODO
	return 0, 0, 0
}

//func tristimulus2ChromaticityXYZ(x1, y1, z1 float64) (x, y, z float64) {
//	v := y1 // refer to CIE
//
//
//	return
//}

func xyz2rgb(x, y, _ float64) (r, g, b float64) {
	/*
		|X|   |b11 b12 b13| |R|
		|Y| = |b21 b22 b23| |G|
		|Z|   |b31 b32 b33| |B|
	*/
	b11, b12, b13 := 2.7689, 1.7517, 1.1302
	b21, b22, b23 := 1.0000, 4.5907, 0.0601
	b31, b32, b33 := 0.0000, 0.0565, 5.5943
	// simplified b
	β11 := b11 - b13
	β12 := b12 - b13
	β13 := b13
	β21 := b21 - b23
	β22 := b22 - b23
	β23 := b23
	β31 := (b11 - b13) + (b21 - b23) + (b31 - b33)
	β32 := (b12 - b13) + (b22 - b23) + (b32 - b33)
	β33 := b13 + b23 + b33

	r = mat.Det(mat.NewDense(3, 3, []float64{
		x, β12, β13,
		y, β22, β23,
		1, β32, β33,
	})) / mat.Det(mat.NewDense(3, 3, []float64{
		β11, β12, x,
		β21, β22, y,
		β31, β32, 1,
	}))
	g = -mat.Det(mat.NewDense(3, 3, []float64{
		β11, x, β13,
		β21, y, β23,
		β31, 1, β33,
	})) / mat.Det(mat.NewDense(3, 3, []float64{
		β11, β12, x,
		β21, β22, y,
		β31, β32, 1,
	}))
	b = 1 - r - g
	return
}
