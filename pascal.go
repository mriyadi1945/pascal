package pascal

import "fmt"

func Pascal(col, row int) int {
	if col == 0 || col == row {
		return 1
	}
	return Pascal(col-1, row-1) + Pascal(col, row-1)
}

func PascalIterative(nilai int) {
	n := nilai
	for i := range n {
		for j := 0; j <= n-i; j++ {
			fmt.Print("  ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("%4d", Pascal(k, i-1))
		}
		fmt.Println()
	}
}
