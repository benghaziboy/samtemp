package samtemp

import (
	. "gopkg.in/check.v1"
	"testing"
)

var (
	defaultHtmlTemplate string
)

func Test(t *testing.T) { TestingT(t) }

type SamtempSuite struct{}

func (s *SamtempSuite) SetUpTest(c *C) {
	defaultHtmlTemplate = "test_template.html"
}

var _ = Suite(&SamtempSuite{})

func (s *SamtempSuite) TestRenderTemplate(c *C) {
	m := map[string]string{
		"FirstName": "Space",
		"LastName":  "Ghost",
	}

	_, err := RenderTemplate(defaultHtmlTemplate, m)
	c.Assert(err, IsNil)
}
