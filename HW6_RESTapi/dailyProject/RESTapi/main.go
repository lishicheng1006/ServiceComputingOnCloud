package main

import (
	"dailyProject/RESTapi/data"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var storage = data.GetInstance()

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	router.HandleFunc("/api/users", userRegist).Methods("POST")
	router.HandleFunc("/api/users", listUsers).Methods("GET")
	router.HandleFunc("/api/users/{username}", updateUser).Methods("PUT")
	router.HandleFunc("/api/users/{username}", deleteUser).Methods("DELETE")

	router.HandleFunc("/api/blogs", createBlog).Methods("POST")
	router.HandleFunc("/api/blogs/user/{username}", getOnesBlog).Methods("GET")
	router.HandleFunc("/api/blogs/{blogID}", updateBlog).Methods("PUT")
	router.HandleFunc("/api/blogs/{blogID}", deleteBlog).Methods("DELETE")
	router.HandleFunc("/api/blogs/{blogID}", getBlog).Methods("GET")

	router.HandleFunc("/api/categories", getAllCategory).Methods("GET")
	router.HandleFunc("/api/categories", createCategory).Methods("POST")
	router.HandleFunc("/api/categories/{categoryID}", getBlogsByCategory).Methods("GET")
	router.HandleFunc("/api/categories/{categoryID}", updateCategory).Methods("PUT")
	router.HandleFunc("/api/categories/{categoryID}", deleteCategory).Methods("DELETE")

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	http.ListenAndServe(":7777", n)
}
