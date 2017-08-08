package main

import (
	"fmt"
	"math/rand"
	"github.com/billyninja/dstructures/linked_list"
)

type Vector4 struct {
	X	float32
	Y	float32
	Z   float32
	W   float32
}

type Point struct {
	Pos 	Vector4
	Color 	Vector4
}

func (p *Point) Print() {
	fmt.Printf("Pos - X: %.2f, Y: %.2f, Z: %.2f, W: %.2f", p.Pos.X, p.Pos.Y, p.Pos.Z, p.Pos.W)
	fmt.Printf("Color - R: %.2f, G: %.2f, B: %.2f, A: %.2f", p.Color.X, p.Color.Y, p.Color.Z, p.Color.W)
} 

func rnd4d() Vector4{
	return Vector4{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}
}

func rndPoint() *Point{
	return &Point{
		Pos: rnd4d(),
		Color: rnd4d(),
	}
}

func main() {
	llw := llist.NewLinkedList(rndPoint(), rndPoint(), rndPoint())

	println(llw.Count, llw.First, llw.Last)
	llw.Append(rndPoint())
	println(llw.Count, llw.First, llw.Last)
	llw.Preppend(rndPoint())
	println(llw.Count, llw.First, llw.Last)
	llw.Pop()
	println(llw.Count, llw.First, llw.Last)
	llw.Pop()
	llw.Pop()
	llw.Pop()
	llw.Pop()
	llw.Pop()
	llw.Pop()
	println(llw.Count, llw.First, llw.Last)
}
