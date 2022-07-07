package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var r = make (chan int)
func student(id int) {
	n := rand.Intn(500)
	rating:=rand.Intn(10)
	r <-rating
	fmt.Printf("Student %d rating %d\n", id, rating)
	time.Sleep(time.Duration(n)*time.Millisecond)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	avg:=0
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go func (){
			student(i)
			defer wg.Done()
		}()
		//wg.Done()
		avg+= <-r
	}
	wg.Wait()
	fmt.Println("average:", float64(avg)/200.0)


}