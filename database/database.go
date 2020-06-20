package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() {
	conURL := os.Getenv("DATABASE_URL")
	if len(conURL) == 0 {
		conURL = "postgres://esdttdpr:oe81cJEBQN2_0Wr-LTfLTdlkME3a2tZw@john.db.elephantsql.com:5432/esdttdpr"
	}
	db, err = sql.Open("postgres", conURL)
	if err != nil {
		log.Fatal("Cannot connect database:", err)
	}

	createTb := `CREATE TABLE IF NOT EXISTS customer (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Cannot create table to database:", err)
	}
}

//CustomerRepository :
type customerRepository interface {
	base
	Search(condition map[string]interface{}) (rows *sql.Rows, err error)
}

//Conn :
// func Conn() *sql.DB {
// 	return db
// }

type posgresDB struct {
	db *sql.DB
}

//Repository :
func Repository() customerRepository {
	return &posgresDB{
		db: db,
	}
}

func (p posgresDB) Create(in map[string]string) (row *sql.Row, err error) {
	var name, email, status string
	var ok bool
	if name, ok = in["name"]; !ok {
		return nil, errors.New("missing name from input map")
	}
	if email, ok = in["email"]; !ok {
		return nil, errors.New("missing email from input map")
	}
	if status, ok = in["status"]; !ok {
		return nil, errors.New("missing status from input map")
	}
	row = db.QueryRow("INSERT INTO customer (name, email, status) values ($1, $2, $3)  RETURNING id, name, email, status", name, email, status)
	return
}

func (p posgresDB) GetOne(id string) (row *sql.Row, err error) {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer where id=$1")
	if err != nil {
		return nil, err
	}
	row = stmt.QueryRow(id)
	return
}

func (p posgresDB) GetAll() (rows *sql.Rows, err error) {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer")
	if err != nil {
		return nil, err
	}
	return stmt.Query()
}

func (p posgresDB) Update(id int, in map[string]string) (err error) {

	var name, email, status string
	var ok bool
	if name, ok = in["name"]; !ok {
		return errors.New("missing name from input map")
	}
	if email, ok = in["email"]; !ok {
		return errors.New("missing email from input map")
	}
	if status, ok = in["status"]; !ok {
		return errors.New("missing status from input map")
	}

	stmt, err := db.Prepare("UPDATE customer SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id, name, email, status); err != nil {
		return err
	}
	return
}

func (p posgresDB) Delete(id string) (err error) {
	stmt, err := db.Prepare("DELETE FROM customer WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return
}

func (p posgresDB) Search(condition map[string]interface{}) (rows *sql.Rows, err error) {
	return
}
