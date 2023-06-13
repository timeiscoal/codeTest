package main

import (
	"fmt"
	"sort"
)

type data struct {
	name string
	age  int
}

func main() {

	var a []data
	a = append(a, data{"Khan", 10})
	a = append(a, data{"KY", 20})

	sort.Slice(a, func(i, j int) bool {
		return a[i].age < a[j].age
	})

	fmt.Println(a)

	fmt.Println(255 / 100)

}
