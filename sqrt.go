package main

import "fmt"
import "math"

func main() {
	var z float64 = 65434567
	fmt.Println(">>>", sqrt(z))
	fmt.Println("go>", math.Sqrt(z))
}

func sqrt(x float64) float64 {
	var step float64

	step = 1
	z := initial(x)

	for i := 0; i < 15; i++ {
		low := z - 5
		hi := z + 5
		n := newton(low, hi, step, x)
		step = step / 10
		if z == n {
			fmt.Println(i, "times")
			return z
		}
		z = n
	}

	fmt.Println("10 times")
	return z
}

func newton(low, hi, step, x float64) float64 {
	var z float64
	result := []float64{}

	for i := low; i < hi; i++ {
		z = (1 + float64(i) + step)
		z -= (z*z - x) / (2 * z)
		result = append(result, z)
	}

	var y [10]float64
	copy(y[:], result[:10])
	return closest(y, x)
}

func initial(x float64) float64 {
	var z float64
	var result [10]float64
	step := stepten(x)

	for i := 0; i < 10; i++ {
		z = (1 + float64(i) + step)
		z -= (z*z - x) / (2 * z)
		result[i] = z
	}
	return closest(result, x)
}

func closest(candidate [10]float64, x float64) float64 {
	var result [10]float64
	var picked int
	for i := range candidate {
		result[i] = x - (candidate[i] * candidate[i])
	}

	for i := 1; i < 10; i++ {
		if math.Abs(result[i]) < math.Abs(result[i-1]) {
			picked = i
		}
	}
	return candidate[picked]
}

func stepten(x float64) float64 {
	return x / 10
}
