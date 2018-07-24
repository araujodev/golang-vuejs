package session

import (
	Users "github.com/araujodev/golang-vuejs/pkg/types/users"
	ORM "github.com/araujodev/golang-vuejs/src/system/db"
	"github.com/araujodev/golang-vuejs/src/system/jwt"
	Passwords "github.com/araujodev/golang-vuejs/src/system/passwords"

	"encoding/json"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	if len(email) < 1 || len(password) < 1 {
		http.Error(w, "Email and password are required.", http.StatusUnauthorized)
		return
	}

	user := Users.User{Email: email}
	err := ORM.FindBy(db, &user)
	if err != nil || user.Id < 1 {
		log.Println(err)
		http.Error(w, "Credentials do not match.", http.StatusUnauthorized)
		return
	}

	if !Passwords.IsValid(user.Password, password) {
		http.Error(w, "Credentials do not match.", http.StatusUnauthorized)
		return
	}

	token := jwt.GetToken(user.Id)
	login := LoginData{User: user, Token: token}

	http.SetCookie(w, &http.Cookie{
		Name:       "api.example.com",
		Value:      token,
		Path:       "/",
		RawExpires: "0",
	})

	packet, err := json.Marshal(login)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to marshal json.", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
