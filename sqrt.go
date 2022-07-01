// function to calculate square root

// https://go.dev/tour/flowcontrol/8

func Sqrt(x float64) float64 {
	before := x
	z      := x
	for i := 0; i < 100; i++ {
		z -= (z * z - x) / (2 * z)
		if delta := z - before; delta > -0.0000001 {
			return z
		}
		before = z
	}
	return z
}
