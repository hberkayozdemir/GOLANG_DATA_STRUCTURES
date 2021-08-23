// burada biraz daha zor bir data structure ile karşı karşıyayız queue ve stackten sonra biraz daha korkutucu
// linked listlerin bir headi bir de taili var
// linked listi bir metal zincir gibi düşünün her çelik halka bir sonraki halkanın referansını tutart
// önceki yapımızda yani stackteki mantık ne idi last in first out lahmacun örneği vermiştim kebapçının son koyduğu lahmacun bizim ilk yiceğimizdir :D
// burdada aynı prensipte haraket ediyoruz

package main

import "fmt"

type LinkedList struct {
	Head *Node
	Size int
}

// Node - a node represent link in our linked list zincir halkası
type Node struct {
	Value string
	Next  *Node
}

// yeni bir halka ekler zincirin en başına
func (l *LinkedList) Insert(elem string) {
	node := Node{
		Value: elem,
		Next:  l.Head,
	}
	l.Head = &node
	l.Size++
}

// Lahmacun
//adana
// eklediğpimizi düşünün
/*
Lahmacun
adana -> lahmacun
Head -> adana
adana->lahmacun->nil
*/
// delete first removes the first element from our linked list

func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

// print all nodes on linked list
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v \n", current)
		current = current.Next
	}
}

// search -searches through every element of our linked list
// returns the firt element that matches the string otherwise it returns nil

func (l *LinkedList) Search(elem string) *Node {
	current := l.Head
	for current != nil {
		if current.Value == elem {
			return current
		}
		current = current.Next
	}
	return nil
}

// delete selected chain with strinbg var if mathches delete the value
func (l *LinkedList) Delete(elem string) {

	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Value == elem {
			previous.Next = current.Next
			l.Size--
		}

		previous = current
		current = current.Next
	}
}

func main() {

	fmt.Println("Singly Linked List")
	var ll LinkedList

	ll.Insert("Lahmacun")
	ll.List()
	ll.Insert("adana")
	ll.List()
	ll.DeleteFirst()
	ll.List()
	ll.Insert("Ali nazik")
	ll.Insert("izmir köfte")
	ll.Insert("Mercimek çorbası")
	ll.Insert("beyti")
	if element := ll.Search("beyti"); element != nil {
		fmt.Println("Its founded")
		fmt.Printf("%+v \n", element)
	} else {
		fmt.Println("Baba bulamadık")
	}
	fmt.Println("Beyti bulunduktan sonra...")
	ll.Delete("Lahmacun")
	ll.List()
	fmt.Println(ll.Size)
	ll.DeleteFirst()
	fmt.Println("List after delete first")
	ll.List()
	fmt.Println(ll.Size)

}
