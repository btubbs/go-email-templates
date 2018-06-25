// Package emailtemplates provides a tool for reading txt and html files from the filesystem and
// turning them into .go source files that can be compiled into your application.

package emailtemplates

import (
	"bytes"
	"go/format"
	"io"
	"strconv"
	"text/template"
)

type RawTemplate struct {
	Name    string
	Subject string
	Text    string
	HTML    string
}

type TemplatesFile struct {
	PackageName string
	TemplateDir string
	File        string
	Templates   []RawTemplate
}

func (tf TemplatesFile) Render(w io.Writer) error {
	// make a "quote" function available inside our Go source template, so it can escape quote marks
	// inside the email template content.
	funcs := template.FuncMap{
		"quote": strconv.Quote,
	}
	t, err := template.New("").Funcs(funcs).Parse(srcTmpl)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = t.Execute(&b, tf)
	if err != nil {
		return err
	}
	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	_, err = w.Write(formatted)
	return err
}
