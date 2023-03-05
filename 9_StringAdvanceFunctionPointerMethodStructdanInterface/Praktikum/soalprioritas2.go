package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

func (s Student) String() string {
	return fmt.Sprintf("%s (%d)", s.Name, s.Score)
}

func (s Student) sScore() float64 {
	return float64(s.Score)
}

type Students []Student

func (s Students) AverageScore() float64 {
	n := len(s)
	var sum float64
	for _, student := range s {
		sum += student.sScore()
	}
	return sum / float64(n)
}

func (s Students) MinScore() Student {
	min := math.Inf(1)
	var minIdx int
	for i, student := range s {
		if student.sScore() < min {
			min = student.sScore()
			minIdx = i
		}
	}
	return s[minIdx]
}

func (s Students) MaxScore() Student {
	max := math.Inf(-1)
	var maxIdx int
	for i, student := range s {
		if student.sScore() > max {
			max = student.sScore()
			maxIdx = i
		}
	}
	return s[maxIdx]
}

func main() {
	var students Students

	input1 := Student{
		"Rizky",
		80,
	}

	input2 := Student{
		"Kobar",
		75,
	}

	input3 := Student{
		"Ismail",
		70,
	}

	input4 := Student{
		"Umam",
		75,
	}

	input5 := Student{
		"Yopan",
		60,
	}

	students = append(students, input1)
	students = append(students, input2)
	students = append(students, input3)
	students = append(students, input4)
	students = append(students, input5)

	fmt.Println("Average Score:", students.AverageScore())
	fmt.Println("Min Score:", students.MinScore())
	fmt.Println("Max Score:", students.MaxScore())

}
