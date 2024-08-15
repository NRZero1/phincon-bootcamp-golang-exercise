package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID string
	Username string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Database struct {
	*sql.DB
}

func GetConnection(driver string, dataSource string) (*sql.DB, error) {
	conn, err := sql.Open(driver, dataSource)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetAllUser(connection *sql.DB) (User, error) {
	rows, err := connection.Query("SELECT * FROM users")

	if err != nil {
		return User{}, err
	}

	defer rows.Close()

	var allUser []User

	for rows.Next() {
		var user User

		var errScan = rows.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt)

		if errScan != nil {
            return User{}, errScan
		}

		allUser = append(allUser, user)
	}

	if err = rows.Err(); err != nil {
        fmt.Println(err.Error()) 
        return
    }

    for _, each := range allUser {
        fmt.Println(each.Username)
    }
}

func main() {
	dbConn := "postgres://postgres:root@localhost:5432/test_db_golang"

	conn, err := GetConnection("pgx", dbConn)

	defer conn.Close()

	fmt.Println(conn)

	rows, err := GetAllUser(conn)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows.Close()
}