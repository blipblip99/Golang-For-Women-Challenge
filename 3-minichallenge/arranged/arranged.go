package main

import (
	"fmt"
	"sync"
)

type Displayable interface {
	Display()
}


type listItem struct {
	str1  string
	str2  string
	
}

func (x listItem) Display() {
	for a := 1; a < 5; a++ {
		fmt.Println(x.str1, a)
		fmt.Println(x.str2, a)
	}
}


func displayList(items []Displayable) {
	var wg sync.WaitGroup
	var mu sync.Mutex 
	

	wg.Add(len(items))

	
		for _, item := range items {
			go func(i Displayable) {
				mu.Lock() 
				i.Display()
				mu.Unlock() 
				wg.Done()
			}(item)
		}
	

	wg.Wait()
}

func main() {
		
		items := []Displayable{
			
			listItem{str1: "[coba1 coba2 coba3] ", str2:"[bisa1 bisa2 bisa3] "},
		
	}
	
	displayList(items)
}
