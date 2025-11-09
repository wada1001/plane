package components

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v *Vector) Sub(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v *Vector) Scale(scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v *Vector) Dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v *Vector) Length() float64 {
	return (v.X*v.X + v.Y*v.Y)
}
func (v *Vector) Normalize() Vector {
	length := v.Length()
	if length == 0 {
		return Vector{0, 0}
	}
	return Vector{
		X: v.X / length,
		Y: v.Y / length,
	}
}

func (v *Vector) Rotate(angle float64) Vector {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vector{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}
