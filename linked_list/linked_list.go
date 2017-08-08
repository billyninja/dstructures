package llist

import (
	"fmt"
	"errors"
)

type Sorting func(a, b Entity) Entity

type Wrapper struct {
	First 		*Node
	Last 		*Node
	Count 		int64
	Sorting  	Sorting	 
}

type Node struct {
	Prev 	*Node
	El 		Entity
	Next 	*Node
}


func NewNode(el Entity, prev, next *Node) *Node {
	this := &Node{
		Prev: prev,
		Next: next,
		El: el,
	}
	if prev != nil {
		prev.Next = this
	}
	if next != nil {
		next.Prev = this
	}

	return this
}


func NewLinkedList(initials...Entity) *Wrapper {
	var first *Node
	var ctn int64

	if len(initials) > 0 {
		first = NewNode(initials[0], nil, nil)
		ctn += 1
	}

	llw := &Wrapper{
		First: first,
		Last: first,
		Count: ctn,
	}

	// If there's more than 1 initial
	// lets insert it the cannonical way
	for i := 1; i < len(initials); i++ {
		llw.Append(initials[i])
	}

	return llw
}

func (llw *Wrapper) Append(el Entity) {
	new := NewNode(el, nil, nil)
	if llw.Last == nil {
		llw.First = new
		llw.Last = new
	} else {
		pLast := llw.Last
		pLast.Next = new
		new.Prev = pLast
		llw.Last = new
		new.Next = nil
	}
	llw.Count += 1
}

func (llw *Wrapper) Preppend(el Entity) {
	new := NewNode(el, nil, nil)
	if llw.First == nil {
		llw.First = new
		llw.Last = new
	} else {
		pFirst := llw.First
		pFirst.Prev = new
		new.Next = pFirst
		llw.First = new
	}
	llw.Count += 1
}

func (llw *Wrapper) Pop() Entity {
	if llw.Last == nil{
		return nil
	}

	pLast := llw.Last
	llw.Last = pLast.Prev
	llw.Count -= 1
	return pLast.El
}

func (llw *Wrapper) InsertAt(el Entity, idx int64) (*Node, error) {
	// the new element should occupy the idx position at the LList
	if idx >= llw.Count {
		msg := fmt.Sprintf("The requested index %d is out bound", idx)
		return nil, errors.New(msg)
	}

	nxt := llw.First
	prv := llw.First

	for i := 0; i < int(idx); i++ {
		nxt = nxt.Next
		prv = nxt.Prev
	}
	new := NewNode(el, prv, nxt)

	return new, nil
}


func (llw *Wrapper) swap(a, b *Node) {
	pa := a.Prev
	if pa != nil {
		pa.Next = b
	}

	nb := b.Next
	if nb != nil {
		nb.Prev = a	
	}

	if a == llw.First {
		llw.First = b
	}

	if b == llw.Last {
		llw.Last = b
	}

	a.Next = b.Next
	b.Next = a
	b.Prev = a.Prev
	a.Prev = b

}


func (llw *Wrapper) BubbleSort() {
	// [*, 3, *] [*, 1, *] [*, 2, *]
	var runs uint8

	for {
		curr := llw.First
		next := curr.Next
		swapped := false
		for {
			if next == nil {
				curr = llw.First
				break
			}

			if curr.El.Weight() < next.El.Weight() {
				fmt.Printf("Indeed, %.2f > %.2f, let's swap\n", curr.El.Weight(), next.El.Weight())
				llw.swap(curr, next)
				swapped = true
			}

			if curr.Next == nil {
				curr = llw.First
				break
			}

			curr = curr.Next
			next = curr.Next			
		}

		if !swapped {
			fmt.Printf("\nHaven't swapped this run (%d). Done BubbleSort!!!\n", runs)
			return
		}
		runs += 1
		fmt.Printf("Done pass #%d in BubbleSort", runs)
	}
}

func (llw *Wrapper) Present() {
	curr := llw.First
	next := curr.Next

	for { 
		curr = next 
		fmt.Printf("W: %.2f - ", curr.El.Weight())
		next = curr.Next
		if next == nil {
			return
		}
	}
}

type Entity interface {
	Print()
	Weight() float64
}
