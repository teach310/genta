package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/teach310/genta"
	"github.com/teach310/genta/internal/version"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}

	if err := genta.Run(); err != nil {
		log.Fatalln(err)
	}
}
