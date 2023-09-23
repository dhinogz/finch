package data

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID             int
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(email, password string) error {
	fmt.Println(email)
	fmt.Println(password)

	return nil
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (u *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
