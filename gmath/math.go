package gmath

import "math"

// Various useful float32 constants.
const (
	Pi   = math.Pi
	Phi  = math.Phi
	E    = math.E
	DtoR = Pi / 180
)

// Abs returns the absolute value of a float32.
func Abs(f float32) float32 {
	if f < 0.0 {
		return f * -1.0
	}
	return f
}

// Sign returns the sign of a float32.
func Sign(f float32) float32 {
	if f < 0.0 {
		return -1.0
	}
	return 1.0
}

// Min returns the minimum of two float32s.
func Min(a, b float32) float32 {
	if a <= b {
		return a
	}
	return b
}

// Max returns the maximum of two float32s.
func Max(a, b float32) float32 {
	if a >= b {
		return a
	}
	return b
}

// Clamp returns the float32 number closest to the input number within the input bounds.
func Clamp(num, min, max float32) float32 {
	return Min(Max(num, min), max)
}

// Sqrt returns the square root of a float32.
func Sqrt(f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

// Pow returns a float32 (base) to the power of a float32 (exp).
func Pow(base, exp float32) float32 {
	return float32(math.Pow(float64(base), float64(exp)))
}

// Round returns the rounded result of a float32.
func Round(f float32) float32 {
	return float32(math.Round(float64(f)))
}

// Floor returns the floored result of a float32.
func Floor(f float32) float32 {
	return float32(math.Floor(float64(f)))
}

// Ceil returns the ceiled result of a float32.
func Ceil(f float32) float32 {
	return float32(math.Ceil(float64(f)))
}

// ToRadians converts degrees to radians.
func ToRadians(degrees float32) float32 {
	return degrees * DtoR
}

// ToDegrees converts radians to degrees.
func ToDegrees(radians float32) float32 {
	return radians / DtoR
}

// Sin returns the sine of a float32.
func Sin(f float32) float32 {
	return float32(math.Sin(float64(f)))
}

// Cos returns the cosine of a float32.
func Cos(f float32) float32 {
	return float32(math.Cos(float64(f)))
}

// Tan returns the tangent of a float32.
func Tan(f float32) float32 {
	return float32(math.Tan(float64(f)))
}

// AbsI returns the absolute value of an int.
func AbsI(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

// MinI returns the minimum of two ints.
func MinI(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// MaxI returns the maximum of two ints.
func MaxI(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
