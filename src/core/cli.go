package core

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Bc *BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("printchain - print all the blocks of the blockchain")
}

func (cli *CLI) vaildateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("success!")
}

func (cli *CLI) Run() {
	cli.vaildateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printchainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printchainCmd.Parsed() {
		cli.printchain()
	}
}

func (cli *CLI) printchain() {
	cli.Bc.showBlock()
}
