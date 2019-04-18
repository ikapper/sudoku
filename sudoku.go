package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// N is the number of a side.
	N = 9
)

// Board is board of sudoku
type Board struct {
	board [N * N]cell
	// locations keeps current answer state(1-based)
	locations [N * N]int
	// doneCells has indecies having init value
	doneCells []int
}

type cell struct {
	nums [N]int
}

func main() {
	board := &Board{}

	var n string
	fmt.Scanf("%s", &n)
	numberstr := strings.Split(n, "")

	for i, ch := range numberstr {
		d, err := strconv.ParseInt(ch, 10, 0)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("No.%v: %v\n", i, d)
		if d != 0 {
			r, c := i/N, i%N
			board.update(r, c, int(d-1), 1)
			board.doneCells = append(board.doneCells, i)
			board.locations[i] = int(d)
		}
	}

	board.calcAnswer(0)
	board.printAnswer()
}

func (b *Board) printAnswer() {
	for r := 0; r < N; r++ {
		slice := b.locations[r*N : r*N+N]
		fmt.Println(slice)
	}
}

func (b *Board) update(r, c, n, diff int) {
	// add diff to specified n
	// row and column direction
	for i := 0; i < N; i++ {
		b.board[r*N+i].nums[n] += diff
		b.board[i*N+c].nums[n] += diff
	}
	// small 3x3 area
	sr, sc := r/3, c/3
	sn := N / 3
	for y := 0; y < sn; y++ {
		for x := 0; x < sn; x++ {
			idx := (sn*sr+y)*N + (sn*sc + x)
			b.board[idx].nums[n] += diff
		}
	}
}

func (b *Board) calcAnswer(idx int) bool {
	if idx >= N*N {
		return true
	}
	for _, v := range b.doneCells {
		if idx == v {
			// skip this index
			return b.calcAnswer(idx + 1)
		}
	}
	for i := 0; i < N; i++ {
		if b.board[idx].nums[i] == 0 {
			// try to put int
			b.locations[idx] = i + 1
			r, c := idx/N, idx%N
			b.update(r, c, i, 1)
			if !b.calcAnswer(idx + 1) {
				// back
				b.update(r, c, i, -1)
			} else {
				return true
			}
		}
	}
	return false
}
