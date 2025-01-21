package handlers

import (
	"net/http"
)

func (s *SiteHandler) Index(w http.ResponseWriter, r *http.Request) {
	exec(s.Template(), w, "Main", nil)
}

func (s *SiteHandler) ApiTest(w http.ResponseWriter, r *http.Request) {

	sendJson(w, JsonSuccess{true, "Success"})
}
