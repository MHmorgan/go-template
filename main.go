package main

import (
	"fmt"
	"github.com/mhmorgan/rodla/data"
	"github.com/mhmorgan/rodla/game"
	"github.com/mhmorgan/rodla/modes"
	"os"
)

func main() {
	err := data.Init(":memory:")
	if err != nil {
		bail(err)
	}
	defer func() {
		if err := data.Close(); err != nil {
			bail(err)
		}
	}()

	state := game.NewState("Test World")
	gameMode := modes.Start

	for gameMode != nil {
		gameMode = gameMode(state)
	}
}

func bail(err error) {
	fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(1)
}
