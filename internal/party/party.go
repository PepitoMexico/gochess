package party

import (
	"gochess/gochess/internal/user"

	"github.com/notnil/chess"
)

type Party struct {
	name        string
	id          string
	blackPlayer user.User
	whitePlayer user.User
	game        *chess.Game
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

func (party *Party) playMove(move *chess.Move) {
	game := party.game
	moves := game.ValidMoves()
	for _, m := range moves {
		if m == move {
			game.Move(move)
		}
	}

}

func New(name string, id string, blackPlayer user.User, whitePlayer user.User) Party {
	return Party{name, id, blackPlayer, whitePlayer, chess.NewGame()}
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
