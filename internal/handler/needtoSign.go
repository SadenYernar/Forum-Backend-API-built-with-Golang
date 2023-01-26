package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) needToSign(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/need-to-sign" {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	html, err := template.ParseFiles(TemplateDir + "html/needtoSign.html")
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusBadRequest)
	err = html.Execute(w, "")
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
