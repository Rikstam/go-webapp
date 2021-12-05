package views

import "html/template"

// ... variadic parameter, takes any number of strings
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	//
	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
