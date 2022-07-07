package main

import (
	"fmt"
	"sync"
)

var freq = make(chan map[string]int, 5)

//var m map[string]int
func Frequency(word string)  {
	m:=make(map[string]int,26)
	for i:=0;i<len(word);i++ {
		m[string(word[i])]+= 1
	}
	freq <-m
}
func main()  {
	var wg sync.WaitGroup
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	for i:= range words{
		wg.Add(1)
		go Frequency(words[i])
		wg.Done()
	}
	wg.Wait()
	for i:=1;i<=5;i++{
		fmt.Println(<-freq)
	}
}
