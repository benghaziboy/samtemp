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

// Returns an Email struct with the following arguments
// subject: The subject title of the email.
// sender: Email address of the sending email account.
// htmlFile: Filepath to the .html template to be rendered.
// textFile: Filepath to the .txt template to rendered.
// recipients: An array of email address for the email to be delivered to.
// context: an object that maps the templates keywords with their intended values.
func NewEmail(subject, sender, htmlFile, textFile string, recipients []string, context interface{}) (*Email, error) {
	email := Email{
		Subject:    subject,
		Sender:     sender,
		Recipients: recipients,
		Html:       htmlFile,
		Text:       textFile,
		Context:    context,
	}

	return email, nil
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
