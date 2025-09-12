package main
/**
import "fmt"

type Node struct{ 
	data int
	next *Node
}

func appendNode(head *Node, data int) *Node {
    newNode := &Node{data: data}
    if head == nil {
        return newNode
    }
    curr := head
    for curr.next != nil {
        curr = curr.next
    }
    curr.next = newNode
    return head
}


func printList(head *Node) {
    curr := head
    for curr != nil {
        fmt.Print(curr.data, " -> ")
        curr = curr.next
    }
    fmt.Println("nil")
}

func main(){
	var head * Node
	head = appendNode(head, 10)
	head = appendNode(head, 20)
	head = appendNode(head, 30)
	head = appendNode(head, 40)
	printList(head)
}
*/
