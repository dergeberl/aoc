package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay24Part1(t *testing.T) {
	t.Run("Test SolveDay24Part1", func(t *testing.T) {
		i := stringListToSlice(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`)
		got := SolveDay24Part1(i)
		expected := 10

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay24Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay24Part1(input)
	}
}

func TestSolveDay24Part2(t *testing.T) {
	t.Run("Test SolveDay24Part2", func(t *testing.T) {
		i := stringListToSlice(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`)
		got := SolveDay24Part2(i)
		expected := 2208
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay24Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay24Part2(input)
	}
}

func TestStringListToSlice(t *testing.T) {
	t.Run("Test stringListToSlice with valid list", func(t *testing.T) {
		s := `string
string 2
string:3`
		got := stringListToSlice(s)
		expected := []string{"string", "string 2", "string:3"}
		if !equalString(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func equalString(a, b []string) bool {
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
