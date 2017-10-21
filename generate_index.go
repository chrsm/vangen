package main

import (
	"fmt"
	"html/template"
	"io"
)

func generate_index(w io.Writer, domain string, r []repository) error {
	const html = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
* { font-family: sans-serif; }
</style>
</head>
<body>
<h1>Go repositories hosted at {{$.Domain}}</h1>
<ul>
{{range $_, $r := .Repositories -}}
<li>
<a href="//{{$.Domain}}/{{$r.Prefix}}">{{$.Domain}}/{{$r.Prefix}}</a>
{{if .Subs -}}<ul>{{end -}}
{{range $_, $s := .Subs -}}<li><a href="//{{$.Domain}}/{{$r.Prefix}}/{{$s}}">{{$.Domain}}/{{$r.Prefix}}/{{$s}}</a></li>{{end -}}
{{if .Subs -}}</ul>{{end -}}
</li>
{{end -}}
</ul>
Generated by <a href="https://4d63.com/vangen">vangen</a>.
</body>
</html>`

	tmpl, err := template.New("").Parse(html)
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	data := struct {
		Domain       string
		Repositories []repository
	}{
		Domain:       domain,
		Repositories: r,
	}

	err = tmpl.ExecuteTemplate(w, "", data)
	if err != nil {
		return fmt.Errorf("generating template: %v", err)
	}

	return nil
}
