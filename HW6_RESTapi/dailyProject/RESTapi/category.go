package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllCategory(w http.ResponseWriter, req *http.Request) {
	data := storage.GetAllCategories()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func getBlogsByCategory(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["categoryID"])

	data := storage.GetBlogsByCategory(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func createCategory(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}

	categoryInfo := struct {
		CategoryName string `json:"categoryName"`
	}{}
	if err := json.Unmarshal(body, &categoryInfo); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}

	categoryID, err := storage.CreateCategory(categoryInfo.CategoryName)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"CategoryID\":\"%v\"}", categoryID)))
}

func updateCategory(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}

	categoryInfo := struct {
		CategoryName string `json:"categoryName"`
	}{}
	if err := json.Unmarshal(body, &categoryInfo); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["categoryID"])
	if err := storage.UpdateCategory(id, categoryInfo.CategoryName); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteCategory(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["categoryID"])

	if err := storage.DeleteCategory(id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}
