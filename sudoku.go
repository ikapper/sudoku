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

	err := board.init(numberstr)
	if err != nil {
		fmt.Println(err)
		return
	}

	board.calcAnswer(0)
	board.printAnswer()
}

func (b *Board) init(numstrs []string) error {
	if len(numstrs) != N*N {
		return fmt.Errorf("Inputted numbers' length must be 81.")
	}
	for i, ch := range numstrs {
		d, err := strconv.ParseInt(ch, 10, 0)
		if err != nil {
			panic(err)
		}
		if 0 < d && d <= 9 {
			r, c := i/N, i%N
			b.update(r, c, int(d-1), 1)
			b.doneCells = append(b.doneCells, i)
			b.locations[i] = int(d)
		}
	}
	return nil
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
