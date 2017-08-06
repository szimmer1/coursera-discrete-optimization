package main

// DPSolve - returns the objective value, whether solution is optimal,
// 			 and the vector of decision variables (x0...xI) in range [0,1]
func DPSolve(k int, w []int, v []int) (int, int, []int) {
	s := make([]int, len(w))
	return O(k, len(w)-1, w, v, s), 1, s
}

func DPSolveMemo(k int, w []int, v []int) (int, int, []int) {
}

func O(k int, i int, w []int, v []int, s []int) int {
	if i < 0 {
		return 0
	}
	if k < w[i] {
		return O(k, i-1, w, v, s)
	} else {
		pick := v[i] + O(k-w[i], i-1, w, v, s)
		unpick := O(k, i-1, w, v, s)
		if pick > unpick {
			s[i] = 1
			return pick
		} else {
			return unpick
		}
	}
}

func OMemo(k int, w []int, v []int, M [][]int) {
}
