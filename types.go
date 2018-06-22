package emailtemplates

import (
	"io"
	"strconv"
	"text/template"
)

// ProtoTemplate stores the raw template strings, as read from the txt/html files on the filesystem.
// A slice of these is included in the context fed into the template render call that generates the
// Go code.
type ProtoTemplate struct {
	Name    string
	Subject string
	Text    string
	HTML    string
}

type TemplatesFile struct {
	PackageName string
	Templates   []ProtoTemplate
}

func Render(tf TemplatesFile, w io.ReadWriter) error {
	funcs := template.FuncMap{
		"quote": strconv.Quote,
	}
	t, err := template.New("").Funcs(funcs).Parse(srcTmpl)
	if err != nil {
		return err
	}
	return t.Execute(w, tf)
}
