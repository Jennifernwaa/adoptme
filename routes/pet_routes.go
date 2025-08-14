package routes

import (
	"adoptme/handlers"
	"adoptme/middleware"
	"net/http"
)

func RegisterPetRoutes(mux *http.ServeMux) {
	mux.Handle("GET /pets", http.HandlerFunc(handlers.GetPets))
	mux.Handle("GET /pets/{id}", http.HandlerFunc(handlers.GetPet))

	mux.Handle("POST /pets", middleware.AuthMiddleware(middleware.AdminOnly(http.HandlerFunc(handlers.CreatePet))))
	mux.Handle("PUT /pets/{id}", middleware.AuthMiddleware(middleware.AdminOnly(http.HandlerFunc(handlers.UpdatePet))))
	mux.Handle("DELETE /pets/{id}", middleware.AuthMiddleware(middleware.AdminOnly(http.HandlerFunc(handlers.DeletePet))))
}
