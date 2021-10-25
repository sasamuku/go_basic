package main

import (
	"fmt"
	"time"
)

func main() {
	// if
	x := 4
	if x%2 == 0 {
		fmt.Println("gusu")
	} else if x%2 == 1 {
		fmt.Println("kisu")
	} else {
		fmt.Print("not integer")
	}

	if x := 2; x%2 == 0 { // if文の中だけで扱う変数(if文の外では利用できない)
		fmt.Println("gusu")
	}
	fmt.Println(x) // if文での宣言時とは異なる値となる

	// for
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("Continue")
			continue // 以降の処理を飛ばしてループに戻る
		}
		if i > 6 {
			fmt.Println("break")
			break // 以降の処理を飛ばしてループを抜ける
		}
		fmt.Println(i)
	}

	// 無限ループもfor文で書ける
	for {
		break
	}

	// range
	l := []string{"banana", "apple", "orange"}
	for i, v := range l {
		fmt.Println(i, v)
	}

	// switch
	os_name := "mac"
	switch os_name {
	case "mac":
		fmt.Println("mac!!!")
	case "windows":
		fmt.Println("win!!!")
	default:
		fmt.Println("default!!!") // どれにも該当しないとき
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Evening")
	}

	// defer
	// 動きとしてはスタックでLIFOとなる
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	// log
	// log.Println("logging")
	// log.Fatalln("error")

	// エラーハンドリング
	// file, err := os.Open("./lesson.go")
	// if err != nil {
	// 	log.Fatalln("error!")
	// }
	// defer file.Close()
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatalln("error!!")
	// }
	// fmt.Println(count, string(data))

	// panic と recover
	save()
	fmt.Println("recovered!")
}

func ExternalFunc() {
	panic("Unknown error!")
}

func save() {
	// ExternalFunc() // deferの前で呼ぶとrecoverできない
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	ExternalFunc()
}
