package main

import (
	"bytes"
	"errors"
	"html/template"
)

var (
	errNoHtmlTemplate = errors.New("An html template must be provided to create the email.")
	errNoRecipients   = errors.New("An email requires recipients to complete the transaction.")
	errNoSubject      = errors.New("Subject is a required field for an email.")
	errNoSender       = errors.New("An email requires a sender address to complete the transaction.")
	errNoTextTemplate = errors.New("A text template must be provided to create the text.")
)

type Email struct {
	Subject string
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Html    string
	Text    string
	Context interface{}
}

// RenderHtml returns a *string containing the rendered html template.
// Utilizes the Email.Html string referencing a template '.html' file.
func (e Email) RenderHtml() (*string, error) {
	return RenderTemplate(e.Html, e.Context)
}

// RenderText returns a *string containing the rendered html template.
// Utilizes the Email.Html string referencing a template '.txt' file.
func (e Email) RenderText() (*string, error) {
	return RenderTemplate(e.Text, e.Context)
}

// IsValid returns an error if the the relevant Email object violates any of the constraints
func (e Email) IsValid() error {
	if e.Subject == "" {
		return errNoSubject
	}

	if e.Sender == "" {
		return errNoSender
	}

	if len(e.To) == 0 {
		return errNoRecipients
	}

	if e.Html == "" {
		return errNoHtmlTemplate
	}

	if e.Text == "" {
		return errNoTextTemplate
	}

	return nil
}

// NewEmail returns an Email struct with the following arguments
// subject: The subject title of the email.
// sender: Email address of the sending email account.
// htmlFile: Filepath to the .html template to be rendered.
// textFile: Filepath to the .txt template to rendered.
// recipients: An array of email address for the email to be delivered to.
// context: an object that maps the templates keywords with their intended values.
func NewEmail(subject, sender, htmlFile, textFile string, to, cc, bcc []string, context interface{}) (*Email, error) {
	email := Email{
		Subject: subject,
		Sender:  sender,
		To:      to,
		Cc:      cc,
		Bcc:     bcc,
		Html:    htmlFile,
		Text:    textFile,
		Context: context,
	}

	err := email.IsValid()
	if err != nil {
		return nil, err
	}

	return &email, nil
}

// RenderTemplate renders the template file provided
// Context is interpretted from the context interface argument
// Currently all data associated with the interface must be relevant to the template.
// Returns a pointer to a bytes.Buffer interface and an error
func RenderTemplate(src string, context interface{}) (*string, error) {
	temp, err := template.ParseFiles(src)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	err = temp.Execute(&out, context)
	if err != nil {
		return nil, err
	}

	s := out.String()

	return &s, nil
}
