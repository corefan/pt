package pt

import "image"

type Pixel struct {
	Samples int
	M, V    Color
}

func (p *Pixel) Color() Color {
	return p.M.DivScalar(float64(p.Samples))
}

func (p *Pixel) Variance() Color {
	if p.Samples < 2 {
		return Black
	}
	return p.V.DivScalar(float64(p.Samples - 1)).Pow(0.5)
}

func (p *Pixel) AddSample(sample Color) {
	p.Samples++
	if p.Samples == 1 {
		p.M = sample
		return
	}
	oldMean := p.M
	p.M = p.M.Add(sample.Sub(p.M).DivScalar(float64(p.Samples)))
	p.V = p.V.Add(sample.Sub(oldMean).Mul(sample.Sub(p.M)))
}

type Buffer struct {
	W      int
	H      int
	Pixels []Pixel
}

func NewBuffer(w, h int) *Buffer {
	pixels := make([]Pixel, w*h)
	return &Buffer{w, h, pixels}
}

func (b *Buffer) Image() image.Image {
	result := image.NewRGBA64(image.Rect(0, 0, b.W, b.H))
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			c := b.Pixels[y*b.W+x].Color()
			result.SetRGBA64(x, y, c.Pow(1/2.2).RGBA64())
		}
	}
	return result
}

func (b *Buffer) AddSample(x, y int, sample Color) {
	b.Pixels[y*b.W+x].AddSample(sample)
}
