package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay5Part1(t *testing.T) {
	t.Run("Test SolveDay5Part1 with valid list", func(t *testing.T) {
		i := listToSlice(`BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`)
		got := SolveDay5Part1(i)
		expected := 820

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay5Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay5Part1(input)
	}
}

func TestSolveDay5Part2(t *testing.T) {
	t.Run("Test SolveDay5Part2", func(t *testing.T) {
		i := listToSlice(`BFFFBBFRRR
FFFBBBFRRR
FFFBBBFLLL
FFFBBBFLLR
FFFBBBFLRL
FFFBBBFLRR
FFFBBBFRLR
FFFBBBFRRL
FFFBBBFRRR
BBFFBBFRLL`)
		got := SolveDay5Part2(i)
		expected := 116
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay5Part2 with no seed match", func(t *testing.T) {
		i := listToSlice(`BFFFBBFRRR
FFFBBBFRRR
FFFBBBFLLL
FFFBBBFLLR
FFFBBBFLRL
FFFBBBFLRR
FFFBBBFRLR
FFFBBBFRRL
FFFBBBFRRR
BBFFBBFRLL
FFFBBBFRLL`)
		got := SolveDay5Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay5Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay5Part2(input)
	}
}

func TestGetSeatValue(t *testing.T) {
	t.Run("Test getSeatValue", func(t *testing.T) {
		row, column, seatid := getSeatValue("BFFFBBFRRR")
		expectedRow := 70
		expectedColumn := 7
		expectedSeatId := 567
		if row != expectedRow || column != expectedColumn || seatid != expectedSeatId {
			t.Errorf("expected '%d' '%d' '%d' but got '%d' '%d' '%d'", expectedRow, expectedColumn, expectedSeatId, row, column, seatid)
		}
	})
	t.Run("Test getSeatValue", func(t *testing.T) {
		row, column, seatid := getSeatValue("FFFBBBFRRR")
		expectedRow := 14
		expectedColumn := 7
		expectedSeatId := 119
		if row != expectedRow || column != expectedColumn || seatid != expectedSeatId {
			t.Errorf("expected '%d' '%d' '%d' but got '%d' '%d' '%d'", expectedRow, expectedColumn, expectedSeatId, row, column, seatid)
		}
	})
	t.Run("Test getSeatValue", func(t *testing.T) {
		row, column, seatid := getSeatValue("BBFFBBFRLL")
		expectedRow := 102
		expectedColumn := 4
		expectedSeatId := 820
		if row != expectedRow || column != expectedColumn || seatid != expectedSeatId {
			t.Errorf("expected '%d' '%d' '%d' but got '%d' '%d' '%d'", expectedRow, expectedColumn, expectedSeatId, row, column, seatid)
		}
	})
	t.Run("Test getSeatValue", func(t *testing.T) {
		row, column, seatid := getSeatValue("BBBBBBBRRR")
		expectedRow := 127
		expectedColumn := 7
		expectedSeatId := 1023
		if row != expectedRow || column != expectedColumn || seatid != expectedSeatId {
			t.Errorf("expected '%d' '%d' '%d' but got '%d' '%d' '%d'", expectedRow, expectedColumn, expectedSeatId, row, column, seatid)
		}
	})
	t.Run("Test getSeatValue", func(t *testing.T) {
		row, column, seatid := getSeatValue("FFFBBBFRLL")
		expectedRow := 14
		expectedColumn := 4
		expectedSeatId := 116
		if row != expectedRow || column != expectedColumn || seatid != expectedSeatId {
			t.Errorf("expected '%d' '%d' '%d' but got '%d' '%d' '%d'", expectedRow, expectedColumn, expectedSeatId, row, column, seatid)
		}
	})
	t.Run("Test getSeatValue with invalid seat", func(t *testing.T) {
		row, column, seatid := getSeatValue("BBBBBBBRRS")
		expectedRow := 0
		expectedColumn := 0
		expectedSeatId := 0
		if row != expectedRow || column != expectedColumn || seatid != expectedSeatId {
			t.Errorf("expected '%d' '%d' '%d' but got '%d' '%d' '%d'", expectedRow, expectedColumn, expectedSeatId, row, column, seatid)
		}
	})
	t.Run("Test getSeatValue with invalid seat length", func(t *testing.T) {
		row, column, seatid := getSeatValue("BBBBBBBRRSS")
		expectedRow := 0
		expectedColumn := 0
		expectedSeatId := 0
		if row != expectedRow || column != expectedColumn || seatid != expectedSeatId {
			t.Errorf("expected '%d' '%d' '%d' but got '%d' '%d' '%d'", expectedRow, expectedColumn, expectedSeatId, row, column, seatid)
		}
	})
}

func BenchmarkGetSeatValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _, _ = getSeatValue("BBBBBBBRRR")
		_, _, _ = getSeatValue("FFFFFFFLLL")
		_, _, _ = getSeatValue("SBBBBBBRRR")
		_, _, _ = getSeatValue("BBBBBBBRRS")

	}
}

func TestGetSeatValueFor(t *testing.T) {
	t.Run("Test getSeatID", func(t *testing.T) {
		seatid := getSeatID("BFFFBBFRRR")
		expectedSeatId := 567
		if seatid != expectedSeatId {
			t.Errorf("expected '%d' but got '%d'", expectedSeatId, seatid)
		}
	})
	t.Run("Test getSeatID", func(t *testing.T) {
		seatid := getSeatID("FFFBBBFRRR")
		expectedSeatId := 119
		if seatid != expectedSeatId {
			t.Errorf("expected '%d' but got '%d'", expectedSeatId, seatid)
		}
	})
	t.Run("Test getSeatID", func(t *testing.T) {
		seatid := getSeatID("BBFFBBFRLL")
		expectedSeatId := 820
		if seatid != expectedSeatId {
			t.Errorf("expected '%d' but got '%d'", expectedSeatId, seatid)
		}
	})
	t.Run("Test getSeatID", func(t *testing.T) {
		seatid := getSeatID("BBBBBBBRRR")
		expectedSeatId := 1023
		if seatid != expectedSeatId {
			t.Errorf("expected '%d' but got '%d'", expectedSeatId, seatid)
		}
	})
	t.Run("Test getSeatID", func(t *testing.T) {
		seatid := getSeatID("FFFBBBFRLL")
		expectedSeatId := 116
		if seatid != expectedSeatId {
			t.Errorf("expected '%d' but got '%d'", expectedSeatId, seatid)
		}
	})
	t.Run("Test getSeatID with invalid seat", func(t *testing.T) {
		seatid := getSeatID("BBBBBBBRRS")
		expectedSeatId := 0
		if seatid != expectedSeatId {
			t.Errorf("expected '%d' but got '%d'", expectedSeatId, seatid)
		}
	})
	t.Run("Test getSeatID with invalid seat length", func(t *testing.T) {
		seatid := getSeatID("BBBBBBBRRSS")
		expectedSeatId := 0
		if seatid != expectedSeatId {
			t.Errorf("expected '%d' but got '%d'", expectedSeatId, seatid)
		}
	})
}

func BenchmarkGetSeatValueFor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = getSeatID("BBBBBBBRRR")
		_ = getSeatID("FFFFFFFLLL")
		_ = getSeatID("SBBBBBBRRR")
		_ = getSeatID("BBBBBBBRRS")

	}
}

func TestListToSlice(t *testing.T) {
	t.Run("Test ListToSlice with valid list", func(t *testing.T) {
		s := `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`
		got := listToSlice(s)
		expected := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
		if !equal(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
