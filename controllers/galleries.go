package controllers

import (
	"net/http"

	"lenslocked.com/views"
)

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/galleries"),
	}
}

type Galleries struct {
	NewView *views.View
}

// GET /galleries
// GET /signup
func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
