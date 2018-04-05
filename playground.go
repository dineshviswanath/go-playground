package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Variable declaration
	foo := 5
	bar := true
	name := `dinesh`

	// Parameter reading
	msg := os.Args[1]
	fmt.Println(msg)

	// Formatting
	fmt.Printf("foo %+v \n", foo)
	fmt.Printf("var a %T = %+v\n", foo, foo)
	fmt.Println("bar := ", bar)
	fmt.Println("name := ", name)

	// Array
	names := [4]string{"Ram", "Raghu", "Din", "Raj"}

	// Classic for loop
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Range loop
	for i, n := range names {
		fmt.Printf("%d  :  %s\n", i, n)
	}

	// Slice
	sliceA := []string{}
	sliceB := []string{"Raj", "Raghu"}
	sliceA = append(sliceA, "Dinesh", "Arun")
	// Append another slice
	sliceA = append(sliceA, sliceB...)

	for i, item := range sliceA {
		fmt.Printf("From Slice %d : %s\n", i, item)
	}

	fmt.Println(len(sliceA))
	fmt.Println(cap(sliceA))

	sliceA = append(sliceA, "AnotherDinesh", "AnotherArun")

	fmt.Println(len(sliceA))
	fmt.Println(cap(sliceA))

	// Another way of creating slice
	makeSliceA := make([]string, 10, 30)

	fmt.Println(len(makeSliceA))
	fmt.Println(cap(makeSliceA))
	makeSliceA = append(makeSliceA, sliceA...)

	fmt.Println(len(makeSliceA))
	fmt.Println(cap(makeSliceA))

	a := []string{"A", "B", "C", "D"}

	fmt.Println(a)
	fmt.Println(a[1:4])
	fmt.Println(a[:3])
	fmt.Println(a[1:])

	subsetSlice := a[1:2]
	fmt.Println("Subset: Before mutation", subsetSlice)
	// Mutate the subset
	for i, s := range subsetSlice {
		subsetSlice[i] = strings.ToLower(s)
	}

	fmt.Println("Subset: After mutation", subsetSlice)

	// Map
	mapA := map[int]string{}

	mapA[1] = "A"
	mapA[2] = "B"
	mapA[3] = "C"
	mapA[4] = "D"

	fmt.Println(mapA)

	// Key value in range
	for k, v := range mapA {
		fmt.Println(k, v)
	}

	// Check if key exists: version1
	k, ok := mapA[1]
	if ok {
		fmt.Println(k)
	}

	// Check if key exists: version2
	if k, ok := mapA[1000]; ok {
		fmt.Println(k)
	}
}
