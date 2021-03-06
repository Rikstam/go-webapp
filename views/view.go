package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir         string = "views/layouts/"
	TemplateDir       string = "views/"
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
	addTemplatePath(files)
	addTemplateExtension(files)

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
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExtension(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExtension
	}
}
