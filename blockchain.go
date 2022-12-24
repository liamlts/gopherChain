/*
***
GopherChain
2022 By: https://github.com/liamlts
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
***
*/
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

	//timestamp + amount (0) + init message
	timeCreated := []byte(gBlock.Timestamp.String())
	transactionAmountData := []byte(strconv.FormatInt(gBlock.Data.Amount, 10))
	initMessage := "started gopherChain"

	unHashedData := bytes.Join([][]byte{timeCreated, transactionAmountData, []byte(initMessage)}, []byte{})

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

// Returns false and malicious block upon failure. Else will return true and nil
func (bchain *Blockchain) isChainValid() (bool, *Block) {
	curBlock := bchain.getLatestBlock()

	for curBlock.PreviousBlock != nil {
		if !curBlock.validateBlock(bchain.difficulty) {
			return false, curBlock
		}
		if curBlock.PreviousHash != curBlock.PreviousBlock.Hash {
			return false, curBlock
		}
		curBlock = curBlock.PreviousBlock
	}
	return true, nil
}

func (bchain *Blockchain) getLatestBlock() *Block {
	return bchain.LatestBlock
}
