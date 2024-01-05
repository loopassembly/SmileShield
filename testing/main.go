package main

import "fmt"

type Node struct {
    Data int
    Next *Node
}

type LinkedList struct {
    Head *Node
}

func (ll *LinkedList) Append(data int) {
    newNode := &Node{Data: data, Next: nil}
    if ll.Head == nil {
        ll.Head = newNode
        return
    }
    current := ll.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode
}

func (ll *LinkedList) Display() {
    current := ll.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    linkedList := &LinkedList{}
    linkedList.Append(1)
    linkedList.Append(2)
    linkedList.Append(3)

    fmt.Println("linklNode")
    linkedList.Display()
}
