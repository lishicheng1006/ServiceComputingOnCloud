package main

import (
	"dailyProject/RESTapi/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func checkAuth(w http.ResponseWriter, req *http.Request) bool {
	vars := mux.Vars(req)
	username, password, ok := req.BasicAuth()
	if ok {
		if err := storage.CheckPwd(username, password); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
			return false
		}
		if username != vars["username"] {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", "This user does not have permission")))
			return false
		}
		return true
	}
	w.WriteHeader(http.StatusNonAuthoritativeInfo)
	return false
}

func userRegist(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}
	var userInfo data.User
	if err := json.Unmarshal(body, &userInfo); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}

	if err := storage.UserRegist(userInfo.Username, userInfo.Password); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func listUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := struct {
		Usernames []string
	}{Usernames: storage.ListUser()}

	json.NewEncoder(w).Encode(data)
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if checkAuth(w, req) {
		fmt.Println(vars["username"])
		storage.DeleteUser(vars["username"])
		w.WriteHeader(http.StatusOK)
	}
}

func updateUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}

	data := struct {
		Password string
	}{}
	if err := json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err)))

		return
	}
	fmt.Println(data)
	if checkAuth(w, req) {
		storage.UpdatePassword(vars["username"], data.Password)
		w.WriteHeader(http.StatusNoContent)
	}
}
