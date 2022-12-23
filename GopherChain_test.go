package gopherCoin

import (
	"testing"
)

func TestBlockChain(t *testing.T) {
	//make new block chain
	gopherCoin := new(Blockchain)
	gopherCoin.startBlockchain()
	gopherCoin.difficulty = 4
	//manually add new blocks to it
	gopherCoin.newBlock(90)
	gopherCoin.newBlock(90)
	gopherCoin.newBlock(90)
	gopherCoin.newBlock(90)
	gopherCoin.newBlock(90)
	gopherCoin.newBlock(90)
	gopherCoin.newBlock(90)
	//use channel to walk the blockchain
	c := make(chan *Block)
	go gopherCoin.ReverseWalk(c)

	//display all blocks by reading off channel
	num := 1
	for i := range c {
		t.Logf("Num: %d Block: %v\n", num, i)
		num++
	}

}
