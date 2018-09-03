package renderers

import (
	"image/color"

	"github.com/fogleman/gg"
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
	dc := gg.NewContext(width, height)

	// draw a white background
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetColor(color.White)
	dc.Fill()

	for iY, generation := range generations {
		for iX, cell := range generation {
			if !cell {
				continue
			}

			dc.DrawRectangle(float64(iX*scale), float64(iY*scale), float64(scale), float64(scale))
			dc.SetColor(color.Black)
			dc.Fill()
		}
	}

	err := dc.SavePNG(r.opts.File)

	if err != nil {
		log.WithField("filename", r.opts.File).Error("Unable to save rendered image")
		panic("")
	}
}
