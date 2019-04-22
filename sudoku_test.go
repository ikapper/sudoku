package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestSudoku(t *testing.T) {
	b := &Board{}
	content, err := ioutil.ReadFile("input01.txt")
	if err != nil {
		log.Fatal(err)
	}
	numstr := string(content)
	numstrs := strings.Split(numstr, "")
	if err := b.init(numstrs); err != nil {
		t.Errorf("Board failed to initialize itself by %v", numstrs)
	}
	if len(b.doneCells) != 44 {
		t.Errorf("Board failed to scan cells. Scanned cells must be %v but %v", 44, len(b.doneCells))
	}
	copyDoneCells := make([]int, len(b.doneCells))
	copy(copyDoneCells, b.doneCells)

	b.calcAnswer(0)

	for i := 0; i < len(b.doneCells); i++ {
		if b.doneCells[i] != copyDoneCells[i] {
			t.Errorf("Board.doneCells should not be changed.")
			t.FailNow()
		}
	}

	ansContent, err := ioutil.ReadFile("answer01.txt")
	if err != nil {
		log.Fatal(err)
	}
	expectedAnsStr := string(ansContent)
	builder := &strings.Builder{}
	for _, i := range b.locations {
		builder.WriteString(strconv.Itoa(i))
	}
	ans := builder.String()
	if expectedAnsStr != ans {
		t.Errorf("Incorrect answer; expected: %v, but %v", expectedAnsStr, ans)
	}
}
