package main

import (
	"fmt"
)

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func intSign(i int) int {
	switch {
	case i < 0:
		return -1
	case i > 0:
		return 1
	default:
		return 0
	}
}

type Vector struct {
	x, y, z int
}

func (this Vector) Add(other Vector) Vector {
	return Vector{
		x: this.x + other.x,
		y: this.y + other.y,
		z: this.z + other.z,
	}
}

func (v Vector) Energy() int {
	return intAbs(v.x) + intAbs(v.y) + intAbs(v.z)
}

type Moon struct {
	pos, vel Vector
}

func (m *Moon) Energy() int {
	return m.pos.Energy() * m.vel.Energy()
}

func (this *Moon) UpdateVelocity(other *Moon) {
	this.vel.x += intSign(other.pos.x - this.pos.x)
	this.vel.y += intSign(other.pos.y - this.pos.y)
	this.vel.z += intSign(other.pos.z - this.pos.z)
}

func (this *Moon) Move() {
	this.pos = this.pos.Add(this.vel)
}

func main() {
	// moons := []*Moon{
	// 	{pos: Vector{-1, 0, 2}},
	// 	{pos: Vector{2, -10, -7}},
	// 	{pos: Vector{4, -8, 8}},
	// 	{pos: Vector{3, 5, -1}},
	// }
	// moons := []*Moon{
	// 	{pos: Vector{-8, -10, 0}},
	// 	{pos: Vector{5, 5, 10}},
	// 	{pos: Vector{2, -7, 3}},
	// 	{pos: Vector{9, -8, -3}},
	// }
	moons := []*Moon{
		{pos: Vector{6, 10, 10}},
		{pos: Vector{-9, 3, 17}},
		{pos: Vector{9, -4, 14}},
		{pos: Vector{4, 14, 4}},
	}

	const steps = 1000

	for s := 0; s < steps; s++ {
		for i, m1 := range moons {
			for j, m2 := range moons {
				if i == j {
					continue
				}

				m1.UpdateVelocity(m2)
			}
		}

		for _, m := range moons {
			m.Move()
		}
	}

	total := 0
	for _, m := range moons {
		total += m.Energy()
	}

	fmt.Println(total)
}
