package db_model

import "fmt"

type User struct {
	ID       int
	USERNAME string
	EMAIL    string
	PASSWORD string
}

func (u User) CreateUser(id int, username, email, password string) *User {
	return &User{
		ID:       id,
		USERNAME: username,
		EMAIL:    email,
		PASSWORD: password,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("{\nID: %d, Username: %s, Email: %s, Password: %s,\n}", u.ID, u.USERNAME, u.EMAIL, u.PASSWORD)
}