package gopherCoin

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Timestamp     time.Time
	PreviousBlock *Block
	PreviousHash  string
	Data          *Data
	Hash          string
	Nonce         int64
}

func (b *Block) calculateHash() string {
	timeCreated := []byte(b.Timestamp.String())
	transactionAmountData := []byte(strconv.FormatInt(b.Data.Amount, 10))
	nusedOnce := []byte(strconv.FormatInt(b.Nonce, 10))

	unHashedData := bytes.Join([][]byte{timeCreated, transactionAmountData, []byte(b.PreviousHash), nusedOnce}, []byte{})

	hash := sha256.Sum256(unHashedData)
	encodedHash := fmt.Sprintf("%x", hash)

	return encodedHash
}

func (b *Block) mineBlock(difficulty int) *Block {
	curHash := b.Hash[0:difficulty]

	desiredHashBuilder := strings.Builder{}
	for i := 0; i < difficulty; i++ {
		desiredHashBuilder.WriteString("0")

	}
	desiredHash := desiredHashBuilder.String()

	for curHash != desiredHash {
		b.Hash = b.calculateHash()
		curHash = b.Hash[0:difficulty]
		b.Nonce++
	}
	return b
}

func (b *Block) validateBlock(difficulty int) bool {
	return b.Hash == b.mineBlock(difficulty).Hash
}

type Data struct {
	TransactionID string
	Amount        int64
}
