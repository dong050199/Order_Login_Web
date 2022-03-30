package main

import (
	c "SQLite_JWT/controller"
	"SQLite_JWT/driver"
	h "SQLite_JWT/handler"
	"SQLite_JWT/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func dosth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deo ai lam gi may")
}

func main() {
	driver.ConnectMongoDB()
	r := mux.NewRouter()
	r.Use(middleware.CORS)
	login := r.PathPrefix("/login").Subrouter()
	login.HandleFunc("/user", h.Login)
	login.HandleFunc("/register", h.Register)
	config := r.PathPrefix("/data").Subrouter()
	config.Use(middleware.Authentication)
	config.HandleFunc("/product", c.GetAllProduct).Methods("GET")
	config.HandleFunc("/product", c.CreateProduct).Methods("POST")
	config.HandleFunc("/product/{id}", c.GetProductById).Methods("GET")
	config.HandleFunc("/product/{id}", c.UpdateProduct).Methods("PUT")
	config.HandleFunc("/product/{id}", c.DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
