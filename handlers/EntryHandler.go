package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/render"
    "github.com/russross/blackfriday"
    "github.com/kixpanganiban/g3/lib"
)

// Handles entries via "/blog/{name}"
// Loads the corresponding .md file using ReadContent, parses it via
// blackfriday then returns it as html through PageFactory.
func EntryHandler (c *gin.Context) {
    md, err := lib.ReadContent(c.Param("name"))
    if err != nil {
        c.String(404, "Entry not found")
    } else {
        html := blackfriday.MarkdownCommon(md)
        pf := lib.PageFactory{}
        pf.Init(html, "SomeTitle", "templates/base.tmpl")
        c.Render(
            200,
            render.Data {
            ContentType: "text/html",
            Data:        []byte(pf.Render()),
        })
    }
}
