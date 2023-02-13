package game

type Player struct {
	Living

	Name string
}

func NewPlayer(name string) Player {
	return Player{
		Name: name,
	}
}
