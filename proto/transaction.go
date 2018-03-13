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
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/zipper-project/zipper/account"
	"github.com/zipper-project/zipper/common/crypto"
)

var (
	// ErrEmptySignature represents no signature
	ErrEmptySignature = errors.New("Signature Empty Error")
)

//NewTransaction initialization transaction
func NewTransaction(from, to account.ChainCoordinate,
	txType TransactionType,
	nonce uint32,
	sender, recipient account.Address,
	assetID uint32,
	amount, fee int64,
	createTime uint32) *Transaction {
	tx := &Transaction{
		Header: &TxHeader{
			FromChain:  from,
			ToChain:    to,
			Type:       txType,
			Nonce:      nonce,
			Sender:     sender.String(),
			Recipient:  recipient.String(),
			AssetID:    assetID,
			Amount:     amount,
			Fee:        fee,
			CreateTime: createTime,
		},
	}
	return tx
}

// Hash returns the hash of a transaction
func (tx *Transaction) Hash() crypto.Hash {
	return crypto.DoubleSha256(tx.Serialize())
}

// SignHash returns the hash of a raw transaction before sign
func (tx *Transaction) SignHash() crypto.Hash {
	rawTx := &Transaction{
		Header: &TxHeader{
			FromChain:  tx.Header.FromChain,
			ToChain:    tx.Header.ToChain,
			Type:       tx.Header.Type,
			Nonce:      tx.Header.Nonce,
			Sender:     tx.Header.Sender,
			Recipient:  tx.Header.Recipient,
			AssetID:    tx.Header.AssetID,
			Amount:     tx.Header.Amount,
			Fee:        tx.Header.Fee,
			CreateTime: tx.Header.CreateTime,
		},
		Payload:      tx.Payload,
		Meta:         tx.Meta,
		ContractSpec: tx.ContractSpec,
	}
	return rawTx.Hash()
}

// Serialize marshal txData proto message
func (tx *Transaction) Serialize() []byte {
	bytes, _ := proto.Marshal(tx)
	return bytes
}

// Deserialize deserializes bytes to a transaction
func (tx *Transaction) Deserialize(data []byte) error {
	return proto.Unmarshal(data, tx)
}

// Verfiy Also can use this method verify signature
func (tx *Transaction) Verfiy() (account.Address, error) {
	var (
		a   account.Address
		err error
	)
	switch tx.Header.GetType() {
	case TransactionType_Atomic, TransactionType_AcrossChain, TransactionType_Backfront, TransactionType_Distribut, TransactionType_IssueUpdate,
		TransactionType_JSContractInit, TransactionType_LuaContractInit, TransactionType_ContractInvoke, TransactionType_ContractQuery, TransactionType_Security:
		fallthrough
	case TransactionType_Issue:
		if tx.Header.Signature != nil {
			sig := &crypto.Signature{}
			sig.SetBytes(tx.Header.Signature, false)
			p, err := sig.RecoverPublicKey(tx.SignHash().Bytes())
			if err != nil {
				return a, err
			}
			a = account.PublicKeyToAddress(*p)
		} else {
			err = ErrEmptySignature
		}

	case TransactionType_Merged:
		a = account.ChainCoordinateToAddress(account.HexToChainCoordinate(tx.FromChain()))
	}
	return a, err
}

// Sender returns the address of the sender.
func (tx *Transaction) Sender() account.Address {
	return account.HexToAddress(tx.Header.Sender)
}

// FromChain returns the chain coordinate of the sender
func (tx *Transaction) FromChain() string {
	return account.NewChainCoordinate(tx.Header.FromChain).String()
}

// ToChain returns the chain coordinate of the recipient
func (tx *Transaction) ToChain() string {
	return account.NewChainCoordinate(tx.Header.ToChain).String()
}

// IsLocalChain returns whether or not local chain
func (tx *Transaction) IsLocalChain() bool {
	return bytes.Compare([]byte(tx.FromChain()), []byte(tx.ToChain())) == 0
}

// Recipient returns the address of the recipient
func (tx *Transaction) Recipient() account.Address {
	return account.HexToAddress(tx.Header.Recipient)
}

// AssetID returns the transfer asset of the transaction
func (tx *Transaction) AssetID() uint32 {
	return tx.Header.AssetID
}

// Amount returns the transfer amount of the transaction
func (tx *Transaction) Amount() int64 { return tx.Header.Amount }

// Fee returns the nonce of the transaction
func (tx *Transaction) Fee() int64 { return tx.Header.Fee }

// Nonce returns the nonce of the transaction
func (tx *Transaction) Nonce() uint32 { return tx.Header.Nonce }

// GetType returns transaction type
func (tx *Transaction) GetType() TransactionType { return tx.Header.Type }

// CreateTime returns the create time of the transaction
func (tx *Transaction) CreateTime() uint32 { return tx.Header.CreateTime }

// Compare implements interface consensus need
func (tx *Transaction) Compare(v interface{}) int {
	if tx.CreateTime() >= v.(*Transaction).CreateTime() {
		return 1
	}
	return 0
}

// Transactions represents transaction slice type for basic sorting.
type Transactions []*Transaction

// Len returns the length of s
func (s Transactions) Len() int { return len(s) }

// Swap swaps the i'th and the j'th element in s
func (s Transactions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less compares nonce of the i'th and the j'th element in s
func (s Transactions) Less(i, j int) bool { return s[i].Header.Nonce < s[j].Header.Nonce }
