package spec

import (
	"container/heap"
	"fmt"
	. "github.com/natboehm/blockchain_simulation/simulation"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {

	s := time.Now().UnixNano()
	blockchain := NewBlockchain(s)

	t1 := TransactionArrival{30}
	t2 := TransactionArrival{10}

	heap.Push(blockchain, t1)
	heap.Push(blockchain, t2)

	t3 := heap.Pop(blockchain).(TransactionArrival)
	t4 := heap.Pop(blockchain).(TransactionArrival)

	if t3 != t2 {
		t.Errorf("T3 : %v\n T2 : %v", t3, t2)
	}

	if t1 != t4 {
		t.Errorf("T1 : %v\n T4 : %v", t1, t4)
	}
}

func TestNewBlockchain(t *testing.T) {
	s := time.Now().UnixNano()
	blockchain := NewBlockchain(s)

	if blockchain.BlockArrival() < 0 {
		t.Errorf("Expected number > 0")
	}
	if blockchain.TransactionArrivalLambda() < 0 {
		t.Errorf("Expected number > 0")
	}
	if blockchain.Fee() < 0 {
		t.Errorf("Expected number > 0")
	}
}

func TestEventHeapSize(t *testing.T) {
	s := time.Now().UnixNano()
	blockchain := NewBlockchain(s)
	if blockchain.Len() != 0 {
		t.Errorf("Incorrect size")
	}

	t1 := TransactionArrival{30}
	heap.Push(blockchain, t1)
	if blockchain.Len() != 1 {
		t.Errorf("Incorrect size")
	}
}

func TestBlocksAddedToHeap(t *testing.T) {
	s := time.Now().UnixNano()
	blockchain := NewBlockchain(s)

	b := BlockArrival{Time(5)}
	heap.Push(blockchain, b)

	event := heap.Pop(blockchain).(Event)
	event.Visit(blockchain)
	if blockchain.Len() != 1 {
		t.Errorf("Incorrect heap size: %v", blockchain.Len())
	}

	if len(blockchain.History) != 1 {
		t.Errorf("Incorrect history size: %v", len(blockchain.History))
	}

	if event := heap.Pop(blockchain).(Event); event.When() < 0 {
		t.Errorf("Incorrect time: %v", event.When())
	}
}

func TestTransactionsAddedToBlock(t *testing.T) {
	s := time.Now().UnixNano()
	blockchain := NewBlockchain(s)

	trans := TransactionArrival{Time(5)}
	heap.Push(blockchain, trans)

	event := heap.Pop(blockchain).(Event)
	event.Visit(blockchain)
	if blockchain.Len() != 1 {
		t.Errorf("Incorrect heap size: %v", blockchain.Len())
	}

	if len(blockchain.History) != 0 {
		t.Errorf("Incorrect history size: %v", len(blockchain.History))
	}

	if len(blockchain.CurrentBlock.Transactions) != 1 {
		t.Errorf("Incorrect transaction size: %v", len(blockchain.CurrentBlock.Transactions))
	}

	if event := heap.Pop(blockchain).(Event); event.When() < 0 {
		t.Errorf("Incorrect time: %v", event.When())
	}
}

// Tests to see if transactions are removed from the Tx queue correctly
func TestOffloadTransactionsFromQ(t *testing.T) {

	s := time.Now().UnixNano()
	chain := NewBlockchain(s)

	// First, add 2 tx to the queue
	tx := Transaction{
		ArrivalTime: 1,
		Fee:         chain.Fee(),
	}

	tx2 := Transaction{
		ArrivalTime: 1,
		Fee:         chain.Fee(),
	}

	// Add those tx to the queue
	chain.TransactionWaitQueue.Push(tx, tx.Fee)
	chain.TransactionWaitQueue.Push(tx2, tx.Fee)

	// Process an block arrival event
	blkArrival := BlockArrival{Time(1)}
	blkArrival.Visit(chain)

	// Assert that there are now 2 tx on the current block
	if len(chain.CurrentBlock.Transactions) != 2 {
		t.Errorf("Incorrect number of transactions: %v", chain.CurrentBlock.Transactions)
	}
}

// Tests whether or not a full block will correctly pull transactions from the queue
func TestFullBlockOrdersCorrectly(t *testing.T) {
	data := []Transaction{
		{
			ArrivalTime: 1,
			Fee:         10,
		}, {
			ArrivalTime: 1,
			Fee:         11,
		}, {
			ArrivalTime: 1,
			Fee:         12,
		},
	}

	s := time.Now().UnixNano()
	chain := NewBlockchain(s)
	chain.TransactionMax = 1

	// Add the 3 transactions to the blockchain
	for _, tx := range data {
		chain.AddTx(tx)
	}

	if tx := chain.CurrentBlock.Transactions[0]; tx.Fee != 10 {
		t.Error("Did not find the correct fee in the current block's transactions")
	}

	if chain.TransactionWaitQueue.Size() != 2 {
		t.Errorf("Did not have the correct number of tx waiting in the queue")
	}

	// then fire a block arrival event
	blkArrival := BlockArrival{Time(5)}
	blkArrival.Visit(chain)

	if len(chain.History) != 1 {
		t.Errorf("History not correct len: %v", len(chain.History))
	}

	// make sure that the correct tx is added to the new block
	if fee := chain.CurrentBlock.Transactions[0].Fee; fee != 12 {
		t.Errorf("Incorrect fee on the tx in the current block: %v", fee)
	}

	fmt.Println(chain.TransactionWaitQueue.Size())
	fmt.Println(chain.History)
	fmt.Println(chain.CurrentBlock)

	/*
		// and that the cheaper block is still in the tx queue
		if tx, _ := chain.TransactionWaitQueue.Head(); tx.(Transaction).Fee != 11 {
			t.Errorf("Incorrect fee on the tx in the queue: %v", tx.(Transaction).Fee)
		}
	*/
}
