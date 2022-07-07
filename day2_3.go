package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mutex sync.Mutex
var balance int = 500

func deposit(val int, wg *sync.WaitGroup)  {
	mutex.Lock()
	balance+=val
	fmt.Printf("Credited Rs. %d\t Current Balance = Rs. %d\n",val,balance)
	mutex.Unlock()
	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup){
	mutex.Lock()
	if balance<val{
		fmt.Printf("Insufficient Balance for debiting Rs. %d\t Current Balance = Rs. %d\n",val,balance)
	}else{
		balance-=val
		fmt.Printf("Debited Rs. %d\t Current Balance = Rs. %d\n",val,balance)
	}
	mutex.Unlock()
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	for i:=0;i<3;i++{
		wg.Add(2)
		go withdraw(rand.Intn(1500),&wg)
		go deposit(rand.Intn(500),&wg)
		wg.Wait()
	}
	fmt.Println("Account Balance: ",balance)
}