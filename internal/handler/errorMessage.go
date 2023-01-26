package handler

import (
	"html/template"
	"net/http"
)

type ErrorBody struct {
	Status         int
	Message        string
	SpecialMessage string
}

func errorHeader(w http.ResponseWriter, errorMessage string, status int) {
	w.WriteHeader(status)
	errH := setError(status, errorMessage)
	html, err := template.ParseFiles(TemplateDir + "html/error.html")
	if err != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	html.Execute(w, errH)
}

func setError(status int, errorMessage string) *ErrorBody {
	return &ErrorBody{
		Status:         status,
		Message:        http.StatusText(status),
		SpecialMessage: errorMessage,
	}
}
