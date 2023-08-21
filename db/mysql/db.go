package db

import (
	"database/sql"
	"time"
	"github.com/dimalien/db-web-server-golang.git/db/model"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	)

var (
	ctx context.Context	
)

func ConnnectToDataBase() *sql.DB {
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)

	return db
}

func GetUserByID(ID int, db *sql.DB) (*db_model.User, error) {
	stmt := "SELECT * FROM users WHERE id=?"
	row := db.QueryRow(stmt, ID)

	defer db.Close()

	user := &db_model.User{} 

	err := row.Scan(&user.ID, &user.USERNAME, &user.EMAIL, &user.PASSWORD)

	if errors.Is(err, sql.ErrNoRows) {
		return  nil, err
	}

	return user, nil	
}

func InsertUser(user *db_model.User, db *sql.DB) (sql.Result, error) {

	stmt := "INSERT INTO users(id, username, email, password) VALUES (?, ?, ?, ?);"

	res, err := db.Exec(stmt, user.ID, user.USERNAME, user.EMAIL, user.PASSWORD)

	db.Close()

	return res, err
} 