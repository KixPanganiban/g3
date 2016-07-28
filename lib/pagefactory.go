package lib

import (
    "bytes"
    "html/template"
)

type PageFactory struct {
    Body template.HTML
    Title string
    TemplatePath string
}

// Initialize the PageFactory by converting the string HTML Body
// into a template.HTML type Body, passing other params as normal
func (p *PageFactory) Init(Body []byte, Title string, TemplatePath string) {
    p.Body = template.HTML(Body)
    p.Title = Title
    p.TemplatePath = TemplatePath
}

// Returns a string of the parsed template with the HTML Body
func (p *PageFactory) Render() string {
    var htmlBytes bytes.Buffer
    t, _ := template.ParseFiles(p.TemplatePath)
    t.Execute(&htmlBytes, p)
    return htmlBytes.String()
}
