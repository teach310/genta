package generator

import (
	"fmt"
)

type Generator struct {
}

func (g *Generator) Run() error {
	fmt.Printf("Hello World\n")
	return nil
}
