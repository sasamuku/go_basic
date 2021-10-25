package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine と sync.WaitGroup
func normal(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func normal2(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

// channel
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// Producer and Consumer
func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("ch2 is", i*1000)
		wg.Done()
	}
}

// fan-out と fan-in
func first_fn(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func second_fn(first <-chan int, second chan<- int) { // 矢印で方向を明示的に表現可能
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func third_fn(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 2
	}
}

// select
func goroutine1(ch chan string) {
	for {
		ch <- "goroutine 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "goroutine 2"
		time.Sleep(3 * time.Second)
	}
}

// sync.Mutex (同時書き込みエラーを防げるやつ)
type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// goroutine と sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go normal2("Goroutine", &wg) // goroutine
	normal("Non Gorutine")
	wg.Wait() // wg.Done() が実行されるまで終了を待つ
	fmt.Println("wg.Done() is implemented")

	// channel
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // チャネルは goroutine とやり取りするためのキューとも解釈できる
	go sum(s, c)
	x := <-c
	fmt.Println(x)

	// Buffered channels
	ch := make(chan int, 2)
	ch <- 100
	ch <- 200
	close(ch) // channelを閉じないとFor文が無限ループする

	for ch := range ch {
		fmt.Println(ch)
	}

	// Producer と Consumer
	var wg2 sync.WaitGroup
	ch2 := make(chan int)
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go producer(ch2, i)
	}
	go consumer(ch2, &wg2)
	wg2.Wait() // wg.Done()を待たないとゴルーチンが結果を出力する前に終了してしまう
	close(ch2) // chを閉じないとconsumerのfor文が終了しない

	// fan-out と fan-in
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)
	go first_fn(first)
	go second_fn(first, second)
	go third_fn(second, third)
	for i := range third {
		fmt.Println(i)
	}

	// select
	ch_one := make(chan string)
	ch_two := make(chan string)
	go goroutine1(ch_one)
	go goroutine2(ch_two)
	// for {
	// 	select {
	// 	case msg1 := <-ch_one: // チャネル待ちのブロックが発生しない
	// 		fmt.Println(msg1)
	// 	case msg2 := <-ch_two:
	// 		fmt.Println(msg2)
	// 	}
	// }

	// default selection
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	// sync.Mutex (同時書き込みエラーを防げるやつ)
	counter := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			counter.Inc("Key") // lockをかけるので同じstructにアクセスしても大丈夫
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			counter.Inc("Key") // lockをかけるので同じstructにアクセスしても大丈夫
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(counter.v["Key"])
}
