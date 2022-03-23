package database

import (
	"database/sql"
	"github.com/kasrasaeed/news_app/domain/entity"
	"github.com/kasrasaeed/news_app/pkg/id"
)

type userPostgre struct {
	db *sql.DB
}

func NewUserPostgre(db *sql.DB) *userPostgre {
	return &userPostgre{db: db}
}

func (up *userPostgre) Create(e *entity.User) (id.UUID, error) {
	rawQuery := `insert into users (id, user_name, pass_word, role, last_visit, created_date, updated_date)
				 values (?,?,?,?,?,?,?)`
	stmt, err := up.db.Prepare(rawQuery)
	if err != nil {
		return e.Id, err
	}

	_, err = stmt.Exec(
		e.Id,
		e.UserName,
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
