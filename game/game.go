package game

type Mode func(*State) Mode

type State struct {
	Player    Player
	WorldName string
}

func NewState(worldName string) *State {
	return &State{WorldName: worldName}
}
