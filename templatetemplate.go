package emailtemplates

const srcTmpl = `// Generated code. DO NOT EDIT.
//go:generate eml ingest --templatedir {{.TemplateDir}} --packagename {{.PackageName}} --file {{.File}}
package {{.PackageName}}

import (
	"bytes"
	"text/template"
)

type Template struct {
	Subject       string
	Text          string
	HTML          string
	parsedSubject *template.Template
	parsedText    *template.Template
	parsedHTML    *template.Template
}

func (t *Template) Render(data interface{}) (RenderedTemplate, error) {
	out := RenderedTemplate{}
	var b bytes.Buffer
	err := t.parsedSubject.Execute(&b, data)
	if err != nil {
		return out, err
	}
	out.Subject = b.String()
	b.Reset()
	err = t.parsedText.Execute(&b, data)
	if err != nil {
		return out, err
	}
	out.Text = b.String()
	b.Reset()
	err = t.parsedHTML.Execute(&b, data)
	if err != nil {
		return out, err
	}
	out.HTML = b.String()
	return out, nil
}

type RenderedTemplate struct {
	Subject string
	Text    string
	HTML    string
}

{{range .Templates}}var {{.Name}} = Template{
	Subject: {{quote .Subject}},
	Text: {{quote .Text}},
	HTML: {{quote .HTML}},
}

{{end}}

func init() {
{{- range .Templates}}
	{{.Name}}.parsedSubject = template.Must(template.New("").Parse({{.Name}}.Subject))
	{{.Name}}.parsedText = template.Must(template.New("").Parse({{.Name}}.Text))
	{{.Name}}.parsedHTML = template.Must(template.New("").Parse({{.Name}}.HTML))
{{- end}}
}
`
