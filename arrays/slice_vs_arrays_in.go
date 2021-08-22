package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// array ile slice ın syntax farkı da cok farksızdır başlatırken eğer size belirlemezsen bu bir slicetır mesela

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(myArray[7])
	for i, n := range myArray {

		fmt.Printf("Index : %d\n", i)
		fmt.Println(n)
	}
	
	for i, n := range mySlice {

		fmt.Printf("Index of slice : %d\n", i)
		fmt.Println(n)
	}
}

/*

The main difference between an Array object in Go and a Slice object is that the Slice is effectively an abstraction built on-top of the Array.
They give us programmers a greater degree of flexibility when it comes to working with this particular data structure
 as they handle things like the automatic resizing of arrays whenever we hit a limit.
In most situations, you are not likely going to want to specify an exact length of an array so it’s likely that you’l end up just resorting to
Slices for the majority of your use cases.


Go'daki bir Array nesnesi ile bir Slice nesnesi arasındaki temel fark, Slice'ın etkin bir şekilde Dizinin üzerine inşa edilmiş bir soyutlama olmasıdır.
Bu özel veri yapısıyla çalışma söz konusu olduğunda programcılara daha fazla esneklik sağlıyorlar.
 bir sınıra ulaştığımızda dizilerin otomatik olarak yeniden boyutlandırılması gibi şeyleri ele aldıkları için.
Çoğu durumda, bir dizinin tam uzunluğunu belirtmek istemeyeceksiniz, bu nedenle muhtemelen sadece bir diziye başvuracaksınız.
Kullanım durumlarınızın çoğu için dilimler.

*/
