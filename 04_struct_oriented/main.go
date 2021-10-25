package main

import "fmt"

type Vertex struct {
	X, Y int
}

func Area(v Vertex) int { // 普通の関数
	return v.X * v.Y
}

func (v Vertex) Area() int { // メソッド (値レシーバー)
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) { // メソッド (ポインタレシーバー)
	v.X = v.X * i
	v.Y = v.Y * i
}

func New(x, y int) *Vertex { // コンストラクタ
	return &Vertex{x, y}
}

type Vertex3D struct { // Embedded (継承みたいな)
	Vertex
	Z int
}

func (v Vertex3D) Area3D() int { // メソッド (値レシーバー)
	return v.X * v.Y * v.Z
}

func (v *Vertex3D) Scale3D(i int) { // メソッド (ポインタレシーバー)
	v.X = v.X * i
	v.Y = v.Y * i
	v.Z = v.Z * i
}

func New3D(x, y, z int) *Vertex3D { // コンストラクタ
	return &Vertex3D{Vertex{x, y}, z}
}

// non-struct
type MyInt int

// non-struct のメソッド
func (i MyInt) Double() int {
	return int(i * 2)
}

// interface
type Human interface {
	Say() string
}

type Japanese struct {
	Name string
}

type Dog struct {
	Name string
}

func (p Japanese) Say() string {
	return "konitiha" + p.Name
}

func (d Dog) Wan() string {
	return "wanwan" + d.Name
}

// switch type 文
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println("Hello" + v)
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

// stringer
func (j Japanese) String() string {
	return fmt.Sprintf("Watashi no namae ha" + j.Name)
}

// カスタムエラー
type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string { // UserNotFount という struct に対してカスタムエラーを実装
	return fmt.Sprintf("User not found: %v\n", e.Username)
}

func FuncError() error {
	return &UserNotFound{Username: "mike"}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
	fmt.Println(v.Area())
	v.Scale(10)
	fmt.Println(v.Area())

	// コンストラクタ
	vv := New(5, 6)
	fmt.Println(vv.Area())

	// Embedded
	v3d := New3D(2, 3, 4)
	v3d.Scale3D(10)
	fmt.Println(v3d.Area3D())

	// non-struct のメソッド
	myInt := MyInt(5)
	fmt.Println(myInt.Double())

	// interface
	var h Human = Japanese{"Taro"}
	fmt.Println(h.Say())

	// Error:DogにはSayメソッドがないため宣言できない
	// var dog Human = Dog{"Poti"}
	// fmt.Println(dog)

	// switch type 文
	do(10)
	do("Mike")
	do(true)

	// stringer
	var hu Human = Japanese{"Jiro"}
	fmt.Println(hu)

	if err := FuncError(); err != nil {
		fmt.Println(err)
	}
}
