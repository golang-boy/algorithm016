package Week_01

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}

	if n < 3 {
		return n
	}

	first, second, third := 1, 2, 0
	for i := 3; i <= n; i++ {
		third = first + second
		first, second = second, third
	}

	return third
}
