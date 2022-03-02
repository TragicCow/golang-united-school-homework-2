package square

import "math"

type newInt int

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const SidesTriangle newInt = 3
const SidesSquare newInt = 4
const SidesCircle newInt = 0

func CalcSquare(sideLen float64, sidesNum newInt) float64 {
	const pi = math.Pi
	switch sidesNum {

	case SidesTriangle:
		return math.Sqrt(3) / 4 * sideLen * sideLen
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return pi * sideLen * sideLen
	default:
		return 0
	}
}
