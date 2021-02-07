package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type byAge []person

// Len ...
func (a byAge) Len() int { return len(a) }

// Less ...
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

// Swap ...
func (a byAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})
	fmt.Println(people)
}
