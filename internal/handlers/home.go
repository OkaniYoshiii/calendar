package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/OkaniYoshiii/calendar/internal/calendar"
	"github.com/OkaniYoshiii/calendar/internal/repository"
)

type Anniversary struct {
	Childs []repository.Child
}

type HomeHandler struct {
	Queries  *repository.Queries
	Template *template.Template
}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	childs, err := handler.Queries.ListChilds(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cal := calendar.New(2025, func(day *calendar.Day[Anniversary]) {
		for _, child := range childs {
			if day.Valid() == true && child.Birthday.Month() == day.Month() && child.Birthday.Day() == day.Day() {
				day.Payload.Childs = append(day.Payload.Childs, child)
			}
		}
	})

	now := time.Now()

	data := struct {
		Childs   []repository.Child
		Calendar calendar.Calendar[Anniversary]
		Now      time.Time
	}{
		Childs:   childs,
		Calendar: cal,
		Now:      now,
	}

	if err := handler.Template.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
