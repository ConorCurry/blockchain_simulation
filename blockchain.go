package main

import (
	. "github.com/natboehm/blockchain_simulation/simulation"
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("Beginning simulation.")

	t := time.Now().UnixNano()
	blockchain := NewBlockchain(t)

	blockchain2 := NewBlockchain(t)
	blockchain2.TransactionMax = blockchain2.TransactionMax << 1

	blockchain4 := NewBlockchain(t)
	blockchain4.TransactionMax = blockchain4.TransactionMax << 2

	blockchain8 := NewBlockchain(t)
	blockchain8.TransactionMax = blockchain8.TransactionMax << 3

	blockchain16 := NewBlockchain(t)
	blockchain16.TransactionMax = blockchain16.TransactionMax << 4

	var wg sync.WaitGroup
	wg.Add(5)
	go blockchain.Run(&wg, "data/output1.json")
	go blockchain2.Run(&wg, "data/output2.json")
	go blockchain4.Run(&wg, "data/output4.json")
	go blockchain8.Run(&wg, "data/output8.json")
	go blockchain16.Run(&wg, "data/output16.json")
	wg.Wait()
}
