package generator

import (
	"bytes"
	"text/template"
)

type Generator struct {
}

func (g *Generator) Run(csharpFile *CSharpFile) (string, error) {
	filepath := "templates/example.pb.cs.tmpl"
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
