package handler

import (
	"forumv2/internal/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

// // авторизация
func (h *Handler) userSignIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-in" {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		html, err := template.ParseFiles(TemplateDir + "html/signIn.html")
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		html.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		email, ok := r.Form["email"]
		if !ok {
			errorHeader(w, "email is not found", http.StatusBadRequest)
			return
		}
		password, ok := r.Form["password"]
		if !ok {
			errorHeader(w, "password is incorrect", http.StatusBadRequest)
			return
		}
		data := models.User{
			Email:    email[0],
			Password: password[0],
		}
		token, err := h.service.AuthorizationUserService(data)
		if err != nil {
			errorHeader(w, "Email or Password is incorrect", http.StatusBadRequest)
			return
		}
		cookie := http.Cookie{
			Name:   "session_name",
			Value:  token,
			MaxAge: 300 * 50,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

// регистрация
func (h *Handler) userSignUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-up" {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		html, err := template.ParseFiles(TemplateDir + "html/signUp.html")
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		html.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		username, ok := r.Form["username"]
		if !ok {
			errorHeader(w, "username is indicated incorrectly", http.StatusBadRequest)
			return
		}
		name, ok := r.Form["name"]
		if !ok {
			errorHeader(w, "name is indicated incorrectly", http.StatusBadRequest)
			return
		}
		password, ok := r.Form["password"]
		if !ok {
			errorHeader(w, "password is indicated incorrectly", http.StatusBadRequest)
			return
		}
		email, ok := r.Form["email"]
		if !ok {
			errorHeader(w, "email is indicated incorrectly", http.StatusBadRequest)
			return
		}
		data := models.User{
			Name:     name[0],
			Username: username[0],
			Password: password[0],
			Email:    email[0],
		}
		EmailExist, err := h.service.CheckUserEmail(email[0])
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if EmailExist {
			errorHeader(w, "user with this email already exists", http.StatusBadRequest)
			return
		}
		UsernameExist, err := h.service.CheckUserUsername(username[0])
		if err != nil {
			log.Print(err)
			errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if UsernameExist {
			errorHeader(w, "user with this username already exists", http.StatusBadRequest)
			return
		}
		status, err := h.service.User.CreateUserService(data)
		if err != nil {
			errorHeader(w, "One field is filled out incorrectly", status)
			// http.Error(w, http.StatusText(status), status)
			log.Printf("User not created")
			return
		}
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
	default:
		errorHeader(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) logOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		errorHeader(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	uuidCtx := r.Context().Value("uuid")
	if uuidCtx == nil {
		errorHeader(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	uuidStmt := uuidCtx.(uuid.UUID)

	cookie := http.Cookie{
		Name:   "session_name",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	err := h.service.DeleteSessionService(uuidStmt)
	if err != nil {
		errorHeader(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
