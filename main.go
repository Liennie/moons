package main

import (
	"fmt"
	"sync"
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

// stolen from https://play.golang.org/p/SmzvkDjYlb
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// stolen from https://play.golang.org/p/SmzvkDjYlb
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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

func (this *Moon) UpdateVelocityX(other *Moon) {
	this.vel.x += intSign(other.pos.x - this.pos.x)
}

func (this *Moon) UpdateVelocityY(other *Moon) {
	this.vel.y += intSign(other.pos.y - this.pos.y)
}

func (this *Moon) UpdateVelocityZ(other *Moon) {
	this.vel.z += intSign(other.pos.z - this.pos.z)
}

func (this *Moon) UpdateVelocity(other *Moon) {
	this.UpdateVelocityX(other)
	this.UpdateVelocityY(other)
	this.UpdateVelocityZ(other)
}

func (this *Moon) MoveX() {
	this.pos.x += this.vel.x
}

func (this *Moon) MoveY() {
	this.pos.y += this.vel.y
}

func (this *Moon) MoveZ() {
	this.pos.z += this.vel.z
}

func (this *Moon) Move() {
	this.MoveX()
	this.MoveY()
	this.MoveZ()
}

func (this *Moon) Copy() *Moon {
	return &Moon{
		pos: this.pos,
		vel: this.vel,
	}
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

	initial := make([]*Moon, len(moons))
	for i, moon := range moons {
		initial[i] = moon.Copy()
	}

	stepsX := 0
	stepsY := 0
	stepsZ := 0

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		steps := 0

		for {
			for i, m1 := range moons {
				for j, m2 := range moons {
					if i == j {
						continue
					}

					m1.UpdateVelocityX(m2)
				}
			}

			for _, m := range moons {
				m.MoveX()
			}

			steps++

			repeating := true
			for i, m1 := range moons {
				m2 := initial[i]
				if m1.pos.x != m2.pos.x || m1.vel.x != m2.vel.x {
					repeating = false
					break
				}
			}

			if repeating {
				stepsX = steps
				break
			}
		}
	}()

	go func() {
		defer wg.Done()

		steps := 0

		for {
			for i, m1 := range moons {
				for j, m2 := range moons {
					if i == j {
						continue
					}

					m1.UpdateVelocityY(m2)
				}
			}

			for _, m := range moons {
				m.MoveY()
			}

			steps++

			repeating := true
			for i, m1 := range moons {
				m2 := initial[i]
				if m1.pos.y != m2.pos.y || m1.vel.y != m2.vel.y {
					repeating = false
					break
				}
			}

			if repeating {
				stepsY = steps
				break
			}
		}
	}()

	go func() {
		defer wg.Done()

		steps := 0

		for {
			for i, m1 := range moons {
				for j, m2 := range moons {
					if i == j {
						continue
					}

					m1.UpdateVelocityZ(m2)
				}
			}

			for _, m := range moons {
				m.MoveZ()
			}

			steps++

			repeating := true
			for i, m1 := range moons {
				m2 := initial[i]
				if m1.pos.z != m2.pos.z || m1.vel.z != m2.vel.z {
					repeating = false
					break
				}
			}

			if repeating {
				stepsZ = steps
				break
			}
		}
	}()

	wg.Wait()

	fmt.Println(LCM(stepsX, stepsY, stepsZ))
}
