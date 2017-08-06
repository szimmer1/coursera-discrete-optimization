package fib

// Fib - produces the fibonacci number at index n of the sequence
// [1, 2, 3, 5, 8, 13, 21, ...]
func Fib(n int, M []int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 2
	} else {
		if M == nil {
			return Fib(n-2, nil) + Fib(n-1, nil)
		} else {
			if M[n-2] == 0 {
				M[n-2] = Fib(n-2, M)
			}
			if M[n-1] == 0 {
				M[n-1] = Fib(n-1, M)
			}
			return M[n-2] + M[n-1]
		}
	}
}

func FibMemo(n int) int {
	M := make([]int, n)
	return Fib(n, M)
}
