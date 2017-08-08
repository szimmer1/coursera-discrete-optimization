package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var n, k, tempV, tempW int

	fmt.Fscanf(f, "%d %d", &n, &k)

	w := make([]int, n)
	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(f, "%d %d", &tempV, &tempW)
		w[i] = tempW
		v[i] = tempV
	}

	fmt.Printf("n=%d [%T], k=%d, w=%v, v=%v\n", n, n, k, w, v)
	//fmt.Println("No memo")
	//fmt.Println(DPSolve(k, w, v))
	fmt.Println("Memo")
	fmt.Println(DPSolveMemo(k, w, v))
}
