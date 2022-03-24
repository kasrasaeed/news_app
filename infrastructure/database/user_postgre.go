package database

import (
	"database/sql"
	"fmt"
	"github.com/kasrasaeed/news_app/domain/entity"
	"github.com/kasrasaeed/news_app/pkg/id"
	"strings"
)

type userPostgre struct {
	db *sql.DB
}

func NewUserPostgre(db *sql.DB) *userPostgre {
	return &userPostgre{db: db}
}

func (up *userPostgre) GetById(id id.UUID) (*entity.User, error) {
	rawQuery := `select id, user_name, pass_word, role, last_visit, created_date, updated_date from users where id = ?`
	stmt, err := up.db.Prepare(rawQuery)
	if err != nil {
		return nil, err
	}
	var user entity.User
	row := stmt.QueryRow(id)
	err = row.Scan(user.Id, user.UserName, user.PassWord, user.Role, user.LastVisit, user.CreatedDate, user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (up *userPostgre) GetByName(name string) (*entity.User, error) {
	rawQuery := `select from users where user_name = ?`
	stmt, err := up.db.Prepare(rawQuery)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(name)

	var user *entity.User
	err = row.Scan(
		user.Id,
		user.UserName,
		user.PassWord,
		user.Role,
		user.LastVisit,
		user.CreatedDate,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (up *userPostgre) Create(e *entity.User) (id.UUID, error) {
	user, err := up.GetByName(strings.ToLower(e.UserName))
	if err != nil {
		return [16]byte{}, err
	}
	if user != nil {
		return [16]byte{}, fmt.Errorf("this username already exists")
	}
	rawQuery := `insert into users (id, user_name, pass_word, role, last_visit, created_date, updated_date)
				 values (?,?,?,?,?,?,?)`
	stmt, err := up.db.Prepare(rawQuery)
	if err != nil {
		return e.Id, err
	}

	_, err = stmt.Exec(
		e.Id,
		strings.ToLower(e.UserName),
		e.PassWord,
		e.Role,
		e.LastVisit,
		e.CreatedDate.Format("2020-01-01"),
		e.UpdatedAt.Format("2020-01-01"),
	)
	if err != nil {
		return e.Id, err
	}
	err = stmt.Close()
	if err != nil {
		return e.Id, err
	}
	return e.Id, nil
}
