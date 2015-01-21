package samtemp

import (
	"bytes"
	"html/template"
)

type Email struct {
	Subject    string
	Sender     string
	Recipients []string
	Html       string
	Text       string
	Context    interface{}
}

// Returns a *bytes.Buffer value containing the rendered html template.
// Utilizes the Email.Html string representing the `template.html` file.
func (e Email) RenderHtml() (*bytes.Buffer, error) {
	return RenderTemplate(e.Html, e.Context)
}

// Returns a *bytes.Buffer value containing the rendered text template.
// Utilizes the Email.Html string representing the `template.text` file.
func (e Email) RenderText() (*bytes.Buffer, error) {
	return RenderTemplate(e.Text, e.Context)
}

// RenderTemplate renders the template file provided
// Context is interpretted from the context interface argument
// Currently all data associated with the interface must be relevant to the template.
// Returns a pointer to a bytes.Buffer interface and an error
func RenderTemplate(html string, context interface{}) (*bytes.Buffer, error) {
	temp, err := template.ParseFiles(html)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	err = temp.Execute(&out, context)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
