/*
	Given an m x n integer matrix `matrix`, if an element is 0, set its entire row and column to 0's.

	You must do it in place.

	Example 1:

	Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
	1 1 1
	1 0 1
	1 1 1
	Output: [[1,0,1],[0,0,0],[1,0,1]]

	Example 2:

	Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]


    m == matrix.length
    n == matrix[0].length
    1 <= m, n <= 200
    -231 <= matrix[i][j] <= 231 - 1

*/

/*
	iterate over levels (i, lvl)
	at each level,
		iterate, (i, n)
			if a 0 is found,
				save the index it's found at in `zero`
				save the index of the level in `blanks`

	iterate levels again (i, lvl)
		if `blanks[i]` set all level to 0s and skip to next
		iterate `zeroes` (_, idx)
			set lvl[idx] to 0
*/

package main

import "fmt"

func main() {
	inpt := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	//[[1,0,1],[0,0,0],[1,0,1]]
	setZeroes(inpt)
	fmt.Println(inpt)
}

func setZeroes(matrix [][]int) {
	var zrs []int
	rws := map[int]bool{}

	for li, lvl := range matrix {
		for ni, n := range lvl {
			if n == 0 {
				zrs = append(zrs, ni)
				rws[li] = true
			}
		}
	}

	for i, lvl := range matrix {
		if rws[i] {
			matrix[i] = make([]int, len(lvl))
			continue
		}
		for _, idx := range zrs {
			lvl[idx] = 0
		}
	}
}
