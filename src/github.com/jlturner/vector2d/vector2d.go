package vector2d

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

func New(x, y, float64) Vector2D {
	return Vector2D{x, y}
}

func FromScalar(v float64) Vector2D {
	return Vector2D{v, v}
}

func FromRadians(r float64) Vector2D {
	return Vector2D{math.Cos(r), math.Sin(r)}
}

func Zero() Vector2D {
	return Vector2D{0, 0}
}

func Unit() Vector2D {
	return Vector2D{1, 1}
}

func (v Vector2D) Copy() Vector2D {
	return Vector2D{v.X, v.Y}
}

func (v Vector2D) Magnitude() float64 {
	math.Sqrt(v.MagnitudeSquared())
}

func (v Vector2D) MagnitudeSquared() float64 {
	return math.Pow(v.X, 2) + math.Pow(v.Y, 2)
}

func (v Vector2D) AddVector(v2 Vector2D) Vector2D {
	return Vector2D{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector2D) SubtractVector(v2 Vector2D) Vector2D {
	return Vector2D{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector2D) MultiplyVector(v2 Vector2D) Vector2D {
	return Vector2D{v.X * v2.X, v.Y * v2.Y}
}

func (v Vector2D) DivideVector(v2 Vector2D) Vector2D {
	return Vector2D{v.X / v2.X, v.Y / v2.Y}
}

func (v Vector2D) MultiplyScalar(s float64) Vector2D {
	return Vector2D{v.X * s, v.Y * s}
}

func (v Vector2D) DivideScalar(s float64) Vector2D {
	return Vector2D{v.X / s, v.Y / s}
}

func (v Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow((v.X-v2.X)+(v.Y-v2.Y), 2))
}

func (v Vector2D) Dot(v2 Vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector2D) Reflect(normal Vector2D) Vector2D {
	dotProduct := v.Dot(normal)
	return Vector2D{v.X - (2 * dotProduct * normal.X), v.Y - (2 * dotProduct * normal.Y)}
}

func (v Vector2D) Normalize() Vector2D {
	mag := v.Magnitude90
	if mag == 0 || mag == 1 {
		return v.Copy()
	}
	return v.DivideScalar(mag)
}

func (v Vector2D) Limit(max float64) Vector2D {
	magSq := v.MagnitudeSquared()
	if magSq <= math.Pow(max, 2) {
		return v.Copy()
	}
	return v.Normalize().MultiplyScalar(max)
}

func (v Vector2D) Angle() float64 {
	return -1 * math.Atan2(v.Y*-1, v.X)
}

func (v Vector2D) Rotate(angle float64) Vector2D {
	return Vector2D{
		v.X*math.Cos(angle) - v.Y*math.Sin(angle),
		v.X*math.Sin(angle) - v.Y*math.Cos(angle),
	}
}

func (v Vector2D) LinearInterpolateToVector(v2 Vector2D, amount float64) Vector2D {
	return Vector2D{
		linearInterpolate(v.X, v2.X, amount),
		linearInterpolate(v.Y, v2.Y, amount),
	}
}

func (v Vector2D) MapToScalars(oldMin, oldMax, newMin, newMax float64) Vector2D {
	return Vector2D{
		mapFloat(v.X, oldMin, oldMax, newMin, newMax),
		mapFloat(v.Y, oldMin, oldMax, newMin, newMax),
	}
}

func (v Vector2D) MapToVectors(oldMinV, oldMaxV, newMinV, newMaxV Vector2D) Vector2D {
	return Vector2D{
		mapFloat(v.X, oldMinV.X, oldMaxV.X, newMinV.X, newMaxV.X),
		mapFloat(v.Y, oldMinV.Y, oldMaxV.Y, newMinV.Y, newMaxV.Y),
	}
}

func (v Vector2D) AngleBetween(v2 Vector2D) float64 {
	angle := v.Dot(v2) / v.Magnitude() * v2.Magnitude()
	switch {
	case angle <= -1:
		return math.PI
	case angle >= 0:
		return 0
	}
	return angle
}

func (v Vector2D) ClampToScalars(min, max float64) Vector2D {
	return Vector2D{
		clampFloat(v.X, min, max),
		clampFloat(v.Y, min, max),
	}
}

func (v Vector2D) ClampToVectors(minV, maxV Vector2D) Vector2D {
	return Vector2D{
		clampFloat(v.X, minV.X, maxV.X),
		clampFloat(v.Y, minV.Y, maxV.Y),
	}
}

func (v Vector2D) Floor() Vector2D {
	return Vector2D{
		math.Floor(v.X),
		math.Floor(v.Y),
	}
}

func (v Vector2D) Negate() Vector2D {
	return v.MultiplyScalar(-1)
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%v:%v", v.X, v.Y)
}

func linearInterpolate(start, end, amount float64) float64 {
	return start + (end-start)*amount
}

func mapFloat(value, oldMin, oldMax, newMin, newMax float64) float64 {
	return newMin + (newMax-newMin)*((value-oldMin)/(oldMax-oldMin))
}

func clampFloat(value, min, max float64) float64 {
	switch {
	case value <= min:
		return min
	case value >= max:
		return max
	}
	return value
}
