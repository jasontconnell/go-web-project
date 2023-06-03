package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func sendJson(w http.ResponseWriter, s interface{}) error {
	enc := json.NewEncoder(w)
	return enc.Encode(s)
}

func sendError(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
	w.WriteHeader(http.StatusInternalServerError)
}

func exec(tmp *template.Template, w http.ResponseWriter, name string, data interface{}) {
	err := tmp.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println("error executing template", err)
	}
}
