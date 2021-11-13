package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		// 5head code
		// fmt.Println(strings.Join([]string{strconv.Itoa(i), v}, " "))
		// 4head code
		fmt.Printf("Argument %d: %s\n", i, v)
	}
}
