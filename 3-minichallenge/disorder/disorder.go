package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup 

func routine(i int) {
	defer wg.Done()
	fmt.Println("[coba1 coba2 coba3] ", i)
	fmt.Println("[bisa1 bisa2 bisa3] ", i)
}

func main() {
	for i := 1; i < 5; i++ {
		wg.Add(1)     
		go routine(i) 
	}
	wg.Wait() 

}
