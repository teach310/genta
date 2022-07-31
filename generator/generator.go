package generator

import (
	"bytes"
	"path/filepath"
	"text/template"
)

type Generator struct {
	TemplatesPath string
}

func (g *Generator) Run(csharpFile *CSharpFile) (string, error) {
	filepath := filepath.Join(g.TemplatesPath, "example.pb.cs.tmpl")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, csharpFile); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
