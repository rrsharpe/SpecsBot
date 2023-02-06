package ssd

import (
	"bytes"
	"text/template"
)

type Comment struct {
	Confident     bool
	Title         string
	Interface     string
	FormFactor    string
	Capacities    string
	Controller    string
	Configuration string
	DRAM          string
	HMB           string
	NANDBrand     string
	NANDType      string
	Layers        string
	ReadWrite     string
	Categories    string
	Notes         string
	AltNames      []string
}

func genComment(comment Comment) string {
	t := template.Must(template.ParseFiles("ssd/commentTemplate.md"))
	var buf bytes.Buffer
	t.Execute(&buf, comment)
	return buf.String()
}
