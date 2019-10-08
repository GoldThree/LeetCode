package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / core(x, -n)
	}
	return core(x, n)
}

func core(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := core(x, n/2)
	if n%2 == 0 {
		return res * res
	} else {
		return res * res * x
	}
}
