package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	CreatedDate time.Time
	Id          uint32
	LastVisit   time.Time
	PassWord    string
	Role        string
	UpdatedAt   time.Time
	UserName    string
}

func NewUser(role, userName, Password string) (*User, error) {
	user := &User{
		CreatedDate: time.Now(),
		Role:        role,
		UserName:    userName,
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
