package main

import "fmt"
import "sort"

type Organ struct {
	Name   string
	Weight Grams
}

func (o *Organ) String() string { return fmt.Sprintf("%v (%v)", o.Name, o.Weight) }

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

// separator

type Organs []*Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }

type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }

func main() {
	s := []*Organ{
		{"spleen", 162},
		{"pacreas", 131},
		{"liver", 1494},
		{"heart", 290},
		{"brain", 1340},
	}
	for _, o := range s {
		fmt.Println(o)
	}

	fmt.Println("before sort", s)

	// sort
	sort.Sort(ByWeight{s})
	fmt.Println("Organs by weight", s)

	sort.Sort(ByName{s})
	fmt.Println("Organs by name", s)

	// reverse
	fmt.Println("before reverse", s)

	sort.Sort(sort.Reverse(ByWeight{s}))
	fmt.Println("Organs by weight", s)

	sort.Sort(sort.Reverse(ByName{s}))
	fmt.Println("Organs by name", s)
}
