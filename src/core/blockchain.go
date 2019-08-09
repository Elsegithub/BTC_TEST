package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type BlockChain struct {
	tip []byte
	DB  *bolt.DB
}

func (bc *BlockChain) AddBlock(data string) {
	//prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	//newBlock := NewBlock(data, prevBlock.Hash)
	//bc.Blocks = append(bc.Blocks, newBlock)
	var lastHash []byte
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("Creating a new one")
			firstBlock := NewFirstBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(firstBlock.Hash, firstBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), firstBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = firstBlock.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{tip, db}
	return &bc
}

func (bc *BlockChain) showBlock() {
	var b *Block
	err := bc.DB.View(func(tx *bolt.Tx) error {
		bu := tx.Bucket([]byte(blocksBucket))
		bu.ForEach(func(hash, data []byte) error {
			//fmt.Printf("%x\n",hash)
			//fmt.Printf("%x\n",[]byte("l"))
			if string(hash) == string([]byte("l")) {
				return nil
			}
			b = DeSerializeBlock(data)
			fmt.Printf("Prev.hash: %x\n", b.PrevBlockHash)
			fmt.Printf("data: %s\n", b.Data)
			fmt.Printf("Hash: %x\n", b.Hash)
			return nil
		})
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
