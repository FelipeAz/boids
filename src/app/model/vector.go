package model

import (
	"math"
)

type Vector struct {
	x float64
	y float64
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}

func (v Vector) Subtract(v2 Vector) Vector {
	return Vector{
		x: v.x - v2.x,
		y: v.y - v2.y,
	}
}

func (v Vector) Multiply(v2 Vector) Vector {
	return Vector{
		x: v.x * v2.x,
		y: v.y * v2.y,
	}
}

func (v Vector) AddV(val float64) Vector {
	return Vector{
		x: v.x + val,
		y: v.y + val,
	}
}

func (v Vector) SubtractV(val float64) Vector {
	return Vector{
		x: v.x - val,
		y: v.y - val,
	}
}

func (v Vector) MultiplyV(val float64) Vector {
	return Vector{
		x: v.x * val,
		y: v.y * val,
	}
}

func (v Vector) DivideV(val float64) Vector {
	return Vector{
		x: v.x / val,
		y: v.y / val,
	}
}

func (v Vector) limit(lower, upper float64) Vector {
	return Vector{
		x: math.Min(math.Max(v.x, lower), upper),
		y: math.Min(math.Max(v.y, lower), upper),
	}
}

func (v Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v.x - v2.x, 2) + math.Pow(v.y - v2.y, 2))
}