package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

const Pi = 3.14

func main() {
	// Hello World
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())

	// 変数
	i := 1 // 省略記法 <変数名> := <値>
	f := 1.2
	var s string = "test" // 通常記法 var <変数名> <型> = <値>
	fmt.Println(i, f, s)
	fmt.Printf("%T\n", Pi)
	fmt.Println(1+1, 2+2)

	// string
	st := "Hello world"
	st = strings.Replace(st, "H", "S", 1)
	fmt.Println(st)
	fmt.Println(`test + 
	test`)
	fmt.Println(`"""`)

	// 型変換
	var aa int = 1
	b := float64(aa)
	fmt.Printf("%T %v %f\n", b, b, b)

	var str string = "14"
	bb, _ := strconv.Atoi(str) // err は捨てる
	fmt.Printf("%T %v %d\n", bb, bb, bb)

	// 配列
	var list_1 = [2]int{}
	list_1[0] = 100
	list_1[1] = 200
	fmt.Println(list_1[0], list_1[1])

	var list_2 [3]int = [3]int{100, 200, 300}
	fmt.Println(list_2[0], list_2[1], list_2[2])

	// スライス
	// 配列との宣言の仕方の違いは[]内に要素数を入れるかどうか
	n := []int{10, 20, 30, 40}
	fmt.Println(n)
	n = append(n, 100, 200, 300)
	fmt.Println(n)

	// スライスの初期化
	sli := make([]int, 3, 5) // make([]<型>, len, cap)
	fmt.Printf("len=%d cap=%d value=%v\n", len(sli), cap(sli), sli)

	// スライスの拡張によるメモリ再確保
	slia := make([]int, 3) // capを省略するとlenと同値となる
	fmt.Printf("slia len=%d cap=%d value=%v\n", len(slia), cap(slia), slia)
	slia = append(slia, 1, 2)
	fmt.Printf("slia len=%d cap=%d value=%v\n", len(slia), cap(slia), slia)
	// capは確保しているメモリの長さ、lenは実際に値が入っている分の長さ
	// 分からなくなったら読む -> https://qiita.com/Kashiwara/items/e621a4ad8ec00974f025

	c := make([]int, 5) // [0 0 0 0 0]が初期値
	// c := make([]int, 0, 5) // []が初期値
	c = append(c, 100)
	fmt.Println(c)

	// map (辞書型)
	m := map[string]int{"apple": 100, "banana": 500}
	fmt.Println(m)

	// map の初期化
	m2 := make(map[string]int) // 空のmapを用意（メモリ確保）
	m2["mikan"] = 70
	fmt.Println(m2)

	// これはエラー
	// var m3 map[string]int // mapを宣言（メモリ確保しない）
	// m3["budou"] = 70
	// fmt.Println(m3)

	// バイト型
	byte := []byte{72, 73}
	fmt.Println(byte)
	fmt.Println(string(byte))

	// 関数
	add_result := add(10, 20)
	fmt.Println(add_result)

	result_add, result_sub := calc(100, 200)
	fmt.Println(result_add, result_sub)

	// inner 関数
	test_func := func(x int) {
		fmt.Println("Num is", x)
	}
	test_func(1)

	func(x int) {
		fmt.Println("Num is", x)
	}(1) // 宣言と同時に呼べる

	// クロージャー
	circleArea1 := circleArea(3.14) // 円周率が3.14の世界の面積を求める関数
	fmt.Println(circleArea1(2))
	circleArea2 := circleArea(6.28) // 円周率が6.28の世界の面積を求める関数
	fmt.Println(circleArea2(2))

	// 可変長引数
	variadic_func(10, 20, 30)
	// params := [3]int{40, 50, 60} // 配列を渡すとエラーとなる
	params := []int{40, 50, 60}
	variadic_func(params...)
}

// 関数
func add(a, b int) int {
	return a + b
}

func calc(a, b int) (add, sub int) {
	add = a + b
	sub = a - b
	return
	// return add, sub としてもOK
}

// クロージャー
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

// variadic function
func variadic_func(params ...int) {
	fmt.Printf("Type is %T, len is %d, value is %v\n", params, len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}
