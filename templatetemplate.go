package emailtemplates

const srcTmpl = `// Generated code. DO NOT EDIT.
//go:generate eml ingest --templatedir {{.TemplateDir}} --packagename {{.PackageName}} --file {{.File}}
package {{.PackageName}}
 
type Template struct {
	Subject string
	Text string
	HTML string
}

{{range .Templates}}var {{.Name}} = Template{
	Subject: {{quote .Subject}},
	Text: {{quote .Text}},
	HTML: {{quote .HTML}},
}

{{end}}
`
