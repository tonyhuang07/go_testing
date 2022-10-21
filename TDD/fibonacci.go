package tdd

func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
	fibs := []int{0, 1}

	for i := 2; i <= n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}

	return fibs[n]
}
