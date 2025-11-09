package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/OkaniYoshiii/calendar/internal/repository"
)

type HomeHandler struct {
	Queries  *repository.Queries
	Template *template.Template
}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	childs, err := handler.Queries.ListChilds(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	data := struct{ Childs []repository.Child }{
		Childs: childs,
	}

	if err := handler.Template.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
