package main

import (
    "log"
    "os"
    "text/template"
)

type Note struct {
    Title	string
    Description string
}

const tmpl = "Notes are:\n{{range .}}\nTitle: {{.Title}}, Description: {{.Description}}\n{{end}}"

func main() {
    notes := []Note{
	    {"text/template", "Template generates textual output"},
	    {"html/template", "Template generates HTML output"},
    }
    t := template.New("note")
    t, err := t.Parse(tmpl)
    if err != nil {
	log.Fatal("Parse: ",err)
	return
    }
    if err := t.Execute(os.Stdout, notes); err != nil {
	log.Fatal("Parse: ",err)
	return
    }
}