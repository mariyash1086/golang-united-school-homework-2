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

const SideTriangle sideNameType = 3
const SideSquare sideNameType = 4
const SideCircle sideNameType = 0

func CalcSquare(sideLen float64, sidesNum sideNameType) float64 {
	switch sidesNum {
	case SideTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case SideSquare:
		return math.Pow(sideLen, 2)
	case SideCircle:
		return math.Pow(sideLen, 2) * math.Pi
	default:
		return 0
	}
}
