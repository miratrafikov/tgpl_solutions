package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func m1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func m2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func m3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	funcs := [3]func(){m1, m2, m3}
	var results [3]int64
	for i, v := range funcs {
		fmt.Println("Func #", i)
		start := time.Now()
		v()
		results[i] = time.Since(start).Nanoseconds()
	}
	for i, v := range results {
		fmt.Printf("Func #%d: %d ns\n", i, v)
	}
}
