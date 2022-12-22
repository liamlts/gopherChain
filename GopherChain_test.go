package gopherCoin

import (
	"testing"
)

func TestBlockChain(t *testing.T) {
	//make new block chain
	gopherCoin := new(Blockchain)
	gopherCoin.startBlockchain()
	//manually add new blocks to it
	gopherCoin.newBlock(90)
	gopherCoin.newBlock(77489321)
	gopherCoin.newBlock(90890)
	gopherCoin.newBlock(902346798)
	gopherCoin.newBlock(93217483920)
	gopherCoin.newBlock(9324789)
	bID := gopherCoin.getLatestBlock().PreviousBlock.PreviousBlock.PreviousBlock
	t.Logf("%s\n", bID.Data.TransactionID)
	findBlock, err := gopherCoin.tIDLookup(bID.Data.TransactionID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Block: %v", findBlock)
	//use channel to walk the blockchain
	//c := make(chan *Block)
	//go gopherCoin.ReverseWalk(c)

	//display all blocks by reading off channel
	/****
	num := 1
	for i := range c {
		t.Logf("Num: %d Block: %v\n", num, i)
		num++
	}
	****/
	gopherCoin.isChainValid()

}
