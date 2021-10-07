package color

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	stdfnt "golang.org/x/image/font"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/palette/brewer"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type Data []Sample

type Sample struct {
	Name   string
	Values []Value
}
type Value struct {
	X float64
	Y float64
}

func saveRGB(data Data, filename string) error {
	if len(data) != 3 {
		panic("data length should be 3")
	}
	f, err2 := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err2 != nil {
		return err2
	}
	if _, err := f.WriteString("λ r g b\n"); err != nil {
		return err
	}
	for i := 0; i < len(data[0].Values); i++ {
		lambda := data[0].Values[i].X
		x, y, z := data[0].Values[i].Y, data[1].Values[i].Y, data[2].Values[i].Y
		r, g, b := xyz2rgb(x, y, z)
		if _, err := fmt.Fprintf(f, "%.5f %.5f %.5f %.5f\n", lambda, r, g, b); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func load(filename string) (Data, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := make(Data, 3)
	data[0].Name = "r"
	data[1].Name = "g"
	data[2].Name = "b"
	for scanner.Scan() {
		lineText := scanner.Text()
		values := make([]float64, 4)
		for i, value := range strings.Split(lineText, " ") {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("sample value not float: %s", value)
			}
			values[i] = f
		}
		lambda := values[0]
		for i := 0; i+1 < len(values); i++ {
			data[i].Values = append(data[i].Values, Value{X: lambda, Y: values[i+1]})
		}
	}
	return data, nil
}

// Plot creates a plot from metric data and saves it to a temporary file.
// It's the callers' responsibility to remove the returned file when no longer needed.
func Plot(data Data, title, format string) error {
	p := plot.New()
	p.Title.Text = title
	p.Title.TextStyle.Font = font.From(font.Font{
		Typeface: "Liberation",
		Variant:  "Mono",
		Style:    stdfnt.StyleItalic,
		Weight:   stdfnt.WeightBold,
	}, 0.35*vg.Centimeter)
	p.Title.Padding = 1 * vg.Centimeter
	p.X.Tick.Marker = plot.DefaultTicks{}
	normalFont := font.From(font.Font{
		Typeface: "Liberation",
		Variant:  "Mono",
	}, 3*vg.Millimeter)
	p.X.Tick.Label.Font = normalFont
	p.Y.Tick.Label.Font = normalFont
	p.Legend.TextStyle.Font = normalFont
	p.Legend.Top = true
	p.Legend.YOffs = 15 * vg.Millimeter

	// Color palette for drawing lines
	paletteSize := 8
	palette, err := brewer.GetPalette(brewer.TypeAny, "Dark2", paletteSize)
	if err != nil {
		return fmt.Errorf("failed to get color palette: %v", err)
	}
	colors := palette.Colors()

	for s, sample := range data {
		data := make(plotter.XYs, len(sample.Values))
		for i, v := range sample.Values {
			data[i].X = v.X
			data[i].Y = v.Y
		}

		l, err := plotter.NewLine(data)
		if err != nil {
			return fmt.Errorf("failed to create line: %v", err)
		}
		l.LineStyle.Width = vg.Points(1)
		l.LineStyle.Color = colors[s%paletteSize]

		p.Add(l)
		if len(data) > 1 {
			if sample.Name != "" {
				p.Legend.Add(sample.Name, l)
			}
		}
	}

	// Draw plot in canvas with margin
	margin := 6 * vg.Millimeter
	width := 24 * vg.Centimeter
	height := 10 * vg.Centimeter
	format = strings.TrimPrefix(format, ".")
	c, err1 := draw.NewFormattedCanvas(width, height, format)
	if err1 != nil {
		return fmt.Errorf("failed to create canvas: %v", err1)
	}
	p.Draw(draw.Crop(draw.New(c), margin, -margin, margin, -margin))
	f, err2 := os.OpenFile(title+"."+format, os.O_CREATE|os.O_WRONLY, 0666)
	if err2 != nil {
		return err2
	}

	if _, err := c.WriteTo(f); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
