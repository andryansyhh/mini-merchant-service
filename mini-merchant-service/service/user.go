package service

import (
	"errors"
	"fmt"
	"mini-merchant-service/storage"
	"mini-merchant-service/formatter"
	entity "mini-merchant-service/entity"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(userUser entity.UserInputs) (formatter.UserFormat, error)
	LoginUser(input entity.LoginUserInputs) (formatter.UserFormat, error)
	ShowAllUser() ([]formatter.UserFormat, error)
	FindUserByID(userID string) (entity.Users, error)
	UpdateUserByID(userID string, input entity.UpdateUserInputs) (formatter.UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	CreateOutletUser(outlet entity.OutletInput, userID string) (entity.Outlets, error)
	ShowAllOutletUser() ([]formatter.OutletFormat, error)
}

type userservice struct {
	dao storage.UserDao
}

func NewUserService(dao storage.UserDao) *userservice {
	return &userservice{dao}
}

func (s *userservice) RegisterUser(userUser entity.UserInputs) (formatter.UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(userUser.Password), bcrypt.MinCost)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if err != nil {
		return formatter.UserFormat{}, err
	}

	useruuid, err := uuid.NewV4()

	if err != nil {
		return formatter.UserFormat{}, err
	}

	var newUser = entity.Users{
		UserID:    useruuid.String(),
		FullName:  userUser.FullName,
		Email:     userUser.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.dao.RegisterUser(newUser)
	formatUser := formatter.FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *userservice) LoginUser(input entity.LoginUserInputs) (formatter.UserFormat, error) {
	userUser, err := s.dao.FindUserByEmail(input.Email)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if len(userUser.UserID) != 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(userUser.Password), []byte(input.Password)); err != nil {
			return formatter.UserFormat{}, errors.New("password invalid")
		}

		formatter := formatter.FormatUser(userUser)

		return formatter, nil
	}

	newError := fmt.Sprintf("user id not found")
	return formatter.UserFormat{}, errors.New(newError)
}

func (s *userservice) ShowAllUser() ([]formatter.UserFormat, error) {
	userUser, err := s.dao.ShowAllUser()
	var formatuserUser []formatter.UserFormat

	for _, users := range userUser {
		formatUser := formatter.FormatUser(users)
		formatuserUser = append(formatuserUser, formatUser)

	}
	if err != nil {
		return formatuserUser, err
	}

	return formatuserUser, nil
}

func (s *userservice) FindUserByID(userID string) (entity.Users, error) {
	user, err := s.dao.FindUserByID(userID)

	if err != nil {
		return entity.Users{}, err
	}

	if len(user.UserID) == 0 {
		newError := fmt.Sprintf("user id not found")
		return entity.Users{}, errors.New(newError)
	}

	return user, nil
}

func (s *userservice) UpdateUserByID(userID string, input entity.UpdateUserInputs) (formatter.UserFormat, error) {
	var dataUpdate = entity.UpdateUserInputs{}

	user, err := s.dao.FindUserByID(userID)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if len(user.UserID) == 0 {
		newError := fmt.Sprintf("user id not found")
		return formatter.UserFormat{}, errors.New(newError)
	}

	if input.FullName != "" || len(input.FullName) != 0 {
		dataUpdate.FullName = input.FullName
	}
	if input.Email != "" || len(input.Email) != 0 {
		dataUpdate.Email = input.Email
	}
	if input.Password != "" || len(input.Password) != 0 {
		dataUpdate.Password = input.Password
	}

	userUpdated, err := s.dao.UpdateUserByID(userID, dataUpdate)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	formatUser := formatter.FormatUser(userUpdated)

	return formatUser, nil
}

func (s *userservice) DeleteUserByID(userID string) (interface{}, error) {
	user, err := s.dao.FindUserByID(userID)

	if err != nil {
		return nil, err
	}

	if len(user.UserID) == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.dao.DeleteUserByID(userID)

	if err != nil {
		return nil, err
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := formatter.FormatDeleteUser(msg)

	return formatDelete, nil
}

func (s *userservice) CreateOutletUser(outlet entity.OutletInput, userID string) (entity.Outlets, error) {

	checkStatus, err := s.dao.FindOutletUserByID(userID)

	if err != nil {
		return checkStatus, err
	}

	if checkStatus.UserID == userID {
		errorStatus := fmt.Sprintf("Outlet for user id : %s has been created", userID)
		return checkStatus, errors.New(errorStatus)
	}

	outletuuid, err := uuid.NewV4()

	if err != nil {
		return entity.Outlets{}, err
	}

	var newOutlet = entity.Outlets{
		OutletID:   outletuuid.String(),
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserID:     userID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createOutlet, err := s.dao.CreateOutletUser(newOutlet)

	if err != nil {
		return createOutlet, err
	}

	return createOutlet, nil
}

func (s *userservice) ShowAllOutletUser() ([]formatter.OutletFormat, error) {
	outlet, err := s.dao.ShowAllOutletUser()
	var formatOutlet []formatter.OutletFormat

	for _, outlets := range outlet {
		formatOutlets := formatter.FormatOutlet(outlets)
		formatOutlet = append(formatOutlet, formatOutlets)

	}

	if err != nil {
		return formatOutlet, err
	}

	return formatOutlet, nil
}
