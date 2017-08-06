package main

// DPSolve - returns the objective value, whether solution is optimal,
// 			 and the vector of decision variables (x0...xI) in range [0,1]
func DPSolve(k int, w []int, v []int) (int, int, []int) {
	s := make([]int, len(w))
	return DPSolveAux(k, len(w)-1, w, v, s), 1, s
}

func DPSolveMemo(k int, w []int, v []int) (int, int, []int) {
}

func DPSolveAux(k int, i int, w []int, v []int, s []int) int {
	if i < 0 {
		return 0
	}
	if k < w[i] {
		return DPSolveAux(k, i-1, w, v, s)
	} else {
		pick := v[i] + DPSolveAux(k-w[i], i-1, w, v, s)
		unpick := DPSolveAux(k, i-1, w, v, s)
		if pick > unpick {
			s[i] = 1
			return pick
		} else {
			return unpick
		}
	}
}
