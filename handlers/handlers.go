package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/EverardoJorge/twitter-backend-clone/middlewares"
	"github.com/EverardoJorge/twitter-backend-clone/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers()  {
	router := mux.NewRouter()

	router.HandleFunc("/resgiter", middlewares.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}