package gopherCoin

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Blockchain struct {
	GenesisBlock *Block
	LatestBlock  *Block
	difficulty   int
}

func (bchain *Blockchain) startBlockchain() {
	if bchain.difficulty <= 0 {
		bchain.difficulty = 2
	}
	gBlock := &Block{
		Timestamp:     time.Now(),
		PreviousBlock: nil,
		PreviousHash:  "",
	}

	data := &Data{
		TransactionID: uuid.New().String(),
		Amount:        0,
	}
	gBlock.Data = data

	//timestamp + amount (0)
	timeCreated := []byte(gBlock.Timestamp.String())
	transactionAmountData := []byte(strconv.FormatInt(gBlock.Data.Amount, 10))

	unHashedData := bytes.Join([][]byte{timeCreated, transactionAmountData}, []byte{})

	hash := sha256.Sum256(unHashedData)
	encodedHash := fmt.Sprintf("%x", hash)

	gBlock.Hash = encodedHash

	bchain.GenesisBlock = gBlock
	bchain.LatestBlock = gBlock
}

// Adds new block to blockchain, takes in transaction amount calls mineblock to add a valid block to bchain
func (bchain *Blockchain) newBlock(amount int64) {
	data := &Data{
		TransactionID: uuid.New().String(),
		Amount:        amount,
	}
	block := &Block{
		Timestamp:     time.Now(),
		PreviousBlock: bchain.getLatestBlock(),
		PreviousHash:  bchain.getLatestBlock().Hash,
		Data:          data,
	}
	//mine the block
	block.mineBlock(bchain.difficulty)
	//add it to the blockchain
	bchain.LatestBlock = block
}

func (bchain *Blockchain) isChainValid() bool {
	curBlock := bchain.getLatestBlock()

	for curBlock.PreviousBlock != nil {
		if !curBlock.validateBlock(bchain.difficulty) {
			return false
		}
		if curBlock.PreviousHash != curBlock.PreviousBlock.Hash {
			return false
		}
		curBlock = curBlock.PreviousBlock
	}
	return true
}

func (bchain *Blockchain) getLatestBlock() *Block {
	return bchain.LatestBlock
}
