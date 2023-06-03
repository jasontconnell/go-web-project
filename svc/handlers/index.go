package handlers

import (
	"net/http"

	"github.com/jasontconnell/go-web-project/data"
)

func (s *SiteHandler) Index(w http.ResponseWriter, r *http.Request) {
	exec(s.Template(), w, "Main", nil)
}

func (s *SiteHandler) ApiTest(w http.ResponseWriter, r *http.Request) {
	pkg := s.repo.GetPackage()
	if pkg.Message == "" {
		pkg = data.Package{Message: "Hello World! from database with id "}
		s.repo.SavePackage(&pkg)
	}
	sendJson(w, pkg)
}
