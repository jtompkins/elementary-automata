package renderers

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	log "github.com/sirupsen/logrus"
	"joshtompkins.com/elementary-automata/generation"
)

type PngRenderer struct {
	opts *RenderOptions
}

func NewPngRenderer(opts *RenderOptions) *PngRenderer {
	return &PngRenderer{opts: opts}
}

func (r *PngRenderer) Render(generations []generation.Generation) {
	scale := r.opts.Scale
	width := len(generations[0]) * scale
	height := len(generations) * scale

	// create the image
	img := image.NewGray(image.Rect(0, 0, width, height))

	// draw a white background
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	for iY, generation := range generations {
		for iX, cell := range generation {
			if !cell {
				continue
			}

			r.drawCell(img, iX, iY, color.Black)
		}
	}

	f, err := os.Create(r.opts.File)
	defer f.Close()

	if err != nil {
		log.WithField("filename", r.opts.File).Error("Unable to save rendered image")
		panic(err)
	}

	err = png.Encode(f, img)

	if err != nil {
		panic(err)
	}
}

func (r *PngRenderer) drawCell(img draw.Image, locale, generation int, c color.Color) {
	x := locale * r.opts.Scale
	y := generation * r.opts.Scale

	for iX := x; iX < x+r.opts.Scale; iX++ {
		for iY := y; iY < y+r.opts.Scale; iY++ {
			img.Set(iX, iY, c)
		}
	}
}
