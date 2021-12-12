package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay12Part1(t *testing.T) {
	t.Run("Test SolveDay12Part1", func(t *testing.T) {
		got := SolveDay12Part1(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`)
		expected := 10

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay12Part1", func(t *testing.T) {
		got := SolveDay12Part1(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`)
		expected := 19

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay12Part1", func(t *testing.T) {
		got := SolveDay12Part1(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`)
		expected := 226

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay12Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay12Part1(string(i))
	}
}

func TestSolveDay12Part2(t *testing.T) {
	t.Run("Test SolveDay12Part2", func(t *testing.T) {
		got := SolveDay12Part2(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`)
		expected := 36
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay12Part2", func(t *testing.T) {
		got := SolveDay12Part2(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`)
		expected := 103
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay12Part2", func(t *testing.T) {
		got := SolveDay12Part2(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`)
		expected := 3509
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

}

func BenchmarkSolveDay12Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay12Part2(string(i))
	}
}
