package repository

import (
	"github.com/ssuxue/mellow-server/domain/model"
	"github.com/ssuxue/mellow-server/internal/initdata"
)

type IUserRepository interface {
	Insert(member *model.User) int
	FindAll() []*model.User
	FindById(id int) *model.User
	FindByUsername(username string) *model.User
	FindByPhone(phone string) *model.User
	Delete(id int)
}

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Insert(member *model.User) int {
	create := initdata.DB.Create(&member)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (u *UserRepository) FindAll() []*model.User {
	var users []*model.User
	initdata.DB.Find(&users)
	return users
}

func (u *UserRepository) FindById(id int) *model.User {
	user := &model.User{}
	initdata.DB.Where("id = ?", id).Find(&user)
	return user
}

func (u *UserRepository) FindByUsername(username string) *model.User {
	user := &model.User{}
	initdata.DB.Where("username = ?", username).Find(&user)
	return user
}

func (u *UserRepository) FindByPhone(phone string) *model.User {
	user := &model.User{}
	initdata.DB.Where("phone = ?", phone).Find(&user)
	return user
}

func (u *UserRepository) Delete(id int) {
	user := &model.User{}
	initdata.DB.Where("id = ?", id).Delete(user)
}
