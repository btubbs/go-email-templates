package emailtemplates

import (
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
	return WriteTemplates(f, path.Base(fileName), templateDir, packageName)
}

func WriteTemplates(w io.Writer, baseName, templateDir, packageName string) error {
	ts, err := getRawTemplatesFromDir(templateDir)
	if err != nil {
		return err
	}
	return TemplatesFile{
		PackageName: packageName,
		TemplateDir: templateDir,
		File:        baseName,
		Templates:   ts,
	}.Render(w)
}

func getRawTemplatesFromDir(dir string) ([]RawTemplate, error) {
	templateNames, err := getSubdirectories(dir)
	out := []RawTemplate{}
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
		out = append(out, RawTemplate{
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
