package database

import "database/sql"

type base interface {
	Create(in map[string]string) (row *sql.Row, err error)
	Update(id int, in map[string]string) (err error)
	GetAll() (rows *sql.Rows, err error)
	GetOne(id string) (row *sql.Row, err error)
	Delete(id string) (err error)
}
