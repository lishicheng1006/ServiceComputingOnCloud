package service

import (
	"dailyProject/Agenda/entity"
	"errors"
	"os"

	"github.com/olekukonko/tablewriter"
)

//Service is called by cobra cmd instruction
type Service struct {
	agendaStorage *entity.Storage
	curUser       entity.CurUser
}

var singleService *Service

//GetInstance returns the pointer of the instance of Service
func GetInstance() *Service {
	if singleService == nil {
		singleService = &Service{agendaStorage: entity.GetInstance()}
		singleService.agendaStorage.ReadCurUser(&singleService.curUser)
	}
	return singleService
}

//CreateUser is called by the cmd class to verify and create a user
func (s *Service) CreateUser(curUsername, curPassword, curEmail, curPhone string) error {
	if s.agendaStorage.QueryUser(func(username, password string) bool {
		if username == curUsername {
			return true
		}
		return false
	}) == nil {
		s.agendaStorage.CreateUser(entity.User{Username: curUsername, Password: curPassword, Email: curEmail, Phone: curPhone})
		return nil
	}

	return errors.New("The username is already occupied")
}

//UserLogin is called by the cmd class to check if a user can log in to the system
func (s *Service) UserLogin(curUsername, curPassword string) error {
	if err := s.ifLogin(); err == nil {
		return errors.New("Existing logged in user")
	}

	if s.agendaStorage.QueryUser(func(username, password string) bool {
		if username == curUsername && password == curPassword {
			return true
		}
		return false
	}) != nil {
		s.curUser = entity.CurUser{CurUsername: curUsername, CurPassword: curPassword}
		return s.agendaStorage.UpdateCurUser(&s.curUser)
	}

	return errors.New("Username or password is wrong")

}

//ListAllUsers is called by the cmd class to determine if users list can be shown
func (s *Service) ListAllUsers() error {
	if err := s.ifLogin(); err != nil {
		return err
	}

	allUser := s.agendaStorage.QueryUser(func(username, password string) bool {
		return true
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Username", "Email", "Phone"})

	for _, val := range allUser {
		table.Append([]string{val.Username, val.Email, val.Phone})
	}

	table.Render()
	return nil
}

//UserLogout is called by the cmd class to log out the system
func (s *Service) UserLogout() error {
	if err := s.ifLogin(); err != nil {
		return err
	}

	s.curUser = entity.CurUser{}
	return s.agendaStorage.UpdateCurUser(&s.curUser)
}

//DeleteUser is called by the cmd class to verify and delete current user
func (s *Service) DeleteUser() error {
	if err := s.ifLogin(); err != nil {
		return err
	}

	if err := s.agendaStorage.DeleteUser(func(username, password string) bool {
		if username == s.curUser.CurUsername {
			return true
		}
		return false
	}); err != nil {
		return nil
	}

	return s.UserLogout()
}

func (s *Service) ifLogin() error {
	if s.curUser.CurUsername == "" && s.curUser.CurPassword == "" {
		return errors.New("Not logged in. No permission")
	}
	return nil
}
