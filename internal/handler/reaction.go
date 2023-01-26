package handler

import (
	"forumv2/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
)

func (h *Handler) LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	uuidString := r.Context().Value("uuid")
	if uuidString == nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	uuid := uuidString.(uuid.UUID)
	postIDStr := r.FormValue("postID")
	if postIDStr == "" {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	comment := r.FormValue("commentID")
	if comment == "" {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	commentID, err := strconv.Atoi(comment)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	statusStr := r.PostFormValue("status")
	if statusStr == "" {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var status models.LikeStatus
	statusInt, err := strconv.Atoi(statusStr)
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	switch statusInt {
	case 1:
		status = models.Like
	case -1:
		status = models.DisLike
	default:
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return

	}

	like := models.LikeComment{
		UserID:     uuid,
		CommentsID: commentID,
		Status:     status,
	}

	err = h.service.Reactions.LikeCommentService(like)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, PostAddress+postIDStr, http.StatusSeeOther)
}

func (h *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	postIDStr := r.PostFormValue("postID")
	if postIDStr == "" {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	statusStr := r.FormValue("status")
	if statusStr == "" {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value("uuid").(uuid.UUID)
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	statusInt, err := strconv.Atoi(statusStr)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var status models.LikeStatus
	switch statusInt {
	case 1:
		status = models.Like
	case -1:
		status = models.DisLike
	default:
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	like := models.LikePost{
		PostID: postID,
		UserID: userID,
		Status: status,
	}
	err = h.service.LikePostService(like)
	if err != nil {
		log.Print(err)
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, PostAddress+postIDStr, http.StatusSeeOther)
}
