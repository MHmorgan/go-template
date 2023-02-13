// Package phys implements geometric primitives.
//
// The x-axis goes west-east, increasing eastwards.
// The y-axis goes north-south, increasing northwards.
// The z-axis goes up-down, increasing upwards.
//
// Geometric ranges are inclusive of the minimum and exclusive of the maximum.
//
// TODO Handle edge cases, max/min values, etc.
package phys

type Sized interface {
	X() int32
	Y() int32
	Z() int32
	MidX() int32
	MidY() int32
	MidZ() int32
	MaxX() int32
	MaxY() int32
	MaxZ() int32
	DX() uint32
	DY() uint32
	DZ() uint32
	Pos() Point
	Rect() Rect
	Cube() Cube
	Area() Area
	Volume() Volume
}

// -----------------------------------------------------------------------------
// AREA

type Area struct {
	dx, dy uint32
}

func NewArea(dx, dy uint32) Area {
	return Area{dx, dy}
}

func (a Area) DX() uint32 {
	return a.dx
}

func (a Area) DY() uint32 {
	return a.dy
}

// -----------------------------------------------------------------------------
// VOLUME

type Volume struct {
	dx, dy, dz uint32
}

func NewVolume(dx, dy, dz uint32) Volume {
	return Volume{dx, dy, dz}
}

func (v Volume) DX() uint32 {
	return v.dx
}

func (v Volume) DY() uint32 {
	return v.dy
}

func (v Volume) DZ() uint32 {
	return v.dz
}

// -----------------------------------------------------------------------------
// POSITION

type Point struct {
	x, y, z int32
}

func NewPoint(x, y, z int32) Point {
	return Point{x, y, z}
}

func (p Point) X() int32 {
	return p.x
}

func (p Point) Y() int32 {
	return p.y
}

func (p Point) Z() int32 {
	return p.z
}

func (p Point) MidX() int32 {
	return p.x
}

func (p Point) MidY() int32 {
	return p.y
}

func (p Point) MidZ() int32 {
	return p.z
}

func (p Point) MaxX() int32 {
	return p.x + 1
}

func (p Point) MaxY() int32 {
	return p.y + 1
}

func (p Point) MaxZ() int32 {
	return p.z + 1
}

func (p Point) DX() uint32 {
	return 0
}

func (p Point) DY() uint32 {
	return 0
}

func (p Point) DZ() uint32 {
	return 0
}

func (p Point) Pos() Point {
	return p
}

func (p Point) Area() Area {
	return Area{}
}

func (p Point) Volume() Volume {
	return Volume{}
}

func (p Point) Rect() Rect {
	return Rect{p.x, p.y, p.z, 0, 0}
}

func (p Point) Cube() Cube {
	return Cube{p.x, p.y, p.z, 0, 0, 0}
}

var _ Sized = Point{}

// -----------------------------------------------------------------------------
// RECTANGLE

type Rect struct {
	x, y, z int32
	dx, dy  uint32
}

func NewRect(x, y, z int32, dx, dy uint32) Rect {
	return Rect{x, y, z, dx, dy}
}

func NewRectFromEdges(minX, minY, maxX, maxY, z int32) Rect {
	dx := maxX - minX - 1
	if dx < 0 {
		dx = 0
	}
	dy := maxY - minY - 1
	if dy < 0 {
		dy = 0
	}
	return Rect{minX, minY, z, uint32(dx), uint32(dy)}
}

func (r Rect) X() int32 {
	return r.x
}

func (r Rect) Y() int32 {
	return r.y
}

func (r Rect) Z() int32 {
	return r.z
}

func (r Rect) MidX() int32 {
	return r.x + int32(r.dx)/2
}

func (r Rect) MidY() int32 {
	return r.y + int32(r.dy)/2
}

func (r Rect) MidZ() int32 {
	return r.z
}

func (r Rect) MaxX() int32 {
	return r.x + int32(r.dx) + 1
}

func (r Rect) MaxY() int32 {
	return r.y + int32(r.dy) + 1
}

func (r Rect) MaxZ() int32 {
	return r.z + 1
}

func (r Rect) DX() uint32 {
	return r.dx
}

func (r Rect) DY() uint32 {
	return r.dy
}

func (r Rect) DZ() uint32 {
	return 0
}

func (r Rect) Pos() Point {
	return Point{r.x, r.y, r.z}
}

func (r Rect) Rect() Rect {
	return r
}

func (r Rect) Cube() Cube {
	return Cube{r.x, r.y, r.z, r.dx, r.dy, 0}
}

func (r Rect) Area() Area {
	return Area{r.dx, r.dy}
}

func (r Rect) Volume() Volume {
	return Volume{r.dx, r.dy, 0}
}

var _ Sized = Rect{}

// -----------------------------------------------------------------------------
// CUBE

type Cube struct {
	x, y, z    int32
	dx, dy, dz uint32
}

func NewCube(x, y, z int32, dx, dy, dz uint32) Cube {
	return Cube{x, y, z, dx, dy, dz}
}

func NewCubeFromEdges(minX, minY, minZ, maxX, maxY, maxZ int32) Cube {
	dx := maxX - minX - 1
	if dx < 0 {
		dx = 0
	}
	dy := maxY - minY - 1
	if dy < 0 {
		dy = 0
	}
	dz := maxZ - minZ - 1
	if dz < 0 {
		dz = 0
	}
	return Cube{minX, minY, minZ, uint32(dx), uint32(dy), uint32(dz)}
}

func (c Cube) X() int32 {
	return c.x
}

func (c Cube) Y() int32 {
	return c.y
}

func (c Cube) Z() int32 {
	return c.z
}

func (c Cube) MidX() int32 {
	return c.x + int32(c.dx)/2
}

func (c Cube) MidY() int32 {
	return c.y + int32(c.dy)/2
}

func (c Cube) MidZ() int32 {
	return c.z + int32(c.dz)/2
}

func (c Cube) MaxX() int32 {
	return c.x + int32(c.dx) + 1
}

func (c Cube) MaxY() int32 {
	return c.y + int32(c.dy) + 1
}

func (c Cube) MaxZ() int32 {
	return c.z + int32(c.dz) + 1
}

func (c Cube) DX() uint32 {
	return c.dx
}

func (c Cube) DY() uint32 {
	return c.dy
}

func (c Cube) DZ() uint32 {
	return c.dz
}

func (c Cube) Pos() Point {
	return Point{c.x, c.y, c.z}
}

func (c Cube) Rect() Rect {
	return Rect{c.x, c.y, c.z, c.dx, c.dy}
}

func (c Cube) Cube() Cube {
	return c
}

func (c Cube) Area() Area {
	return Area{c.dx, c.dy}
}

func (c Cube) Volume() Volume {
	return Volume{c.dx, c.dy, c.dz}
}

var _ Sized = Cube{}
