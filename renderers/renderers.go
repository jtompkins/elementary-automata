package renderers

import "joshtompkins.com/elementary-automata/generation"

type Renderer interface {
	Render(generations []generation.Generation)
}
