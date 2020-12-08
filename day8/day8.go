package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay8Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay8Part2(stringListToSlice(input)))
	fmt.Printf("Part 1r: %v\n", SolveDay8Part1r(stringListToSlice(input)))
	fmt.Printf("Part 2r: %v\n", SolveDay8Part2r(stringListToSlice(input)))

}

//SolveDay8Part1 returns the number of the latest number before go to loop
func SolveDay8Part1(i []string) (s int) {
	checked := make(map[int]bool)
	cur := 0
	for {
		if checked[cur] {

					return

		}
		checked[cur] = true
		if strings.HasPrefix(i[cur], "acc "){
			t := strings.TrimPrefix(i[cur], "acc ")
			temp, err := strconv.Atoi(t)
			if err != nil {
				return 0
			}
			s = s + temp
		} else if strings.HasPrefix(i[cur], "jmp "){
			t := strings.TrimPrefix(i[cur], "jmp ")
			temp, err := strconv.Atoi(t)
			if err != nil {
				return 0
			}
			cur += temp
			continue
		}
		cur ++
	}
}

//SolveDay8Part1r returns the number of the latest number before go to loop - refactor
func SolveDay8Part1r(i []string) (s int) {
	checked := make(map[int]bool)
	tasks := getTasks(i)
	cur := 0
	for {
		if checked[cur] {

			return

		}
		checked[cur] = true
		for task, int := range tasks[cur]{
			if task == "acc"{
				s +=int
				cur ++
				break
			}
			if task == "jmp" {
				cur += int
                break
			}
			cur ++
		}
	}
}

//SolveDay8Part2 returns the value after finish with exact one change
func SolveDay8Part2(i []string) (s int) {
	input := i
	checked := make(map[int]bool)
	cur := 0
	changed:= 0
	for {
		if len(input) <= cur {
			return s
		}
		if checked[cur] {
			s, cur, checked = 0, 0, make(map[int]bool)
			changed++
		}
		checked[cur] = true
		if strings.HasPrefix(input[cur], "acc ") {
			temp, err := strconv.Atoi(strings.TrimPrefix(input[cur], "acc "))
			if err != nil {
				return 0
			}
			s += temp
			cur ++
		} else if strings.HasPrefix(input[cur], "jmp") && cur != changed || (strings.HasPrefix(input[cur], "nop") && cur == changed){
			jump, err := strconv.Atoi(strings.TrimPrefix(input[cur], "jmp "))
			if err != nil {
				jump, err = strconv.Atoi(strings.TrimPrefix(input[cur], "nop "))
				if err != nil {
					return 0
				}
			}
			cur += jump
		} else if strings.HasPrefix(input[cur], "nop") && cur != changed|| (strings.HasPrefix(input[cur], "jmp") && cur == changed){
			cur++
		}


	}
}

//SolveDay8Part2r returns the value after finish with exact one change - refactor
func SolveDay8Part2r(i []string) (s int) {
	input := i
	checked := make(map[int]bool)
	cur := 0
	changed:= 0
	tasks := getTasks(i)
	for {
		if len(input) <= cur {
			return s
		}
		if checked[cur] {
			s, cur, checked = 0, 0, make(map[int]bool)
			changed++
		}
		checked[cur] = true
		for task, int := range tasks[cur]{
			if task == "acc"{
				s +=int
				cur ++
				break
			}
			if task == "jmp"  && cur != changed || task == "nop" && cur == changed {
				cur += int
				break
			}
			if task == "nop"  && cur != changed || task == "jmp" && cur == changed {
				cur++
				break
			}
		}



	}
}

//getTasks returns a map with the tasks
func getTasks(input []string)  (tasks map[int]map[string]int) {
	tasks = make(map[int]map[string]int)
	for i, task := range input {
		splitTask := strings.Split(task, " ")
		steps, err := strconv.Atoi(splitTask[1])
		if err != nil {
			return nil
		}
		tasks[i] = map[string]int{}
		tasks[i][splitTask[0]] = steps
	}

	return tasks
}

//Helper functions
//stringListToSlice converts the list of strings (each string one row) to a slice
func stringListToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}

//intListToSlice converts the list of numbers (each number one row) to a slice
func intListToSlice(list string) (i []int) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		i = append(i, lineInt)
	}
	return
}