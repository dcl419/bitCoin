package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("小明给小红转了一个比特币")
	bc.AddBlock("张三给李四转了一个比特币")
	for i, block := range bc.blocks {
		// %x 打印十六进制
		fmt.Printf("当前区块高度: %d, %x, %x, %s\n", i, block.PreHash, block.Hash, block.Data)
	}

}
