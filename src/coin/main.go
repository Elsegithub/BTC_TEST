package main

import (
	"core"
	"fmt"
)

func main()  {
	bc := core.NewBlockChain()
	bc.AddBlock("send 1 BTC to ME")
	bc.AddBlock("send 2 more BTC to ME")
	for _, block := range bc.Blocks{
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
