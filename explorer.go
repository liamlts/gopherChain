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
