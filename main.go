package main

import (
	"log"
	"net/http"

	"adoptme/config"
	"adoptme/handlers"
	"adoptme/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectToDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handlers.Signup)
	mux.HandleFunc("/login", handlers.Login)
	routes.RegisterPetRoutes(mux)

	log.Println("✅ Supabase connection successful")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
