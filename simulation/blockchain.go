package simulation

import (
	"container/heap"
	_ "github.com/leesper/go_rng"
	"github.com/oleiade/lane"
	"math/rand"
)

const (
	BlockArrivalLambda = iota
	TransactionArrivalLambda
)

type (
	EventHeap []Event
	Time      int

	Blockchain struct {
		BlockArrival           func() int
		TransactionArrivalTime func() int
		Fee                    func() int
		TransactionMax         int
		History                []Block
		TransactionWaitQueue   *lane.Queue

		heap.Interface
	}

	Transaction struct {
		ArrivalTime Time
	}

	Block struct {
		Transactions []Transaction
	}

	Event interface {
		Visit(*Blockchain)
		When() int
	}

	BlockArrival       struct{ Time }
	TransactionArrival struct{ Time }
	Exit               struct{ Time }
)

func NewBlockchain(s rand.Source) *Blockchain {

	random := rand.New(s)
	_ = random
	eventHeap := make(EventHeap, 0)

	return &Blockchain{
		nil,
		nil,
		nil,
		0,
		make([]Block, 10),
		lane.NewQueue(),
		&eventHeap,
	}
}

func (b BlockArrival) When() int       { return b.Time.Time() }
func (b TransactionArrival) When() int { return b.Time.Time() }
func (b Exit) When() int               { return b.Time.Time() }

func (t Time) Time() int { return int(t) }

func (h EventHeap) Len() int           { return len(h) }
func (h EventHeap) Less(i, j int) bool { return h[i].When() < h[j].When() }
func (h EventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Event))
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
