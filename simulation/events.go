package simulation

import "container/heap"
import "encoding/json"
import "log"
import "io/ioutil"
import "bytes"

func (block BlockArrival) Visit(chain *Blockchain) {
	// Adjust the lambda value
	chain.TxLambda = chain.TransactionArrivalLambda()

	// Write out the new block to the blockchain
	chain.History = append(chain.History, chain.CurrentBlock)
	chain.CurrentBlock = Block{
		block.Time,
		make([]Transaction, 0),
	}

	// Generate a new block arrival event
	next := BlockArrival{Time(int(block.Time) + chain.BlockArrival())}
	heap.Push(chain, next)

	// check to see if there are queued up transactions
	// and pull them off the queue and into the next block
	if size := chain.TransactionWaitQueue.Size(); size != 0 {
		for ; size != 0 && !chain.BlockIsFull(); size = chain.TransactionWaitQueue.Size() {
			next, _ := chain.TransactionWaitQueue.Pop()
			nextTx := next.(Transaction)
			chain.CurrentBlock.Transactions = append(chain.CurrentBlock.Transactions, nextTx)
		}
	}
}

func (chain *Blockchain) AddTx(tx Transaction) {
	// check to see if the block is full. If so, add this to the transaction queue and return
	if chain.BlockIsFull() {
		chain.TransactionWaitQueue.Push(tx, tx.Fee)
		return
	}

	// Add a new transaction to the tx list
	chain.CurrentBlock.Transactions = append(chain.CurrentBlock.Transactions, tx)
}

func (tArrival TransactionArrival) Visit(chain *Blockchain) {

	// Make the new transaction
	tx := Transaction{
		ArrivalTime: tArrival.Time,
		Fee:         chain.Fee(),
	}

	// Add the transaction to the blockchain
	chain.AddTx(tx)

	// Generate a new transaction arrival event
	nextTime := int(tArrival.Time)
	nextTime += chain.TxArrival(chain.TxLambda)

	next := TransactionArrival{Time(nextTime)}
	heap.Push(chain, next)
}

// Print out all of the results from the experiment
func (e Exit) Visit(chain *Blockchain) {

	var buffer bytes.Buffer
	// Collect all of the data as JSON
	// Then write it to the buffer

	payload := struct {
		MaxTransactions int     `json:"max_transactions_per_block"`
		Blocks          []Block `json:"blocks"`
	}{
		chain.TransactionMax,
		chain.History,
	}

	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(payload)
	if err != nil {
		log.Println(err)
	}

	// Write the buffer to the file
	err = ioutil.WriteFile(e.filename, buffer.Bytes(), 0655)
	if err != nil {
		log.Println(err)
	}
}
