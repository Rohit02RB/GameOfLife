package main

import (
	"fmt"
)

// Matrix represents the grid for the Game of Life
type Matrix struct {
	grid [][]int
	row  int
	col  int
}

// NewMatrix initializes a new matrix with the given dimensions
func NewMatrix(row, col int) *Matrix {
	m := &Matrix{
		grid: make([][]int, row),
		row:  row,
		col:  col,
	}
	for i := range m.grid {
		m.grid[i] = make([]int, col)
	}
	return m
}

func main() {
	row, col := getMatrixDimensions()

	mat := NewMatrix(row, col)
	mat.PrintMatrix()

	mat.SetLiveCell()
	fmt.Println("Matrix after setting live cells:")
	mat.PrintMatrix()

	mat.grid = mat.GetUpdateMatrix()
	fmt.Println("Updated Matrix:")
	mat.PrintMatrix()

}

// getMatrixDimensions helps the user to enter the dimensions of the matrix
func getMatrixDimensions() (int, int) {
	var row, col int
	fmt.Println("Enter the number of rows and columns:")
	fmt.Scan(&row, &col)
	return row, col
}

// PrintMatrix prints the current state of the matrix
func (m *Matrix) PrintMatrix() {
	for _, row := range m.grid {
		fmt.Println(row)
	}
}

// SetLiveCell set live cell in the matrix
func (mat *Matrix) SetLiveCell() {

	var totalLiveCells int
	fmt.Println("Enter the total number of live cells:")
	fmt.Scan(&totalLiveCells)
	for i := 0; i < totalLiveCells; i++ {
		var row, col int
		fmt.Println("Enter the row and column of a live cell:")
		fmt.Scan(&row, &col)
		mat.grid[row][col] = 1
	}

}

// GetUpdateMatrix update matrix with its current state
func (mat *Matrix) GetUpdateMatrix() [][]int {
	fMat := NewMatrix(mat.row, mat.col)

	for i := 0; i < mat.row; i++ {
		for j := 0; j < mat.col; j++ {

			liveNeighbours := mat.FindLiveNeighbours(i, j)
			fMat.grid[i][j] = mat.ChangeGameState(mat.grid[i][j], liveNeighbours)

		}
	}
	return fMat.grid

}

func (mat *Matrix) FindLiveNeighbours(row, col int) (liveNeighbours int) {
	rDimension := []int{1, -1, 0, 0, -1, 1, -1, 1}
	cDimension := []int{0, 0, 1, -1, -1, -1, 1, 1}
	liveNeighbours = 0

	for i := 0; i < 8; i++ {
		currRow := row + rDimension[i]
		currCol := col + cDimension[i]

		if currRow >= 0 && currRow < mat.row && currCol >= 0 && currCol < mat.col && mat.grid[currRow][currCol] == 1 {
			liveNeighbours++
		}
	}
	return
}

func (mat *Matrix) ChangeGameState(currState int, liveNeighbours int) (val int) {

	if currState == 1 {
		if liveNeighbours < 2 || liveNeighbours > 3 {
			return 0
		}

	} else {
		if liveNeighbours == 3 {
			return 1
		}
	}
	return currState
}
