package genta

import (
	"github.com/teach310/genta/generator"
)

func Run() error {
	gen := &generator.Generator{}
	return gen.Run()
}
