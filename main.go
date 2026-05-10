package main

import (
	"fmt"

	"github.com/XEQTIONR/gollection/gollection"
)

func main() {

	collection := gollection.Init[int](10, 20, 30, 40, 50)
	collectionFromArray := gollection.InitFromSlice[int]([]int{10, 20, 30, 40, 50})

	fmt.Println("Collection: ", collection, 1)
	fmt.Println("Collection: ", collectionFromArray, 2)

	after, err := collection.After(20)
	after2, err2 := collection.After(50)

	if err == nil {
		fmt.Println("After 30: ", *after)
	}
	if err2 == nil {
		fmt.Println("After 50: ", *after2)
	}

	// fmt.Println("After 50: ", after2)
}
