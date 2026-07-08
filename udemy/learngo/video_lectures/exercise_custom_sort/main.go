package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

type ByName []Person

func (bn ByAge) Len() int           { return len(bn) }
func (bn ByAge) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByAge) Less(i, j int) bool { return bn[i].Age < bn[j].Age }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	//Nope, don't need to convert the slice of Person (struct) to a slice of string
	//fmt.String(people)

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	//Nope, don't need to call Strings to convert, just note the methods used by the interface and tweak the one that matters (i.e. Less())
	//by replacing the field that we're sorting by. (The interface will apply the appropriate method based on the type)
	//sort.Strings(people)
	fmt.Println(people)

}
