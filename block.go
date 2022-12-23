/****
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
	"math/rand"
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
	RandInd       int
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

func (b *Block) mineBlock(difficulty int) {
	b.Hash = b.calculateHash()

	//Desired hash is (00) at random int 0-29 to 0-29+difficulty
	if b.RandInd == 0 {
		b.RandInd = rand.Intn(15)
	}

	curHash := b.Hash[b.RandInd : b.RandInd+difficulty]

	desiredHashBuilder := strings.Builder{}
	for i := 0; i < difficulty; i++ {
		desiredHashBuilder.WriteString("0")

	}
	desiredHash := desiredHashBuilder.String()

	for curHash != desiredHash {
		b.Hash = b.calculateHash()
		curHash = b.Hash[b.RandInd : b.RandInd+difficulty]
		b.Nonce++
	}
	b.Nonce = 0
}

func (b *Block) validateBlock(difficulty int) bool {
	initHash := b.Hash
	b.mineBlock(difficulty)

	return initHash == b.Hash
}

type Data struct {
	TransactionID string
	Amount        int64
}
