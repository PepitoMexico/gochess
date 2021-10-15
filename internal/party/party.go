package party

import (
	"gochess/gochess/internal/user"
)

type Party struct {
	name        string
	id          string
	blackPlayer user.User
	whitePlayer user.User
}

func launchParty() {

}

func stopParty() {

}

func getBlackPlayer() {

}

func getWhitePlayer() {

}

func getPartyByID() {

}

func New(name string, id string, blackPlayer user.User, whitePlayer user.User) Party {
	return Party{name, id, blackPlayer, whitePlayer}
}

func (party *Party) setName(name string) {
	party.name = name
}

func (party *Party) setID(id string) {
	party.id = id
}

func (party *Party) setblackPlayer(blackPlayer user.User) {
	party.blackPlayer = blackPlayer
}

func (party *Party) setwhitePlayer(whitePlayer user.User) {
	party.whitePlayer = whitePlayer
}
