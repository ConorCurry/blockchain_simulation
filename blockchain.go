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

	blockchain05 := NewBlockchain(t)
	blockchain05.TransactionMax = blockchain2.TransactionMax >> 1

	log.Println("Blockchains successfully created\nNow to run the simuations...")

	var wg sync.WaitGroup
	wg.Add(3)
	go blockchain.Run(&wg, "data/output1.json")
	go blockchain2.Run(&wg, "data/output2.json")
	go blockchain05.Run(&wg, "data/output05.json")
	wg.Wait()
	log.Println("Simulation complete.")
}
