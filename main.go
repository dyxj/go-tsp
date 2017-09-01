package main

import (
	"fmt"
	b "go-tsp/base"
	"math/rand"
	"time"
	"reflect"
)

func main() {
	fmt.Println("Traveling sales person")
	//cityListContain()
}

func tspRandom() {
	fmt.Println("Traveling sales person - Standard Random")
}


// Functions below are for experimentation
func cityListContain() {
	c1 := b.City{}
	c1.SetLocation(10,20)
	c2 := b.City{}
	c2.SetLocation(30,40)
	c3 := b.City{}
	c3.SetLocation(50,60)

	c4 := b.City{}
	c4.SetLocation(30,40)

	cslice := make([]b.City,0,20)

	cslice = append(cslice,c1)
	cslice = append(cslice,c2)
	cslice = append(cslice,c3)

	fmt.Println(cslice)
	fmt.Println(cslice[0])

	for _, c := range cslice {
		if c == c4 {
			fmt.Println("found same", c)
		}
		if reflect.DeepEqual(c,c4) {
			fmt.Println("deep equal true", c)
		}
	}

}

func xfunc() {
	//rand.Seed(time.Now().Unix())

	x := 5-10
	fmt.Println(x)
	if x < 0 {
		x = -x
	}
	fmt.Println(x)

	c1 := b.City{}
	c1.SetRandomLocation()
	fmt.Println(c1)

	c2 := b.City{}
	c2.SetRandomLocation()
	fmt.Println(c2)

	fmt.Println(c1.DistanceTo(c2))

	random1()
	random2()
}

func random1() {
	fmt.Println("random1")
	r := rand.New(rand.NewSource(time.Now().Unix()))
	src := []int{1,2,3,4,5}
	dest := make([]int, len(src))
	perm := r.Perm(len(src))
	fmt.Println(perm)
	for i, v := range perm {
		dest[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest)

	dest2 := make([]int, len(src))
	perm2 := r.Perm(len(src))
	fmt.Println(perm2)
	for i, v := range perm2 {
		dest2[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest2)
}

func random2() {
	fmt.Println("random2")
	src := []int{1,2,3,4,5}
	dest := make([]int, len(src))
	perm := rand.Perm(len(src))
	fmt.Println(perm)
	for i, v := range perm {
		dest[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest)

	dest2 := make([]int, len(src))
	perm2 := rand.Perm(len(src))
	fmt.Println(perm2)
	for i, v := range perm2 {
		dest2[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest2)
}
