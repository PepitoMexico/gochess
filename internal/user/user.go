package user

import (
	"database/sql"
	"gochess/gochess/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type User struct {
	Name     string
	Id       string
	Password string
}

func getUsers() {

}

func getUserByID() {

}

func getUserbyName() {

}

func New(name string, id string, password string) User {
	return User{name, id, password}
}

func CreateUser(name string) User {

	driver := config.GetConfig().GetString("database.driver")
	host := config.GetConfig().GetString("database.host")
	database := config.GetConfig().GetString("database.database")
	username := config.GetConfig().GetString("database.username")
	password := config.GetConfig().GetString("database.password")
	db, err := sql.Open(driver, username+":"+password+"@tcp("+host+")/"+database)
	if err != nil {
		panic(err)
	}
	id := uuid.New().String()
	user := New(name, id, password)

	db.Exec("INSERT INTO `users` (`id`, `name`) VALUES ('" + user.Id + "', '" + user.Name + "')")
	return user
}

func (user *User) setName(name string) {
	user.Name = name
}

func (user *User) setID(id string) {
	user.Id = id
}
