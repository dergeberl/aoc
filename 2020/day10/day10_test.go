package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay10Part1(t *testing.T) {
	t.Run("Test SolveDay10Part1", func(t *testing.T) {
		i := intListToSlice(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
		got := SolveDay10Part1(i)
		expected := 22 * 10

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay10Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := intListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay10Part1(input)
	}
}

func TestSolveDay10Part2(t *testing.T) {
	t.Run("Test SolveDay10Part2", func(t *testing.T) {
		i := intListToSlice(`16
10
15
5
1
11
7
19
6
12
4`)
		got := SolveDay10Part2(i)
		expected := int64(8)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay10Part2", func(t *testing.T) {
		i := intListToSlice(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
		got := SolveDay10Part2(i)
		expected := int64(19208)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay10Part2 latest not used", func(t *testing.T) {
		i := intListToSlice(`16
10
15
5
1
11
7
19
6
12
4
25`)
		got := SolveDay10Part2(i)
		expected := int64(8)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay10Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := intListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay10Part2(input)
	}
}

func TestSolveDay10Part2Recursive(t *testing.T) {
	t.Run("Test SolveDay10Part2Recursive", func(t *testing.T) {
		i := intListToSlice(`16
10
15
5
1
11
7
19
6
12
4`)
		got := SolveDay10Part2Recursive(i)
		expected := int64(8)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay10Part2Recursive", func(t *testing.T) {
		i := intListToSlice(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)
		got := SolveDay10Part2Recursive(i)
		expected := int64(19208)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay10Part2Recursive latest not used", func(t *testing.T) {
		i := intListToSlice(`16
10
15
5
1
11
7
19
6
12
4
25`)
		got := SolveDay10Part2Recursive(i)
		expected := int64(8)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay10Part2Recursive(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := intListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay10Part2Recursive(input)
	}
}

func TestCountWays(t *testing.T) {
	t.Run("Test countWays", func(t *testing.T) {
		i := intListToSlice(`1
4
5
6
7
10
11
12
15
16
19`)
		wayCache = make(map[int]int64)
		got := countWays(i)
		expected := int64(8)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

}

func TestIntListToSlice(t *testing.T) {
	t.Run("Test intListToSlice with valid list", func(t *testing.T) {
		s := `123
1234
000
00123`
		got := intListToSlice(s)
		expected := []int{123, 1234, 0, 123}
		if !equalInt(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test intListToSlice with invalid list", func(t *testing.T) {
		s := `123
1234a
000
00123`
		got := intListToSlice(s)
		expected := []int{}
		if !equalInt(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func equalInt(a, b []int) bool {
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
