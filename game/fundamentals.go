package game

import "github.com/mhmorgan/rodla/phys"

type Model interface {
	// Pos returns the position of the model in the world.
	Pos() phys.Point
}

// ENTITIES

type Entity interface {
	Model
	Id() int // Database ID
}

type Living struct {
	Name       string
	dx, dy, dz int32
}

type Object struct {
	Name       string
	dx, dy, dz int32
}

// STRUCTURES

type Structure interface {
	Model
}
