package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
	Z    string
	x1   float64
}

func main() {
	// ポインタ
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	// ポインタ型を宣言するときはnew()
	var po *int = new(int)
	fmt.Printf("%T\n", po)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	// ポインタ型でないときはmake()
	li := make([]int, 0)
	fmt.Printf("%T\n", li)

	// struct
	v := Vertex{X: 1, Y: 2, x1: 0.1} // Zを宣言しなければ構造体のフィールドに値が追加されない
	// v := Vertex{X: 1, Y: 2, x1: 0.1, Z: "test"}
	fmt.Println(v)
	v.x1 = 0.55
	fmt.Println(v.X, v.Y, v.x1)
}
