package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay4Part1(t *testing.T) {
	t.Run("Test SolveDay4Part1", func(t *testing.T) {
		i := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
		got := SolveDay4Part1(i)
		expected := 2

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay4Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay4Part1(input)
	}
}

func TestSolveDay4Part2(t *testing.T) {
	t.Run("Test SolveDay4Part2", func(t *testing.T) {
		i := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719

iyr:2010 hgt:158cmm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
		got := SolveDay4Part2(i)
		expected := 4
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay4Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay4Part2(input)
	}
}

func TestCheckRequired(t *testing.T) {
	t.Run("Test checkRequired with valid fields", func(t *testing.T) {
		i := map[string]string{
			"byr": "a",
			"iyr": "a",
			"eyr": "a",
			"hgt": "a",
			"hcl": "a",
			"ecl": "a",
			"pid": "a",
		}
		got := checkRequired(i)
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with one field missing", func(t *testing.T) {
		i := map[string]string{
			"byr": "a",
			"iyr": "a",
			"hgt": "a",
			"hcl": "a",
			"ecl": "a",
			"pid": "a",
		}
		got := checkRequired(i)
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}


func TestCheckValue(t *testing.T) {
	t.Run("Test checkRequired with byr valid value", func(t *testing.T) {
		got := checkValue("2000", "byr")
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with iyr valid value", func(t *testing.T) {
		got := checkValue("2019", "iyr")
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with eyr valid value", func(t *testing.T) {
		got := checkValue("2025", "eyr")
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with hgt valid value", func(t *testing.T) {
		got := checkValue("180cm", "hgt")
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with hcl valid value", func(t *testing.T) {
		got := checkValue("#123546", "hcl")
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with ecl valid value", func(t *testing.T) {
		got := checkValue("blu", "ecl")
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with pid valid value", func(t *testing.T) {
		got := checkValue("012345678", "pid")
		expected := true
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Test checkRequired with byr invalid value", func(t *testing.T) {
		got := checkValue("2020", "byr")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with iyr invalid value", func(t *testing.T) {
		got := checkValue("2021", "iyr")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with eyr invalid value", func(t *testing.T) {
		got := checkValue("2019", "eyr")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with hgt invalid value", func(t *testing.T) {
		got := checkValue("80in", "hgt")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with hcl invalid value", func(t *testing.T) {
		got := checkValue("#123u46", "hcl")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with ecl invalid value", func(t *testing.T) {
		got := checkValue("blue", "ecl")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with pid invalid value", func(t *testing.T) {
		got := checkValue("0123456789", "pid")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test checkRequired with invalid key", func(t *testing.T) {
		got := checkValue("", "pssid")
		expected := false
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
