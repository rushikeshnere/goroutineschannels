package main

import (
	"fmt"
	"sync"
)


var evenChannel, oddChannel chan int
var num int = 1
var wg sync.WaitGroup

	
func main() {
	evenChannel = make(chan int)
	oddChannel = make(chan int)

	wg.Add(1)	
	go oddNumPrint()	
	wg.Add(1)	
	go evenNumPrint()
	
	wg.Wait()
}

func oddNumPrint() {
	defer wg.Done()
	for  {
		fmt.Println(num)
		num ++
		evenChannel <- num
		num = <- oddChannel
	}
}

func evenNumPrint() {
	defer wg.Done()
	for {
		num = <- evenChannel
		fmt.Println(num)
		num ++
		oddChannel <- num
		
	}
}
