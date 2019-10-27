package entity

//User contains user information
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type CurUser struct {
	CurUsername string `json:"username"`
	CurPassword string `json:"password"`
}
