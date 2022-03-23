package entity

import (
	"github.com/kasrasaeed/news_app/pkg/id"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	CreatedDate time.Time
	Id          id.UUID
	LastVisit   time.Time
	PassWord    string
	Role        string
	UserName    string
	UpdatedAt   time.Time
}

func NewUser(role, userName, Password string) (*User, error) {
	user := &User{
		Id:          id.NewUUID(),
		CreatedDate: time.Now(),
		Role:        role,
		UserName:    userName,
		LastVisit:   time.Now(),
	}
	pwd, err := hashPassWord(Password)
	if err != nil {
		return nil, err
	}
	user.PassWord = pwd
	err = user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateLastVisit(u *User) error {
	updatedUser := u
	updatedUser.LastVisit = time.Now()
	err := updatedUser.Validate()
	if err != nil {
		return InvalidUserEntity
	}
	return nil
}

func (u *User) Validate() error {
	if u.UserName == "" || u.PassWord == "" {
		return InvalidUserEntity
	}
	return nil
}

func hashPassWord(rawPass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPass), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
