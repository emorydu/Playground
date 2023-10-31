package main

import "fmt"

func main() {
	u := []int{11, 12, 13, 14, 15}
	fmt.Println("array:", u)                                     // [11, 12, 13, 14, 15]
	s := u[1:3]                                                  // [11, 12]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=2, cap=4): [12, 13]

	s = append(s, 24)
	fmt.Println("after append 24, array:", u)                                     // [11, 12, 13, 24, 15]
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=3, cap=4): [12, 13, 24]

	s = append(s, 25)
	fmt.Println("after append 25, array:", u)                                     // [11, 12, 13, 24, 25]
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=4, cap=4): [12, 13, 24, 25]

	s = append(s, 26)
	fmt.Println("after append 26, array:", u)                                     // [11, 12, 13, 24, 25]
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=5, cap=8): [12, 13, 24, 25, 26]

	s[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", u) // after reassign 1st elem of slice, array: [11 12 13 24 25]

	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // after reassign 1st elem of slice, slice(len=5, cap=8): [22 13 24 25 26]

}
