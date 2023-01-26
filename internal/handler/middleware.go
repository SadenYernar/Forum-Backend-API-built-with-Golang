package handler

import (
	"context"
	"log"
	"net/http"
)

func (h *Handler) IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("session_name")
		if err != nil {
			http.Redirect(w, r, "/need-to-sign", http.StatusSeeOther)
			return
		}
		// по токену запрашиваем uuid пользователя
		uuid, err := h.service.GetSessionService(token.Value)
		if err != nil {
			log.Printf("Get session from handler don`t work %e", err)
			errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "uuid", uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (h *Handler) IfAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("session_name")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		uuid, err := h.service.GetSessionService(token.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), "uuid", uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
