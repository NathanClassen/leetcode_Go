package main

/*
	Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

    Each row must contain the digits 1-9 without repetition.
    Each column must contain the digits 1-9 without repetition.
    Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

	Note:

		A Sudoku board (partially filled) could be valid but is not necessarily solvable.
		Only the filled cells need to be validated according to the mentioned rules.

	

	Example 1:

	Input: board = 
	[["5","3",".",".","7",".",".",".","."]
	,["6",".",".","1","9","5",".",".","."]
	,[".","9","8",".",".",".",".","6","."]
	,["8",".",".",".","6",".",".",".","3"]
	,["4",".",".","8",".","3",".",".","1"]
	,["7",".",".",".","2",".",".",".","6"]
	,[".","6",".",".",".",".","2","8","."]
	,[".",".",".","4","1","9",".",".","5"]
	,[".",".",".",".","8",".",".","7","9"]]
	Output: true

	Example 2:

	Input: board = 
	[["8","3",".",".","7",".",".",".","."]
	,["6",".",".","1","9","5",".",".","."]
	,[".","9","8",".",".",".",".","6","."]
	,["8",".",".",".","6",".",".",".","3"]
	,["4",".",".","8",".","3",".",".","1"]
	,["7",".",".",".","2",".",".",".","6"]
	,[".","6",".",".",".",".","2","8","."]
	,[".",".",".","4","1","9",".",".","5"]
	,[".",".",".",".","8",".",".","7","9"]]
	Output: false
	Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

	

	Constraints:

		board.length == 9
		board[i].length == 9
		board[i][j] is a digit 1-9 or '.'.


	p

	ensure a board is valid sudoku (not necessarily solvable) by validating that of the existing entries, the following is true
		- no slice has duplicate entries
		- no 2 or more slices have the same entry at a given index
		- for each group of 3 consecutive rows, every 3 entries ((0,1,2), (3,4,5), (6,7,8)), contain only unique entires or '.'

	it seems likely that careful use of maps which will track which numbers have been seen should enable this to operate with O(n) time and O(n) space.
	given that the input is always the same size, we may as well say O(1)

	a single counter can track to location in all 9 rows.
	as it pass along, we will track the rows in groups of 3.
		- every singe row's elements will be tracked
		- every group of 3 rows will be tracked together with a map, which will essentially be cleared every 3 positions (this is to track the "boxes")

	
*/

import "fmt"

func main() {
	board1 := [][]byte{
		{5,3,'.','.',7,'.','.','.','.'},
		{6,'.','.',1,9,5,'.','.','.'},
		{'.',9,8,'.','.','.','.',6,'.'},
		{8,'.','.','.',6,'.','.','.',3},
		{4,'.','.',8,'.',3,'.','.',1},
		{7,'.','.','.',2,'.','.','.',6},
		{'.',6,'.','.','.','.',2,8,'.'},
		{'.','.','.',4,1,9,'.','.',5},
		{'.','.','.','.',8,'.','.',7,9},
	}

	board2 := [][]byte{
		{8,3,'.','.',7,'.','.','.','.'},
		{6,'.','.',1,9,5,'.','.','.'},
		{'.',9,8,'.','.','.','.',6,'.'},
		{8,'.','.','.',6,'.','.','.',3},
		{4,'.','.',8,'.',3,'.','.',1},
		{7,'.','.','.',2,'.','.','.',6},
		{'.',6,'.','.','.','.',2,8,'.'},
		{'.','.','.',4,1,9,'.','.',5},
		{'.','.','.','.',8,'.','.',7,9},
	}
	
	board3 := [][]byte{
		{5,3,'.','.',7,'.','.','.','.'},
		{6,5,'.',1,9,3,'.','.','.'},
		{'.',9,8,'.','.','.','.',6,'.'},
		{8,'.','.','.',6,'.','.','.',3},
		{4,'.','.',8,'.',5,'.','.',1},
		{7,'.','.','.',2,'.','.','.',6},
		{'.',6,'.','.','.','.',2,8,'.'},
		{'.','.','.',4,1,9,'.','.',5},
		{'.','.','.','.',8,'.','.',7,9},
	}

	// board4 := [][]byte{
	// 	{8,3,'.','.',7,'.','.','.','.'},
	// 	{6,'.','.',1,9,5,'.','.','.'},
	// 	{'.',9,8,'.','.','.','.',6,'.'},
	// 	{8,'.','.','.',6,'.','.','.',3},
	// 	{4,'.','.',8,'.',3,'.','.',1},
	// 	{7,'.','.','.',2,'.','.','.',6},
	// 	{'.',6,'.','.','.','.',2,8,'.'},
	// 	{'.','.','.',4,1,9,'.','.',5},
	// 	{'.','.','.','.',8,'.','.',7,9},
	// }

	fmt.Println(isValidSudoku(board1) == true)
	fmt.Println(isValidSudoku(board2) == false)
	fmt.Println(isValidSudoku(board3) == false)
}

func isValidSudoku(board [][]byte) bool {
    seen := make(map[int]map[byte]bool, 9)
	column := make(map[byte]bool, 9)
	boxKey := 0

	// initialize maps
	for i := 0; i < 9; i++ {
		seen[i] = make(map[byte]bool)
		if i % 3 == 0 {
			p := i
			if i == 0 {
				p = 1
			}
			seen[p*10] = make(map[byte]bool)
		}
	}
	
	for p := 0; p < 9; p++ {
		//	first we'll ensure that IF were in a new box (which will be the case for first iteration too),
		//		then we clear the box seen maps
		if p % 3 == 0 {
			clear(seen[10])
			clear(seen[30])
			clear(seen[60])
		}

		// check column, box, and single row uniqueness
		for rIdx, row := range board {
			elmt := row[p]
			if elmt != '.' {

				// column check
				if _, t := column[elmt]; t {
					return false
				} else {
					column[elmt] = true
				}

				// row check
				if _, t := seen[rIdx][elmt]; t {
					return false
				} else {
					seen[rIdx][elmt] = true
				}

				// box evaluation
				// identify which box to check
				if rIdx < 3 {
					boxKey = 10
				} else if rIdx < 6 {
					boxKey = 30
				} else {
					boxKey = 60
				}
				// check if the element at p in the given row has already been seen in the relevant box
				if _, t := seen[boxKey][elmt]; t {
					return false
				} else {
					seen[boxKey][elmt] = true
				}
			}
		}

		// column is verified unique, clear for next one
		clear(column)
		
	}

	//	all checks have passed, the board is valid
	return true
}
/*

	set pointer to 0
	iterate on pointer
	at each index:
		- ensure the element at that index is unique for all rows
		- check whether the element has been seen for each row
			- if not, add it to the 'seen' elements for that row
			- if seen already, return false for the whole input
		- on the group of 3 rows level
			- check that groups 'seen' map to ensure that the element is unique across the entire group of 3 for the current interval
			- if seen, return false for whole group
			- if not, add to seen for that group
		- if pointer % 3 == 0
			- then iteration has passed to a new box of 3
				- reset the 'seen' maps for the row groups
	- return true

	seen = {
		0 {}
		1 {}
		2 {}
		3 {}
		4 {}
		5 {}
		6 {}
		7 {}
		8 {}
		g1 {}
		g2 {}
		g3 {}
	}
*/