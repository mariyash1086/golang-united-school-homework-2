package solution

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sideNameType int

const sideTriangle sideNameType = 3
const sideSquare sideNameType = 4
const sideCircle sideNameType = 0

func CalcSquare(sideLen float64, sidesNum sideNameType) float64 {
	switch sidesNum {
	case sideTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case sideSquare:
		return math.Pow(sideLen, 2)
	case sideCircle:
		return math.Pow(sideLen, 2) * math.Pi
	default:
		return 0
	}
}
