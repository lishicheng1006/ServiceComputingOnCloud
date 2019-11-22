package main

import (
	"dailyProject/RESTapi/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func checkBlogAuth(w http.ResponseWriter, req *http.Request, authUsername bool) bool {
	username, password, ok := req.BasicAuth()
	if ok {
		if err := storage.CheckPwd(username, password); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
			return false
		}
		if authUsername {
			vars := mux.Vars(req)
			id, _ := strconv.Atoi(vars["blogID"])
			if username != storage.GetUserByBlogid(id) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", "This user does not have permission")))
				return false
			}
		}
		return true
	}
	w.WriteHeader(http.StatusNonAuthoritativeInfo)
	return false
}

func createBlog(w http.ResponseWriter, req *http.Request) {
	if !checkBlogAuth(w, req, false) {
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}
	var blogInfo data.BlogPost
	if err := json.Unmarshal(body, &blogInfo); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}

	username, _, _ := req.BasicAuth()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"blogID\":\"%d\"}", storage.CreateBlog(blogInfo, username))))
}

func getOnesBlog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(req)
	data := storage.ListOnesBlog(vars["username"])

	json.NewEncoder(w).Encode(data)

}

func updateBlog(w http.ResponseWriter, req *http.Request) {
	if !checkBlogAuth(w, req, true) {
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}
	var blogInfo data.BlogPost
	if err := json.Unmarshal(body, &blogInfo); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["blogID"])
	if err := storage.UpdateBlog(id, blogInfo); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteBlog(w http.ResponseWriter, req *http.Request) {
	if !checkBlogAuth(w, req, true) {
		return
	}
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["blogID"])
	if err := storage.DeleteBlog(id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getBlog(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["blogID"])
	data, err := storage.GetBlog(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
