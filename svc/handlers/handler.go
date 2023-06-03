package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jasontconnell/go-web-project/conf"
	"github.com/jasontconnell/go-web-project/data"
	"github.com/jasontconnell/repository/v2"
)

type SiteHandler struct {
	http.Handler
	tmpl *template.Template
	repo data.PackageRepo
}

func (s *SiteHandler) Template() *template.Template {
	if s.tmpl != nil {
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

func GetHandler(cfg conf.Config) *SiteHandler {
	h := new(SiteHandler)

	var repo *data.PackageMongoRepo
	repo = new(data.PackageMongoRepo)
	repo.Initialize(repository.Configuration{ConnectionString: cfg.ConnectionString, Database: cfg.Database})
	h.repo = repo

	m := mux.NewRouter()
	m.StrictSlash(true)

	// posts
	m.HandleFunc("/", h.Index)
	m.HandleFunc("/api/test", h.ApiTest)
	h.Handler = m
	return h
}
