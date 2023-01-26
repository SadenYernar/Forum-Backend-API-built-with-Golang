package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

func (h *Handler) FilterByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	categoriesArr, ok := r.URL.Query()["category"]
	if !ok {
		errorHeader(w, "invalid category", http.StatusBadRequest)
		return
	}
	_, err := h.service.CreateCategory(categoriesArr)
	if err != nil {
		errorHeader(w, "invalid category", http.StatusBadRequest)
		return
	}
	posts, err := h.service.FilterPostsByCategories(categoriesArr)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(TemplateDir + "html/index.html")
	if err != nil {
		log.Print(err)
		errorHeader(w, "", http.StatusInternalServerError)
		return
	}
	uuidCtx := r.Context().Value("uuid")
	if uuidCtx == nil {
		res := AllData{
			Post: posts,
		}
		err = tmpl.Execute(w, &res)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}

	uuid := uuidCtx.(uuid.UUID)
	user, err := h.service.GetUsersInfoByUUIDService(uuid)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res := AllData{
		Data: user,
		Post: posts,
	}
	err = tmpl.Execute(w, &res)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
