package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func hello(i int) {
	defer wg.Done()
	fmt.Println("hello", i)
}

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
