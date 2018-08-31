package renderers

import (
	"fmt"

	"joshtompkins.com/elementary-automata/generation"
)

type CliRenderer struct{}

func NewCliRenderer() *CliRenderer {
	return &CliRenderer{}
}

func (r *CliRenderer) Render(generations []generation.Generation) {
	for _, g := range generations {
		for _, cell := range g {
			if cell {
				fmt.Print("X")
			} else {
				fmt.Print("_")
			}
		}

		fmt.Println("")
	}
}
