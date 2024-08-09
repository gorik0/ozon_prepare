package main

import (
	"log"
	"math/rand"
)

/**
 * Definition for singly-linked list.
 */
func main() {
	//var nod = &ListNode{0, nil}
	//var size = 5
	l1 := BuildListNodeWithListAndOffset([]int{3, 2, 4}, 0)
	l2 := BuildListNodeWithListAndOffset([]int{4, 6, 5}, 0)
	log.Println(l1)
	log.Println(l2)
	//nod2 := &ListNode{0, nil}
	a := make([]int, 0)
	no := addTwoNumbers(l1, l2)
	ThroughtLinkedList(no, &a)
	log.Println(a)
}

func CreateListNode(nod *ListNode, size int) *ListNode {
	nod.Val = rand.Intn(10)
	size--
	if size == 0 {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	newNode := new(ListNode)
	nod.Next = CreateListNode(newNode, size)
	return nod
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var accum1 []int
	var accum2 []int
	ThroughtLinkedList(l1, &accum1)
	ThroughtLinkedList(l2, &accum2)

	//log.Println(accum1)
	//log.Println(accum2)

	no := (BuildListNodeWithListAndOffset(SumTwoList(accum1, accum2)))
	return no

}

func BuildListNodeWithListAndOffset(list []int, offset int) *ListNode {
	var nod *ListNode
	if offset != 0 {
		nod = &ListNode{
			Val:  offset,
			Next: ListNodeFromList(list, len(list)-1),
		}
	} else {
		nod = ListNodeFromList(list, len(list)-1)
	}
	return nod
}

func ListNodeFromList(list []int, i int) *ListNode {
	//println("i :::", i)
	//log.Println(list)
	if i == 0 {
		return &ListNode{
			Val:  list[0],
			Next: nil,
		}
	}
	return &ListNode{Val: list[i], Next: ListNodeFromList(list, i-1)}

}

func SumTwoList(accum1 []int, accum2 []int) ([]int, int) {
	var maxLen int
	var offset int
	var diffLen int
	var tail int
	var firstListMore bool
	var sum []int
	maxLen, firstListMore, diffLen = findMaxLen(accum1, accum2, maxLen)
	sum = make([]int, maxLen)
	if firstListMore {
		for i := maxLen - 1; i >= 0; i-- {
			if i < diffLen+1 {
				sum[i] = accum1[i]
			} else if tail = accum1[i] + accum2[i-diffLen] + offset; tail < 10 {
				sum[i] = tail
				offset = 0

			} else {
				offset = 1
				sum[i] = tail - 10
			}
		}
	} else {
		for i := maxLen - 1; i >= 0; i-- {
			if i < diffLen+1 {
				sum[i] = accum2[i]
			} else if tail = accum1[i-diffLen] + accum2[i] + offset; tail < 10 {
				sum[i] = tail
				offset = 0

			} else {
				offset = 1
				sum[i] = tail - 10
			}
		}
	}
	return sum, offset
}

func findMaxLen(accum1 []int, accum2 []int, maxLen int) (int, bool, int) {
	var f bool
	var dif int
	if len(accum1) > len(accum2) {
		dif = len(accum1) - len(accum2)
		maxLen = len(accum1)
		f = true
	} else {
		dif = len(accum2) - len(accum1)

		maxLen = len(accum2)
		f = false
	}
	return maxLen, f, dif
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ThroughtLinkedList(nod *ListNode, accum *[]int) {

	if nod.Next == nil {
		*accum = append(*accum, nod.Val)

		return
	} else {
		*accum = append(*accum, nod.Val)
		ThroughtLinkedList(nod.Next, accum)
	}

}
