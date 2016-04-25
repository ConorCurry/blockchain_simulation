package simulation

func (block *BlockArrival) Visit(chain *Blockchain) {
	// Adjust the lambda value
	// Write out the new block to the blockchain
	// Generate a new block arrival event
}

func (tArrival TransactionArrival) Visit(chain *Blockchain) {
	// Add a new transaction to the T list
	// Generate a new transaction arrival event
}

func (*Exit) Visit(chain *Blockchain) {
	// Print out all of the results from the experiment
}
