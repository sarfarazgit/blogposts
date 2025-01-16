package blogrenderer

import (
	"html/template"
	"io"
)

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func RenderpostTemplate(writer io.Writer, post Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(writer, post); err != nil {
		return err
	}

	return nil
}

// if you're continuing from the read files chapter, you shouldn't redefine this
type Post struct {
	Title, Description, Body string
	Tags                     []string
}
