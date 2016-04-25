package spec

import (
	"container/heap"
	. "github.com/natboehm/blockchain_simulation/simulation"
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {

	s := rand.NewSource(time.Now().UnixNano())
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
