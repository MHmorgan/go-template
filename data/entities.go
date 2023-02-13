package data

import (
	"github.com/mhmorgan/rodla/game"
)

func FetchEntity[E game.Entity](id int) (e E, err error) {
	// TODO: Fetch entity from database by ID
	return
}

func FetchEntities(opts ...FetchOption) (es []game.Entity, err error) {
	// TODO Fetch entities from database
	return
}
