package handler

import (
	"forumv2/internal/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

type Data struct {
	Userinfo models.User
	Post     []models.Post
	LikePost []models.Post
}

func (h *Handler) myprofile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/myprofile" {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	html, err := template.ParseFiles(TemplateDir + "html/myprofile.html")
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	uuidCtx := r.Context().Value("uuid")
	uuid := uuidCtx.(uuid.UUID)

	switch r.Method {
	case http.MethodGet:

		userInfo, err := h.service.GetUsersInfoByUUIDService(uuid)
		if err != nil {
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		usersPost, err := h.service.GetUsersPostInService(uuid)
		if err != nil {
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		userLikePosts, err := h.service.GetUserLikePostsInService(uuid)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := &Data{
			Userinfo: userInfo,
			Post:     usersPost,
			LikePost: userLikePosts,
		}

		err = html.Execute(w, data)
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

	default:
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
