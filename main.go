package main

import (
	"html/template"
	"io"
)

// RenderTemplate renders the template file provided
// Context is interpretted from the context []string argument
// Returns the returns a string of the rendered template
func RenderTemplate(html string, context []string) error {
	temp, err := template.ParseFiles(html)
	if err != nil {
		return err
	}

	var out io.Writer
	err = temp.Execute(out, context)

	return err
}
