package gopherCoin

func (bchain *Blockchain) length() int64 {
	curBlock := bchain.getLatestBlock()
	var len int64

	for curBlock.PreviousBlock != nil {
		len++
		curBlock = curBlock.PreviousBlock

	}

	return len
}

// Send entire blockchain in reverse order on given channel
func (bchain *Blockchain) ReverseWalk(c chan *Block) {
	defer close(c)
	curBlock := bchain.getLatestBlock()

	for curBlock.PreviousBlock != nil {
		c <- curBlock
		curBlock = curBlock.PreviousBlock
	}
}
