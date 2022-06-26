package templateengines

import (
	"bytes"
	"encoding/json"
	"html/template"
)

const GoTemplateEngineKey = "golang"

type GoTemplateEngine struct {
}

func (gte *GoTemplateEngine) Execute(templateHtml *string, model any) (*string, error) {
	t, err := template.New("").
		Funcs(template.FuncMap{
			"marshal": func(v interface{}) template.JS {
				a, _ := json.Marshal(v)
				return template.JS(a)
			},
		}).
		Parse(*templateHtml)

	empty := ""
	if err != nil {
		return &empty, err
	}
	var buff bytes.Buffer

	if err := t.Execute(&buff, model); err != nil {
		return &empty, err
	}

	html := buff.String()

	return &html, nil
}

func (gte *GoTemplateEngine) Test(templateHtml *string, model any) error {
	t, err := template.New("").Option("missingkey=error").Parse(*templateHtml)
	if err != nil {
		return err
	}

	var buff bytes.Buffer

	if err := t.Execute(&buff, model); err != nil {
		return err
	}

	return nil
}
