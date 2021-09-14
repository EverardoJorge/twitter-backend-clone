package main

import (
	"log"

	"github.com/EverardoJorge/twitter-backend-clone/db"
	"github.com/EverardoJorge/twitter-backend-clone/handlers"
)


func main(){
	if db.CheckConnection() == 0 {
		log.Fatal("No connection")
		return
	}

	handlers.Handlers()
}
