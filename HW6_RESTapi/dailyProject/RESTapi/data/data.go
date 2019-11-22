package data

import (
	"errors"
	"fmt"
)

type Data struct {
	User          map[string]string
	Blog          map[int]BlogPost
	Blog2user     map[int]string
	User2blog     map[string][]int
	Category      map[string]int
	CategoryInv   map[int]string
	Blog2category map[int][]int
	Category2blog map[int][]int
}

var singleInstance *Data
var blogCount = 0
var categoryCount = 0

func GetInstance() *Data {
	if singleInstance == nil {
		singleInstance = &Data{User: map[string]string{"admin": "admin"}, Blog: make(map[int]BlogPost), Blog2user: make(map[int]string), User2blog: make(map[string][]int), Category: make(map[string]int), CategoryInv: make(map[int]string), Blog2category: make(map[int][]int), Category2blog: make(map[int][]int)}
	}
	return singleInstance
}

func (s *Data) UserRegist(username, password string) error {
	if _, ok := s.User[username]; ok {
		return errors.New("User already exists")
	}
	s.User[username] = password
	return nil
}

func (s *Data) CheckPwd(username, password string) error {
	pwd, ok := s.User[username]
	if !ok {
		return errors.New("User does not exist")
	}
	if pwd == password {
		return nil
	}
	return errors.New("Wrong password")
}

func (s *Data) UpdatePassword(username, password string) error {
	s.User[username] = password
	return nil
}

func (s *Data) DeleteUser(username string) error {
	delete(s.User, username)
	return nil
}

func (s *Data) ListUser() []string {
	var usernames []string
	for key := range s.User {
		usernames = append(usernames, key)
	}
	return usernames
}

func (s *Data) CreateBlog(blog BlogPost, username string) int {
	s.Blog[blogCount] = blog
	s.Blog2user[blogCount] = username
	s.User2blog[username] = append(s.User2blog[username], blogCount)

	for _, val := range blog.Categories {
		if ans, ok := s.Category[val]; ok {
			s.Category2blog[ans] = append(s.Category2blog[ans], blogCount)
			s.Blog2category[blogCount] = append(s.Blog2category[blogCount], ans)
		} else {
			s.Category[val] = categoryCount
			s.CategoryInv[categoryCount] = val
			s.Category2blog[categoryCount] = append(s.Category2blog[categoryCount], blogCount)
			s.Blog2category[blogCount] = append(s.Blog2category[blogCount], categoryCount)
			categoryCount++
		}
	}
	blogCount++
	return blogCount - 1
}

func (s *Data) ListOnesBlog(username string) []BlogDigest {
	var blogDigest []BlogDigest
	fmt.Println(s.User2blog[username])
	for _, val := range s.User2blog[username] {
		blogDigest = append(blogDigest, BlogDigest{BlogID: val, Title: s.Blog[val].Title})
	}
	return blogDigest
}

func (s *Data) UpdateBlog(blogID int, blog BlogPost) error {
	if _, ok := s.Blog[blogID]; ok {
		s.Blog2category[blogID] = s.Blog2category[blogID][0:0]
		for _, val := range s.Blog[blogID].Categories {
			cateID := s.Category[val]
			var index int
			for key, val := range s.Category2blog[cateID] {
				if val == blogID {
					index = key
					break
				}
			}
			s.Category2blog[cateID] = append(s.Category2blog[cateID][:index], s.Category2blog[cateID][index+1:]...)
		}

		s.Blog[blogID] = blog

		for _, val := range blog.Categories {
			if ans, ok := s.Category[val]; ok {
				s.Category2blog[ans] = append(s.Category2blog[ans], blogID)
				s.Blog2category[blogID] = append(s.Blog2category[blogID], ans)
			} else {
				s.Category[val] = categoryCount
				s.CategoryInv[categoryCount] = val
				s.Category2blog[categoryCount] = append(s.Category2blog[categoryCount], blogID)
				s.Blog2category[blogID] = append(s.Blog2category[blogID], categoryCount)
				categoryCount++
			}
		}

		return nil
	}
	return errors.New("BlogID does not exist")
}

func (s *Data) GetBlog(blogID int) (BlogPost, error) {
	if _, ok := s.Blog[blogID]; ok {
		return s.Blog[blogID], nil
	}
	return BlogPost{}, errors.New("BlogID does not exist")
}

func (s *Data) DeleteBlog(blogID int) error {
	if _, ok := s.Blog[blogID]; ok {
		user := s.Blog2user[blogID]
		var index = 0
		for key, val := range s.User2blog[user] {
			if val == blogID {
				index = key
				break
			}
		}
		s.User2blog[user] = append(s.User2blog[user][:index], s.User2blog[user][index+1:]...)

		cateIDs := s.Blog2category[blogID]
		for _, val := range cateIDs {
			var index = 0
			for key, vals := range s.Category2blog[val] {
				if vals == blogID {
					index = key
					break
				}
			}
			s.Category2blog[val] = append(s.Category2blog[val][:index], s.Category2blog[val][index+1:]...)
		}

		delete(s.Blog2category, blogID)
		delete(s.Blog2user, blogID)
		delete(s.Blog, blogID)
		return nil
	}
	return errors.New("BlogID does not exist")
}

func (s *Data) GetUserByBlogid(blogID int) string {
	return s.Blog2user[blogID]
}

func (s *Data) GetAllCategories() []CategoryInfo {
	var categories []CategoryInfo
	for key, val := range s.Category {
		categories = append(categories, CategoryInfo{CategoryName: key, CategoryID: val})
	}
	return categories
}

func (s *Data) GetBlogsByCategory(CategoryID int) []BlogForCategory {
	var blogs []BlogForCategory
	fmt.Println(s.Category2blog)
	if val, ok := s.Category2blog[CategoryID]; ok {
		for _, value := range val {
			blogs = append(blogs, BlogForCategory{Title: s.Blog[value].Title, Content: s.Blog[value].Content, Username: s.Blog2user[value]})
		}
	}
	return blogs
}

func (s *Data) CreateCategory(CategoryName string) (int, error) {
	if _, ok := s.Category[CategoryName]; ok {
		return 0, errors.New("This category already exists")
	}
	s.Category[CategoryName] = categoryCount
	s.CategoryInv[categoryCount] = CategoryName
	categoryCount++
	return categoryCount - 1, nil
}

func (s *Data) UpdateCategory(CategoryID int, CategoryName string) error {
	if _, ok := s.CategoryInv[CategoryID]; ok {
		if _, ok := s.Category[CategoryName]; !ok {
			delete(s.Category, s.CategoryInv[CategoryID])
			s.CategoryInv[CategoryID] = CategoryName
			s.Category[CategoryName] = CategoryID
			return nil
		} else {
			return errors.New("This category name already exists")
		}
	}
	return errors.New("This category does not exist")
}

func (s *Data) DeleteCategory(CategoryID int) error {
	if _, ok := s.CategoryInv[CategoryID]; !ok {
		return errors.New("This category does not exist")
	}
	if len(s.Category2blog[CategoryID]) > 0 {
		return errors.New("There are blogs in this category, not deleteable")
	}
	delete(s.Category, s.CategoryInv[CategoryID])
	delete(s.CategoryInv, CategoryID)
	return nil
}
