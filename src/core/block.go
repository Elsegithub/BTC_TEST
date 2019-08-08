package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp	int64 //时间戳
	Data		[]byte  //数据
	PrevBlockHash []byte  //前一块hash
	Hash		[]byte  //hash
}

func NewBlock(data string, prevBlockHash []byte) *Block{
	block := Block{Timestamp:time.Now().Unix(), Data:[]byte(data), PrevBlockHash:prevBlockHash, Hash:[]byte{}}
	block.SetHash()
	return &block
}

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewFirstBlock() *Block{
	return  NewBlock("first block", []byte{})
}