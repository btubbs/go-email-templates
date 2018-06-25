package emailtemplates

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	// make a TemplatesFile
	tf := TemplatesFile{
		PackageName: "foo",
		TemplateDir: "email_templates",
		File:        "my_email_templates.go",
		Templates: []RawTemplate{
			{
				Name:    "ForgotPassword",
				Subject: "Password Reset",
				Text:    "If this were a real password reset email, there would be a link here.",
				HTML:    `<body>This email content isn't "real".</body>`,
			},
		},
	}
	var b bytes.Buffer
	err := tf.Render(&b)
	assert.Nil(t, err)

	rendered := b.String()
	assert.Contains(t, rendered, "foo")
	assert.Contains(t, rendered, "Password Reset")
	assert.Contains(t, rendered, "var ForgotPassword")
	assert.Contains(t, rendered, "//go:generate eml ingest --templatedir email_templates --packagename foo --file my_email_templates.go")
}
