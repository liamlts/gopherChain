package gopherCoin

import "errors"

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

func (bchain *Blockchain) hashLookup(hash string) (*Block, error) {
	curBlock := bchain.getLatestBlock()

	for curBlock.PreviousBlock != nil {
		if hash == curBlock.Hash {
			return curBlock, nil
		}

		curBlock = curBlock.PreviousBlock
	}
	return nil, errors.New("error finding block with specified hash")
}

func (bchain *Blockchain) tIDLookup(tID string) (*Block, error) {
	curBlock := bchain.getLatestBlock()

	for curBlock.PreviousBlock != nil {
		if tID == curBlock.Data.TransactionID {
			return curBlock, nil
		}

		curBlock = curBlock.PreviousBlock
	}
	return nil, errors.New("error finding block with specified transaction ID")

}
