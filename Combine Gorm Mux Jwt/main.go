package main

import (
	"Combine-Gorm-Mux-Jwt/controllers"
	"Combine-Gorm-Mux-Jwt/middlewares"
	"Combine-Gorm-Mux-Jwt/users"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
	users.ConnectDBUsers()
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")


	r.HandleFunc("/products", middlewares.MiddlewareJWTAuthorization(controllers.ReadAllData)).Methods("GET")
	r.HandleFunc("/product/{id}", middlewares.MiddlewareJWTAuthorization(controllers.ReadProductById)).Methods("GET")
	r.HandleFunc("/addProduct", middlewares.MiddlewareJWTAuthorization(controllers.AddNewProduct)).Methods("POST")
	r.HandleFunc("/updateProduct/{id}", middlewares.MiddlewareJWTAuthorization(controllers.EditProduct)).Methods("PUT")
	r.HandleFunc("/deleteProduct/{id}", middlewares.MiddlewareJWTAuthorization(controllers.DeleteProduct)).Methods("DELETE")

	http.ListenAndServe(":9000", r)

}