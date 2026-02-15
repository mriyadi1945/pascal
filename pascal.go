package pascal

func Pascal(col, row int) int {
	if col == 0 || col == row {
		return 1
	}
	return Pascal(col-1, row-1) + Pascal(col, row-1)
}
