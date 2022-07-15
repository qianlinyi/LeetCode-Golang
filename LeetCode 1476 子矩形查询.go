package main

type SubrectangleQueries struct {
	rectangle [][]int
	history   [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rectangle,
		[][]int{},
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.history = append(this.history, []int{row1, col1, row2, col2, newValue})
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	for i := len(this.history) - 1; i >= 0; i-- {
		if this.history[i][0] <= row && row <= this.history[i][2] && this.history[i][1] <= col && col <= this.history[i][3] {
			return this.history[i][4]
		}
	}
	return this.rectangle[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */

func main() {}
