package entity

import "errors"

var InvalidUserEntity = errors.New("user info can not be empty")

var InvalidNewsEntity = errors.New("news info can not be empty")

var NotFound = errors.New("not found")
