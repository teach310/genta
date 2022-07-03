package genta

import (
	"github.com/teach310/genta/protogen"
)

func Run() error {
	protogen.NewPlugin().Run()
	return nil
}
