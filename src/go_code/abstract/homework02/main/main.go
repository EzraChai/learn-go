package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// Len is the number of elements in the collection.
func (heroSlice HeroSlice) Len() int {
	return len(heroSlice)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (heroSlice HeroSlice) Less(i, j int) bool {
	return heroSlice[i].Age < heroSlice[j].Age
}

// Swap swaps the elements with indexes i and j.
func (heroSlice HeroSlice) Swap(i, j int) {
	heroSlice[i], heroSlice[j] = heroSlice[j], heroSlice[i]
}

func main() {

	rand.Seed(time.Now().Unix())

	var intSlice = []int{0, -1, 10, 7, 90}

	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroes = HeroSlice{}
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: "Hero" + strconv.Itoa(i),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)

	//	Sort the heros based on age
	sort.Sort(heroes)

	//	[{Superman 24} {Batman 30} {Ironman 40}]
	fmt.Println(heroes)
}
