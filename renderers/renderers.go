package renderers

import "joshtompkins.com/elementary-automata/generation"

type Renderer interface {
	Render(generations []generation.Generation)
}

type RenderOptions struct {
	File  string
	Scale int
}
