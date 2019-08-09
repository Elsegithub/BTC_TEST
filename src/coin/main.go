package main

import (
	"core"
)

func main() {
	bc := core.NewBlockChain()
	defer bc.DB.Close()

	cli := core.CLI{bc}
	cli.Run()
	//bc.AddBlock("send 1 BTC to ME")
	//bc.AddBlock("send 2 more BTC to ME")
	//for _, block := range bc.Blocks {
	//	fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	pow := core.NewProofOfWork(block)
	//	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
	//	fmt.Println()
	//}
}
