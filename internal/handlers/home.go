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

type HomeHandler struct {
	Queries  *repository.Queries
	Template *template.Template
}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	childs, err := handler.Queries.ListChilds(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cal := calendar.New(2025)
	now := time.Now()

	data := struct {
		Childs   []repository.Child
		Calendar calendar.Calendar
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
