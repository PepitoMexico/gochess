package user

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
