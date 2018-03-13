// Copyright (C) 2017, Zipper Team.  All rights reserved.
//
// This file is part of zipper
//
// The zipper is free software: you can use, copy, modify,
// and distribute this software for any purpose with or
// without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// The zipper is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// ISC License for more details.
//
// You should have received a copy of the ISC License
// along with this program.  If not, see <https://opensource.org/licenses/isc>.


package proto

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/zipper-project/zipper/account"
	"github.com/zipper-project/zipper/common/crypto"
	"github.com/zipper-project/zipper/common/utils"
)

var (
	testHashStr string
)

func TestBlockSerialize(t *testing.T) {
	var (
		testBlock = Block{
			Header: &BlockHeader{
				PreviousHash: crypto.DoubleSha256([]byte("xxxx")).String(),
				TimeStamp:    uint32(time.Now().Unix()),
				Nonce:        uint32(100),
			},
		}
	)
	Txs := make([]*Transaction, 0)
	hashs := make([]crypto.Hash, 0)
	reciepent := account.HexToAddress("0xbf6080eaae18a6eb4d9d3b9ef08a8bdf02e3caa8")
	for i := 1; i < 3; i++ {
		tx := &Transaction{
			Header: &TxHeader{
				FromChain:  account.NewChainCoordinate([]byte{byte(i)}),
				ToChain:    account.NewChainCoordinate([]byte{byte(i)}),
				Type:       TransactionType_Atomic,
				Nonce:      1,
				Sender:     reciepent.String(),
				Recipient:  reciepent.String(),
				AssetID:    1,
				Amount:     big.NewInt(1100).Int64(),
				Fee:        big.NewInt(110).Int64(),
				CreateTime: utils.CurrentTimestamp(),
			},
		}
		Txs = append(Txs, tx)
		hashs = append(hashs, tx.Hash())
	}

	testBlock.Transactions = Txs
	testBlock.Header.TxsMerkleHash = crypto.ComputeMerkleHash(hashs)[0].String()

	fmt.Println("Block hash", testBlock.Hash())
	fmt.Printf("Block Raw {'previousHash':%v,\n 'MerkleHash':%v,\n  'Nonce':%v,\n TimeStamp':%v,\n Txs:%v.\n",
		testBlock.Header.PreviousHash,
		testBlock.Header.TxsMerkleHash,
		testBlock.Header.Nonce,
		testBlock.Header.TimeStamp,
		testBlock.Transactions,
	)
	fmt.Println("Block Header serialize()", testBlock.Header.Serialize())
	fmt.Println("Block AtomicTxs", testBlock.Transactions, len(testBlock.Transactions))
	fmt.Println("Block serialize()", hex.EncodeToString(testBlock.Serialize()))
	testHashStr = hex.EncodeToString(testBlock.Serialize())
}

func TestBlockDeserialize(t *testing.T) {
	testBlock := &Block{}
	data, _ := hex.DecodeString(testHashStr)
	fmt.Println("------------block deseriailze---------")
	testBlock.Deserialize(data)
	blkData := testBlock.Serialize()
	if !bytes.Equal(blkData, data) {
		t.Errorf("Block.Serialize error, %0x != %0x ", data, blkData)
	}
}
