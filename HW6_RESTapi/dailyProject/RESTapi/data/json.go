package data

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BlogPost struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}

type BlogDigest struct {
	BlogID int    `json:"blogID"`
	Title  string `json:"blogTitle"`
}

type BlogForCategory struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

type CategoryInfo struct {
	CategoryName string `json:"CategoryName"`
	CategoryID   int    `json:"CategoryID"`
}
