package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

var (
	testSubject      = "Hello Darkness"
	testSender       = "Zboy"
	testHtmlTemplate = "test_template.html"
	testTextTemplate = "test_template.txt"

	testTo = []string{
		"charles@bronson.com",
		"sin@bad.com",
	}

	testContext = map[string]string{
		"FirstName": "Space",
		"LastName":  "Ghost",
	}
)

func Test(t *testing.T) { TestingT(t) }

type SamtempSuite struct{}

var _ = Suite(&SamtempSuite{})

func (s *SamtempSuite) TestRenderHtmlTemplate(c *C) {
	re, err := RenderTemplate(testHtmlTemplate, testContext)
	c.Assert(err, IsNil)
	c.Assert(*re, Equals, "<h1>Hi I am, Space Ghost</h1>\n")
}

func (s *SamtempSuite) TestRenderTextTemplate(c *C) {
	re, err := RenderTemplate(testTextTemplate, testContext)
	c.Assert(err, IsNil)
	c.Assert(*re, Equals, "Hi I am, Space Ghost\n")
}

func (s *SamtempSuite) TestNewEmail(c *C) {
	email, err := NewEmail(testSubject, testSender, testHtmlTemplate, testTextTemplate, testTo, nil, nil, testContext)
	c.Assert(err, IsNil)
	c.Assert(email.Subject, Equals, testSubject)
	c.Assert(email.Sender, Equals, testSender)
	c.Assert(email.Html, Equals, testHtmlTemplate)
	c.Assert(email.Text, Equals, testTextTemplate)
	c.Assert(email.To[0], Equals, testTo[0])
	c.Assert(email.To[1], Equals, testTo[1])

	context, ok := email.Context.(map[string]string)
	c.Assert(ok, Equals, true)
	c.Assert(context["FirstName"], Equals, testContext["FirstName"])
	c.Assert(context["LastName"], Equals, testContext["LastName"])
}

func (s *SamtempSuite) TestNewEmailNotValid(c *C) {
	email, err := NewEmail("", testSender, testHtmlTemplate, testTextTemplate, testTo, nil, nil, testContext)
	c.Assert(email, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoSubject)

	email, err = NewEmail(testSubject, "", testHtmlTemplate, testTextTemplate, testTo, nil, nil, testContext)
	c.Assert(email, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoSender)

	email, err = NewEmail(testSubject, testSender, "", testTextTemplate, testTo, nil, nil, testContext)
	c.Assert(email, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoHtmlTemplate)

	email, err = NewEmail(testSubject, testSender, testHtmlTemplate, "", testTo, nil, nil, testContext)
	c.Assert(email, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoTextTemplate)

	email, err = NewEmail(testSubject, testSender, testHtmlTemplate, testTextTemplate, nil, nil, nil, testContext)
	c.Assert(email, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err, Equals, errNoRecipients)
}
