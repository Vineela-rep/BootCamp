package main

import (
	"fmt"
	"sync"
)

var m = make(map[string]int,26)
func Frequency(word string,wg *sync.WaitGroup)  {
	for i:=0;i<len(word);i++ {
		m[string(word[i])]++
	}
	wg.Done()
}
func main()  {
	var wg sync.WaitGroup
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	for i:= range words{
		wg.Add(1)
		go Frequency(words[i],&wg)
	}
	wg.Wait()
	fmt.Println(m)
}
