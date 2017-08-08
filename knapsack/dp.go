package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// DPSolve - returns the objective value, whether solution is optimal,
// 			 and the vector of decision variables (x0...xI) in range [0,1]
func DPSolve(k int, w []int, v []int) (int, []int) {
	s := make([]int, len(w))
	return O(k, len(w)-1, w, v, s), s
}

func DPSolveMemo(k int, w []int, v []int) (int, []int) {
	// initialize matrix with -1 values
	fmt.Println("start making matrix")
	M := make([][]int, k+1)
	I := len(w)
	for i := 0; i < k+1; i++ {
		fmt.Printf("i=%d ", i)
		M[i] = make([]int, I)
		for j, _ := range M[i] {
			M[i][j] = -1
		}
	}
	fmt.Println("start DPSolveMemo")

	objVal := OMemo(k, I-1, w, v, M)
	// determine decision variables from memo table
	return objVal, []int{}
}

func O(k int, i int, w []int, v []int, s []int) int {
	if i < 0 {
		return 0
	}
	if k < w[i] {
		return O(k, i-1, w, v, s)
	}
	pick := v[i] + O(k-w[i], i-1, w, v, s)
	unpick := O(k, i-1, w, v, s)
	if pick > unpick {
		s[i] = 1
		return pick
	} else {
		return unpick
	}
}

func OMemo(k int, i int, w []int, v []int, M [][]int) int {
	fmt.Printf("k=%d, i=%d\n", k, i)
	// recursive base condition
	if i < 0 || k < 0 {
		return 0
	}

	// use the memo if available
	if M[k][i] >= 0 {
		fmt.Printf("found O(k, i)=%d\n", M[k][i])
		return M[k][i]
	}

	fmt.Println("calculating O(k, i)")
	if k < w[i] {
		return OMemo(k, i-1, w, v, M)
	}

	// determine optimal value & return
	M[k][i] = max(
		v[i]+OMemo(k-w[i], i-1, w, v, M),
		OMemo(k, i-1, w, v, M),
	)
	return M[k][i]
}
