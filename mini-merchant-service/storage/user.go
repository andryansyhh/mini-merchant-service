package storage

import (
	"mini-merchant-service/entity"
	"mini-merchant-service/query"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDao interface {
	RegisterUser(user entity.Users) (entity.Users, error)
	FindUserByEmail(email string) (entity.Users, error)
	ShowAllUser() ([]entity.Users, error)
	FindUserByID(ID string) (entity.Users, error)
	UpdateUserByID(ID string, input entity.UpdateUserInputs) (entity.Users, error)
	DeleteUserByID(ID string) (string, error)
	CreateOutletUser(Outlet entity.Outlets) (entity.Outlets, error)
	FindOutletUserByID(ID string) (entity.Outlets, error)
	ShowAllOutletUser() ([]entity.Outlets, error)
}

type dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) RegisterUser(user entity.Users) (entity.Users, error) {
	qry := query.RegisterUser

	err := r.db.Raw(qry,
		user.UserID,
		user.FullName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) FindUserByEmail(email string) (entity.Users, error) {
	var user entity.Users
	qry := query.LoginUser

	err := r.db.Raw(qry, email).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) ShowAllUser() ([]entity.Users, error) {
	var user []entity.Users
	qry := query.GetAllUsers

	err := r.db.Raw(qry).Scan(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *dao) FindUserByID(ID string) (entity.Users, error) {
	var user entity.Users
	err := r.db.Where("id = ?", ID).Preload("Outlet").Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) UpdateUserByID(ID string, input entity.UpdateUserInputs) (entity.Users, error) {

	var user entity.Users
	genPassword, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err2 != nil {
		return user, err2
	}

	input.UpdatedAt = time.Now()

	qry := query.UpdateUserByID
	err := r.db.Raw(qry, input.FullName, input.Email, genPassword, input.UpdatedAt, ID).Scan(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) DeleteUserByID(ID string) (string, error) {
	user := &entity.Users{}
	qry := query.DeleteUserById

	err := r.db.Raw(qry, ID).Scan(&user).Error

	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *dao) CreateOutletUser(outlet entity.Outlets) (entity.Outlets, error) {
	qry := query.CreateOutletbyUser

	err := r.db.Raw(qry,
		outlet.OutletID,
		outlet.OutletName,
		outlet.Picture,
		outlet.UserID,
		outlet.CreatedAt,
		outlet.UpdatedAt).Scan(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) FindOutletUserByID(ID string) (entity.Outlets, error) {
	var outlet entity.Outlets

	qry := query.FindOutletUserByID

	err := r.db.Raw(qry, ID).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) ShowAllOutletUser() ([]entity.Outlets, error) {
	var outlet []entity.Outlets

	qry := query.GetAllOutlets

	err := r.db.Raw(qry).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
