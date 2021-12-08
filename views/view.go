package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir         string = "views/layouts/"
	TemplateExtension string = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExtension)
	if err != nil {
		panic(err)
	}
	return files
}

// ... variadic parameter, takes any number of strings
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...) // unpack the slice returned by layoutFiles()
	//
	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// Add Render method for the View struct
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
