package main

import "fmt"

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
		Hash:    []byte{}, // 先填空, 后面再计算 //TODO
		Data:    []byte(data),
	}

	return &block
}

//6. 生成哈希
//7. 引入区块链
//8. 添加区块
//9. 重构代码

func main() {
	block := NewBlock("转一个bit", []byte{})
	// %x 打印十六进制
	fmt.Printf("%x, %x, %s", block.PreHash, block.Hash, block.Data)
}
