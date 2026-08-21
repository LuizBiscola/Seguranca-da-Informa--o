package main

import (
	"ApiLogin/Internal/database"
	"ApiLogin/Internal/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	db, err := database.Connect("root:senha@tcp(127.0.0.1:3306)/meuapp")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	userHandler := handlers.NewUserHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /cadastro", userHandler.Cadastro)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("GET /usuarios", userHandler.BuscarUsuarios)

	handler := c.Handler(mux)

	log.Println("servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
