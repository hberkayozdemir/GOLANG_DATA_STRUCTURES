package main

import "fmt"

func main() {

	fmt.Println("Arrays")

	var myArray [8]int          // [0,0,0,0,0,0,0,0,0] default olarak int 0
	var myStringArray [8]string // [        ]  sanırım default olarak string " "
	var myByteArray [8]byte     // [0,0,0,0,0,0,0,0,0]
	fmt.Println(myArray)
	fmt.Println(myStringArray)
	fmt.Println(myByteArray)

	myArray2 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(myArray2)

	fmt.Println(len(myArray2))

	fmt.Println(myArray2[7])
	fmt.Println("##################################################################################################")
	fmt.Println("    ---             For loop  Started                        -----")

	for i, n := range myArray2 {
		fmt.Printf("Index: %d\n", i)
		fmt.Println(n)
	}

}

/*

array:   [1]    []    []    []    []    []    []    [] size 8
 index:  0     1     2     3     4     5     6     7
array[0]:= 1


 */ 