package middlewares

import (
	"net/http"

	"github.com/EverardoJorge/twitter-backend-clone/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if db.CheckConnection() == 0 {
			http.Error(w, "No Connection on DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}