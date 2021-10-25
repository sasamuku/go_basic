package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
	"github.com/sasamuku/go_basic/mylib"
)

func main() {
	// Packages
	s := []int{1, 2, 3, 4}
	fmt.Println(mylib.Average(s))

	// Public と Private
	mylib.Say() // Public: 大文字で始まる関数は外部パッケージからも呼び出せる
	// mylib.say() Private: 小文字で始まる関数は外部パッケージから呼び出せない

	// サードパーティのパッケージ
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

}
