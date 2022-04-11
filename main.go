package main 

import "fmt"

type Node struct {
        value string
        next  *Node
        prev  *Node 
}

type DLL struct {
        head *Node
        tail *Node
}

func (dll *DLL) add_node_end(node *Node) *Node {
        if dll.head == nil {
                dll.head = node 
                dll.tail = node
                node.next = nil 
                node.prev = nil
                return node
        }

        aux := dll.tail
        aux.prev = node 
        node.next = aux 
        node.prev = nil
        dll.tail = node

        return node
}

func (dll *DLL) drop_node_end() {
        aux := dll.tail 
        if aux.next != nil {
               dll.tail = aux.next
               dll.tail.prev = nil
        } else {
                dll.tail = nil
                dll.head = nil
        }


}

func (dll *DLL) replace_last_node (node *Node) (*Node) {
        dll.drop_node_end()
        ret := dll.add_node_end(node)

        return ret

}

func (dll *DLL) move_front(node *Node) {
        if dll.head == node {
                return 
        }
        if node.next != nil {
               node.next.prev = node.prev
        }
        if node.prev != nil {
               node.prev.next = node.next
        }
        aux := dll.head
        node.next = nil
        node.prev = aux
        aux.next = node
        dll.head = node
}

type LRU struct {
        lru_map map[int32]*Node
        ordered *DLL
        size int32
        capacity int32
}

func NewLRU(capacity int32) (*LRU) {
        dll := DLL { head: nil, tail: nil }
        lru := LRU { lru_map: make(map[int32]*Node), ordered: &dll, size: 0, capacity: capacity}

        return &lru
}

func (lru *LRU) set(key int32, value string) {

        v, found := lru.lru_map[key]
        if found {
                lru.lru_map[key] = v
                lru.ordered.move_front(v)
                return 
        }
        node := Node { value: value, next: nil, prev: nil }
        var n *Node = nil
        if lru.size >= lru.capacity {
                fmt.Printf("Adding [%d]%s -> %d\n", key, value, lru.size)
                n = lru.ordered.replace_last_node(&node)
        } else {
                fmt.Printf("Adding [%d]%s -> %d\n", key, value, lru.size)
                n = lru.ordered.add_node_end(&node)
                lru.size += 1
        }
        lru.lru_map[key] = n
}

func (lru *LRU) get(key int32) (*Node, error) {

        v, found := lru.lru_map[key]
        if !found {
                return nil, fmt.Errorf("cannot find key in map")
        }

        lru.ordered.move_front(v)

        return v, nil
}

func main () {
        lru := NewLRU(10)

        lru.set(1, "1")
        lru.set(2, "2")
        lru.set(3, "3")
        lru.set(4, "4")
        lru.set(5, "5")
        lru.set(6, "6")
        lru.set(7, "7")
        lru.set(8, "8")
        lru.set(9, "9")
        lru.set(10, "10")
        lru.set(11, "11")

 //      lru.get(6)
 //      lru.get(3)
 //      lru.get(1)
 //      lru.get(7)
 //      lru.get(7)
        lru.get(11)

        node := lru.ordered.head
        for true {
                if node.prev != nil {
                        if node.next == nil {
                                fmt.Printf("Got node: %s<-[%s]\n", node.prev.value, node.value)
                        } else {
                                fmt.Printf("Got node: %s<-[%s]->%s\n", node.prev.value, node.value, node.next.value)
                        }

                        node = node.prev
                } else {
                        fmt.Printf("Got node:    [%s]->%s\n", node.value, node.next.value)
                        break
                }
        }
}
