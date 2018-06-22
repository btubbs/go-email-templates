package emailtemplates

import (
	"bytes"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path"
)

const subjectFile = "subject.txt"
const txtFile = "content.txt"
const htmlFile = "content.html"

func WriteTemplatesToFile(fileName, templateDir, packageName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteTemplates(f, templateDir, packageName)
}

func WriteTemplates(w io.Writer, templateDir, packageName string) error {
	ts, _ := getProtoTemplatesFromDir(templateDir)
	tf := TemplatesFile{
		PackageName: packageName,
		Templates:   ts,
	}

	var b bytes.Buffer
	Render(tf, &b)
	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	_, err = w.Write(formatted)
	return err
}

func getProtoTemplatesFromDir(dir string) ([]ProtoTemplate, error) {
	templateNames, err := getSubdirectories(dir)
	out := []ProtoTemplate{}
	for _, tName := range templateNames {
		subj, err := ioutil.ReadFile(path.Join(dir, tName, subjectFile))
		if err != nil {
			return out, err
		}
		txt, err := ioutil.ReadFile(path.Join(dir, tName, txtFile))
		if err != nil {
			return out, err
		}
		html, err := ioutil.ReadFile(path.Join(dir, tName, htmlFile))
		if err != nil {
			return out, err
		}
		out = append(out, ProtoTemplate{
			Name:    tName,
			Subject: string(subj),
			Text:    string(txt),
			HTML:    string(html),
		})
	}
	return out, err
}

func getSubdirectories(dir string) ([]string, error) {
	var out []string
	files, err := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			out = append(out, f.Name())
		}
	}
	return out, err
}
