package util

const (
	// AddChar is a value with the type of char.
	AddChar = '+'
	// SubstractChar is a value with the type of char.
	SubstractChar = '-'
	// MultiplyChar is a value with the type of char.
	MultiplyChar = '*'
	// DivideChar is a value with the type of char.
	DivideChar = '/'
	// PercentageChar is a value with the type of char.
	PercentageChar = '%'
	// PiChar is a value with the type of a greek char.
	PiChar = 'π'

	// Pi is the circle number: 3.14159265359
	Pi = 3.14159265359
)

/*
CalculateString is a function which calculates a string and returns the value with calculationWay
*/
func CalculateString(str string) (result int, calculationWay string) {
	// INFO: first braces, then potencies, then point, the stroke
	// 1. Detect braces and solve it
	// 2. Detect chars like *;/;-;+;%;π math functions like "sin" and potencies
	// TODO
	result = 0
	calculationWay = "-/-"
	return
}

/*
SolveEquationString is a function which solves a equation and returns the value with calculationWay
*/
func SolveEquationString(str string) (resultLeftSide int, resultRightSide int, calculationWay string) {
	// INFO: first braces, then potencies, then point, the stroke
	// 1. Detect braces and solve it
	// 2. Detect chars like *;/;-;+;%;π math functions like "sin" and potencies
	// TODO
	resultLeftSide = 0
	resultRightSide = 0
	calculationWay = "-/-"
	return
}

/*
Add is a function, which adds multiple ints to one
*/
func Add(ints ...int) int {
	var out int
	for _, i := range ints {
		out += i
	}
	return out
}

/*
AddInt32 is a function, which adds multiple ints to one
*/
func AddInt32(ints ...int32) int32 {
	var out int32
	for _, i := range ints {
		out += i
	}
	return out
}

/*
AddInt64 is a function, which adds multiple ints to one
*/
func AddInt64(ints ...int64) int64 {
	var out int64
	for _, i := range ints {
		out += i
	}
	return out
}

/*
AddFloat32 is a function, which adds multiple floats to one
*/
func AddFloat32(ints ...float32) float32 {
	var out float32
	for _, i := range ints {
		out += i
	}
	return out
}

/*
AddFloat64 is a function, which adds multiple floats to one
*/
func AddFloat64(ints ...float64) float64 {
	var out float64
	for _, i := range ints {
		out += i
	}
	return out
}

/*
Substract is a function, which substracts multiple ints to one
*/
func Substract(ints ...int) int {
	var out int
	for _, i := range ints {
		out -= i
	}
	return out
}

/*
SubstractInt32 is a function, which substracts multiple ints to one
*/
func SubstractInt32(ints ...int32) int32 {
	var out int32
	for _, i := range ints {
		out -= i
	}
	return out
}

/*
SubstractInt64 is a function, which substracts multiple ints to one
*/
func SubstractInt64(ints ...int64) int64 {
	var out int64
	for _, i := range ints {
		out -= i
	}
	return out
}

/*
SubstractFloat32 is a function, which substracts multiple floats to one
*/
func SubstractFloat32(ints ...float32) float32 {
	var out float32
	for _, i := range ints {
		out -= i
	}
	return out
}

/*
SubstractFloat64 is a function, which substracts multiple floats to one
*/
func SubstractFloat64(ints ...float64) float64 {
	var out float64
	for _, i := range ints {
		out -= i
	}
	return out
}

// TODO: Multiply & Divide should be added and a percentage resolver etc.
