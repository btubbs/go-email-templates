package emailtemplates

const srcTmpl = `// Generated code. DO NOT EDIT.
package {{.PackageName}}

import emailtemplates "github.com/btubbs/go-email-templates"

{{range .Templates}}var {{.Name}} = emailtemplates.Template{
	Subject: {{quote .Subject}},
	Text: {{quote .Text}},
	HTML: {{quote .HTML}},
}

{{end}}
`
