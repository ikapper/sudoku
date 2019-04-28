package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestSudoku(t *testing.T) {
	buf := &bytes.Buffer{}
	out = buf
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

	// Check printed format string
	ansFContent, err := ioutil.ReadFile("answer01f.txt")
	if err != nil {
		log.Fatal(err)
	}
	expectedAnsFStr := string(ansFContent)
	b.printAnswer()
	ans = buf.String()
	if expectedAnsFStr != ans {
		t.Errorf("Incorrect answer format")
	}
}

func TestErrorInit(t *testing.T) {
	b := &Board{}
	err := b.init([]string{"1", "2"})
	if err == nil {
		t.Errorf("Init must be failed")
	}
}

func TestMainFunc(t *testing.T) {
	inContent, err := ioutil.ReadFile("input03.txt")
	if err != nil {
		log.Fatal(err)
	}
	in = strings.NewReader(string(inContent))
	buf := &bytes.Buffer{}
	out = buf
	main()

	ansContent, err := ioutil.ReadFile("answer03f.txt")
	if err != nil {
		log.Fatal(err)
	}
	expectedAns := string(ansContent)
	gotAns := buf.String()
	if expectedAns != gotAns {
		t.Errorf("Failed to solve")
	}
}

func TestFlattenOutput(t *testing.T) {
	inContent, err := ioutil.ReadFile("input01.txt")
	if err != nil {
		log.Fatal(err)
	}
	in = strings.NewReader(string(inContent))
	buf := &bytes.Buffer{}
	out = buf

	flag.Set("flatten", "true")
	main()

	ansContent, err := ioutil.ReadFile("answer01.txt")
	if err != nil {
		log.Fatal(err)
	}
	expectedAns := string(ansContent)
	gotAns := buf.String()
	if expectedAns != gotAns {
		t.Errorf("Wrong format, expected: %v, got: %v", expectedAns, gotAns)
	}
}
