package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type SiteHandler struct {
	http.Handler
	tmpl    *template.Template
	devmode bool
}

func (s *SiteHandler) Template() *template.Template {
	if s.tmpl != nil && !s.devmode {
		return s.tmpl
	}

	var err error
	t := template.New("Site")
	t.Funcs(template.FuncMap{"html": toHtml})
	t, err = t.ParseGlob("./template/*.html")
	if err != nil {
		log.Fatal("couldn't load templates", err)
		panic(err)
	}

	s.tmpl = t
	return s.tmpl
}

func GetHandler(devmode bool) *SiteHandler {
	h := new(SiteHandler)

	h.devmode = devmode

	m := http.NewServeMux()

	// posts
	m.HandleFunc("GET /", h.Index)
	m.HandleFunc("GET /api/test", h.ApiTest)
	h.Handler = m
	return h
}
