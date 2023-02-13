package modes

import (
	"database/sql"
	"errors"
	"github.com/mhmorgan/rodla/data"
	"github.com/mhmorgan/rodla/game"
	"github.com/mhmorgan/rodla/ui"
)

func Start(s *game.State) game.Mode {
	times, err := updateTimesPlayed()
	if err != nil {
		panic(err)
	}
	if times == 1 {
		if err = createPlayer(); err != nil {
			panic(err)
		}
	}

	s.Player, err = data.Fetch[game.Player]("player")
	return Explore
}

func updateTimesPlayed() (uint, error) {
	times, err := data.Fetch[uint]("timesPlayed")
	if errors.Is(err, sql.ErrNoRows) {
		times = 1
	} else if err != nil {
		return 0, err
	} else {
		times++
	}
	err = data.Store("timesPlayed", times)
	return times, err
}

func createPlayer() error {
	name, err := ui.GetString("What is your name?")
	if err != nil {
		return nil
	}

	ui.Printf("Greetings, %s!\n", name)
	return data.Store("player", game.NewPlayer(name))
}
