package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay18Part1(t *testing.T) {
	t.Run("Test SolveDay18Part1", func(t *testing.T) {
		got := SolveDay18Part1(`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`)
		expected := 4140

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay18Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay18Part1(string(i))
	}
}

func TestSolveDay18Part2(t *testing.T) {
	t.Run("Test SolveDay18Part2", func(t *testing.T) {
		got := SolveDay18Part2(`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`)
		expected := 3993
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay18Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay18Part2(string(i))
	}
}

func TestExplode(t *testing.T) {
	t.Run("Test explode", func(t *testing.T) {
		n := parseInput(`[[[[[9,8],1],2],3],4]`)[0]
		n.reduceExplode()
		got := n.toString()
		expected := "[[[[0,9],2],3],4]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test explode", func(t *testing.T) {
		n := parseInput(`[7,[6,[5,[4,[3,2]]]]]`)[0]
		n.reduceExplode()
		got := n.toString()
		expected := "[7,[6,[5,[7,0]]]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test explode", func(t *testing.T) {
		n := parseInput(`[[6,[5,[4,[3,2]]]],1]`)[0]
		n.reduceExplode()
		got := n.toString()
		expected := "[[6,[5,[7,0]]],3]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test explode", func(t *testing.T) {
		n := parseInput(`[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]`)[0]
		n.reduceExplode()
		got := n.toString()
		expected := "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test explode", func(t *testing.T) {
		n := parseInput(`[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]`)[0]
		n.reduceExplode()
		got := n.toString()
		expected := "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test explode", func(t *testing.T) {
		n := parseInput(`[[[[4,0],[5,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]`)[0]
		n.reduceExplode()
		got := n.toString()
		expected := "[[[[4,0],[5,4]],[[0,[7,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test explode", func(t *testing.T) {
		n := parseInput(`[[[[4,0],[5,4]],[[7,7],[6,0]]],[[7,[5,5]],[[0,[[5,6],3]],[[6,3],[8,8]]]]]`)[0]
		n.reduceExplode()
		got := n.toString()
		expected := "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[7,[5,5]],[[5,[0,9]],[[6,3],[8,8]]]]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("Test reduce", func(t *testing.T) {
		n := parseInput(`[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]`)[0]
		n.reduce()
		got := n.toString()
		expected := "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test reduce", func(t *testing.T) {
		n := parseInput(`[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]`)[0]
		n.reduce()
		got := n.toString()
		expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Test add", func(t *testing.T) {
		n1 := parseInput(`[1,2]`)[0]
		n2 := parseInput(`[3,4]`)[0]
		n := n1.add(n2)
		got := n.toString()
		expected := "[[1,2],[3,4]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test add", func(t *testing.T) {
		n1 := parseInput(`[[1,2],[3,4]]`)[0]
		n2 := parseInput(`[6,7]`)[0]
		n := n1.add(n2)
		got := n.toString()
		expected := "[[[1,2],[3,4]],[6,7]]"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestMagnitude(t *testing.T) {
	t.Run("Test magnitude", func(t *testing.T) {
		n := parseInput(`[1,9]`)[0]
		got := n.magnitude()
		expected := 21
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test magnitude", func(t *testing.T) {
		n := parseInput(`[9,1]`)[0]
		got := n.magnitude()
		expected := 29
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test magnitude", func(t *testing.T) {
		n := parseInput(`[[1,2],[[3,4],5]]`)[0]
		got := n.magnitude()
		expected := 143
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
