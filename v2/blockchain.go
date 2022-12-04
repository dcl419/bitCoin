package main

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
