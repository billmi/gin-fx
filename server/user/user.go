package user

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	ID   int
}

func NewUserService(userRepo IUserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

type UserService struct {
	UserRepo IUserRepository
}

type IUserRepository interface {
	GetUserByID(id int) (*User, error)
	PrintInfo()
}

type UserRepoImpl struct {}

func (u *UserRepoImpl) GetUserByID(id int) (*User, error) {
	if id == 1 {
		return &User{Name: "Tom",ID: 1}, nil
	}
	return nil, errors.New("未找到相关数据")
}

func (u *UserRepoImpl)PrintInfo(){
	fmt.Println("do do...")
}

func NewUserRepo() *UserRepoImpl {
	return &UserRepoImpl{}
}

func GetUserService() *UserService {
	userRepoImpl := NewUserRepo()
	userService := NewUserService(userRepoImpl)
	return userService
}

