package user

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type User struct {
	Name string
	Id   string
}

func getUsers() {

}

func getUserByID() {

}

func getUserbyName() {

}

func New(name string, id string) User {
	return User{name, id}
}

func createUser(name string) {
	id := uuid.New().String()

	user := New(name, id)

	viper.SetConfigName("config")

	db, err := sql.Open("mysql", "gochess:password@/dbname")
	if err != nil {
		panic(err)
	}
	db.r

}

func (user *User) setName(name string) {
	user.Name = name
}

func (user *User) setID(id string) {
	user.Id = id
}
