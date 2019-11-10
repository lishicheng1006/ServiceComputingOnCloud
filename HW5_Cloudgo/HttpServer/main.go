package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//User is used to record form data
type User struct {
	username, studentID, password, phone, email string
}

var user User

func dealPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in")
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		username := r.FormValue("username")
		studentID := r.FormValue("studentID")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		user = User{username: username, password: password, studentID: studentID, phone: phone, email: email}

		fmt.Fprintf(w, "Yes")
	}
}

func sendData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Fprintf(w, user.username+","+user.studentID+","+user.phone+","+user.email)
	}
}

func dealUnknown(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "This feature is under development...")
}

func dealForm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		t := template.Must(template.New("form").Parse(templateStr))

		items := struct {
			Username, StudentID, Phone, Email string
		}{Username: r.FormValue("username"), StudentID: r.FormValue("studentID"), Phone: r.FormValue("phone"), Email: r.FormValue("email")}
		fmt.Println(items)
		t.Execute(w, items)
	}
}

func main() {
	http.Handle("/submit/", http.HandlerFunc(dealPost))
	http.Handle("/loaddata/", http.HandlerFunc(sendData))
	http.Handle("/unknown/", http.HandlerFunc(dealUnknown))
	http.Handle("/form/", http.HandlerFunc(dealForm))
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
}

const templateStr = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Form</title>
	</head>

	<body>
		<div>
			<table border="1">
                <tr>
                  <th>username</th>
                  <th>studentID</th>
                  <th>phone</td>
                  <th>email</td>
                </tr>
                <tr>
                  <td>{{.Username}}</td>
                  <td>{{.StudentID}}</td>
                  <td>{{.Phone}}</td>
                  <td>{{.Email}}</td>
                </tr>
              </table>
		</div>
	</body>
</html>
`
