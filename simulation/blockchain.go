package simulation

import (
	"container/heap"
	rng "github.com/leesper/go_rng"
	"github.com/oleiade/lane"
	"math"
	"sync"
)

const (
	BlockArrivalLambda = 0.00174520069808
	TxAlpha            = 3.0
	TxBeta             = 1.0
	FeeAlpha           = 7.0
	FeeBeta            = 1200.0
	MaxTx              = 2048
)

type (
	EventHeap []Event
	Time      int

	Blockchain struct {
		BlockArrival             func() int
		TransactionArrivalLambda func() float64
		Fee                      func() int
		TxArrival                func(float64) int
		TxLambda                 float64

		TransactionMax       int
		History              []Block
		TransactionWaitQueue *lane.PQueue

		CurrentBlock Block
		heap.Interface
	}

	Transaction struct {
		ArrivalTime Time `json:"time"`
		Fee         int  `json:"fee"`
	}

	Block struct {
		Time         `json:"time"`
		Transactions []Transaction `json:"transactions"`
	}

	Event interface {
		Visit(*Blockchain)
		When() int
	}

	BlockArrival       struct{ Time }
	TransactionArrival struct{ Time }
	Exit               struct {
		filename string
		Time
	}
)

func (chain *Blockchain) Run(wg *sync.WaitGroup, filename string) {

	defer wg.Done()
	// Make each of the original elements and add them to the heap
	blkArrival := BlockArrival{Time(chain.BlockArrival())}
	txArrival := TransactionArrival{Time(chain.TxArrival(chain.TxLambda))}
	exit := Exit{filename, Time(100000)}

	heap.Push(chain, blkArrival)
	heap.Push(chain, txArrival)
	heap.Push(chain, exit)

	for chain.Len() != 0 {
		event := heap.Pop(chain).(Event)
		switch event.(type) {
		case Exit:
			event.Visit(chain)
			return
		default:
			event.Visit(chain)
		}
	}

}

func NewBlockchain(seed int64) *Blockchain {

	expGen := rng.NewExpGenerator(seed)
	gammaGen := rng.NewGammaGenerator(seed)

	blkArrival := func() int {
		return round(expGen.Exp(BlockArrivalLambda))
	}

	txArrivalLambda := func() float64 {
		return gammaGen.Gamma(TxAlpha, TxBeta)
	}

	fee := func() int {
		return round(gammaGen.Gamma(FeeAlpha, FeeBeta)) + 11500
	}

	txArrival := func(lambda float64) int {
		return round(expGen.Exp(lambda))
	}

	eventHeap := make(EventHeap, 0)

	return &Blockchain{
		blkArrival,
		txArrivalLambda,
		fee,
		txArrival,
		txArrivalLambda(),
		MaxTx,
		make([]Block, 0),
		lane.NewPQueue(lane.MAXPQ),
		Block{
			Time(0),
			make([]Transaction, 0),
		},
		&eventHeap,
	}
}

func (chain *Blockchain) BlockIsFull() bool {
	return len(chain.CurrentBlock.Transactions) >= chain.TransactionMax
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

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.8, f))
}
