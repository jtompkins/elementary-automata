package renderers

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"

	"github.com/fogleman/gg"
	"joshtompkins.com/elementary-automata/generation"
)

type AnimatedGifRenderer struct {
	opts *RenderOptions
}

func NewImageRenderer(opts *RenderOptions) *AnimatedGifRenderer {
	return &AnimatedGifRenderer{opts: opts}
}

func (r *AnimatedGifRenderer) Render(generations []generation.Generation) {
	scale := r.opts.Scale
	x := len(generations[0]) * scale
	y := len(generations) * scale

	images := []*image.Paletted{}
	delays := []int{}

	// create the image
	dc := gg.NewContext(x, y)

	// draw a white background
	dc.DrawRectangle(0, 0, float64(x), float64(y))
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

		// draw the current state of the context as new frame
		img := dc.Image()
		bounds := img.Bounds()
		palettedImage := image.NewPaletted(bounds, palette.WebSafe)
		draw.Draw(palettedImage, palettedImage.Rect, img, bounds.Min, draw.Over)

		// add the new frame to the list of frames
		images = append(images, palettedImage)
		delays = append(delays, 0)
	}

	// save the image
	//dc.SavePNG(r.opts.File)
	f, err := os.OpenFile(r.opts.File, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
