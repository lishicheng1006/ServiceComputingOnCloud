package entity

import (
	"encoding/json"
	"os"
)

const userPath = "./entity/data/user.json"
const curUserPath = "./entity/data/curUser.txt"

//Storage responsible for reading and writing storage files
type Storage struct {
	userList []User
}

var singleStorage *Storage

//UserFilter is used to implement lambda expressions
type UserFilter func(username, password string) bool

//GetInstance returns the pointer of the instance of Storage
func GetInstance() *Storage {
	if singleStorage == nil {
		singleStorage = &Storage{}
		singleStorage.readFromFile(userPath, &singleStorage.userList)
	}
	return singleStorage
}

func (s *Storage) readFromFile(path string, v interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	return json.NewDecoder(jsonFile).Decode(v)
}

func (s *Storage) writeToFile(path string, v interface{}) error {

	jsonFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	return json.NewEncoder(jsonFile).Encode(v)
}

//CreateUser is called by the service class to create a user
func (s *Storage) CreateUser(u User) error {
	s.userList = append(s.userList, u)

	return s.UpdateUser()
}

//DeleteUser is called by the service class to delete a user
func (s *Storage) DeleteUser(uf UserFilter) error {

	for i, v := range s.userList {
		if uf(v.Username, v.Password) {
			s.userList = append(s.userList[:i], s.userList[i+1:]...)
			break
		}
	}
	return s.UpdateUser()
}

//QueryUser is called by the service class to find some users that meet the criteria
func (s *Storage) QueryUser(uf UserFilter) []User {
	var user []User
	for _, v := range s.userList {
		if uf(v.Username, v.Password) {
			user = append(user, v)
		}
	}
	return user
}

//UpdateUser is called by the service class to update the user json file
func (s *Storage) UpdateUser() error {
	return s.writeToFile(userPath, &s.userList)
}

//UpdateCurUser is called by the service class to update the current user text file
func (s *Storage) UpdateCurUser(cu *CurUser) error {
	return s.writeToFile(curUserPath, cu)
}

//ReadCurUser is called by the service class to read the current user from text file
func (s *Storage) ReadCurUser(cu *CurUser) error {
	return s.readFromFile(curUserPath, cu)
}
