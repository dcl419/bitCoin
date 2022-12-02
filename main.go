package main

import (
	"crypto/sha256"
	"fmt"
)

// 1. 定义结构
type Block struct {
	//2. 前区块哈希
	PreHash []byte
	//3. 当前区块哈希
	Hash []byte
	//4. 数据
	Data []byte
}

// 5. 创建区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		PreHash: preBlockHash,
		Hash:    []byte{},
		Data:    []byte(data),
	}

	block.SetHash()
	return &block
}

// 6. 生成哈希
func (block *Block) SetHash() {
	// 1. 拼装数据
	// block.Data...  将block.Data打散成byte
	blockInfo := append(block.PreHash, block.Data...)
	// 2. sha256
	hash := sha256.Sum256(blockInfo)

	block.Hash = hash[:]
}

// 7. 引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	blocks []*Block
}

// 7.1 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建创世块
	genesisBlock := GenesisBlock()
	return &BlockChain{
		// 使用创世块初始化区块链
		blocks: []*Block{genesisBlock},
	}
}

// 7.2 创世块
func GenesisBlock() *Block {
	return NewBlock("创世块", []byte{})
}

// 8. 添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 0. 获取前区块hash
	// 0.1 获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	// 新增区块的前区块hash
	preHash := lastBlock.Hash

	// 1. 创建新的区块
	block := NewBlock(data, preHash)
	// 2. 添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}

//9. 重构代码

func main() {
	bc := NewBlockChain()
	bc.AddBlock("小明给小红转了一个比特币")
	bc.AddBlock("张三给李四转了一个比特币")
	for i, block := range bc.blocks {
		// %x 打印十六进制
		fmt.Printf("当前区块高度: %d, %x, %x, %s\n", i, block.PreHash, block.Hash, block.Data)
	}

}
