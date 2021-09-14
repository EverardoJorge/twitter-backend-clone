package routers

import (
	"encoding/json"
	"net/http"

	"github.com/EverardoJorge/twitter-backend-clone/db"
	"github.com/EverardoJorge/twitter-backend-clone/models"
)

func Register(w http.ResponseWriter, r *http.Request)  {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "There a error "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The email is required to register", 400)
		return
	}
	if len(t.Password) < 8 {
		http.Error(w, "The password is invalid", 400)
		return
	}
	_,found,_ := db.RegisteredUser(t.Email)
	if found == true {
		http.Error(w, "user already registered", 400)
		return
	}
	_,status,err := db.RegisterUser(t)
	if err != nil {
		http.Error(w, "there was an error registering the user "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "registration could not be completed", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}