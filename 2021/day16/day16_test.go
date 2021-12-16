package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay16Part1(t *testing.T) {
	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		got := SolveDay16Part1(`D2FE28`)
		expected := 6

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		got := SolveDay16Part1(`38006F45291200`)
		expected := 9

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		got := SolveDay16Part1(`EE00D40C823060`)
		expected := 14

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		got := SolveDay16Part1(`8A004A801A8002F478`)
		expected := 16

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		got := SolveDay16Part1(`620080001611562C8802118E34`)
		expected := 12

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		got := SolveDay16Part1(`C0015000016115A2E0802F182340`)
		expected := 23

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		got := SolveDay16Part1(`A0016C880162017C3686B18A3D4780`)
		expected := 31

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay16Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay16Part1(string(i))
	}
}

func TestSolveDay16Part2(t *testing.T) {
	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`D2FE28`)
		expected := 2021
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`C200B40A82`)
		expected := 3
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`04005AC33890`)
		expected := 54
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`880086C3E88112`)
		expected := 7
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`CE00C43D881120`)
		expected := 9
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`D8005AC2A8F0`)
		expected := 1
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`F600BC2D8F`)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`9C005AC2F8F0`)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		got := SolveDay16Part2(`9C0141080250320F1802104A08`)
		expected := 1
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

}

func BenchmarkSolveDay16Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay16Part2(string(i))
	}
}
