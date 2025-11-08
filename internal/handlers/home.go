package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/OkaniYoshiii/calendar/internal/repository"
)

type HomeHandler struct {
	Queries *repository.Queries
}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"timef": func(t time.Time, layoutStr string) string {
			var layout string
			switch layoutStr {
			case "DateOnly":
				layout = time.DateOnly
			default:
				layout = time.Layout
			}

			return t.Format(layout)
		},
	}

	tmpl := template.Must(
		template.New("base.html").Funcs(funcMap).ParseFiles(BASE_TEMPLATE, TEMPLATE_DIR+"/home/index.html"),
	)

	childs, err := handler.Queries.ListChilds(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	data := struct{ Childs []repository.Child }{
		Childs: childs,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
