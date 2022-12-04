package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"
const lastHashKey = "LastHashKey"

// 7. 引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	// blocks []*Block
	db   *bolt.DB
	tail []byte // 存储最后一个区块的hash
}

// 7.1 定义一个区块链
func NewBlockChain() *BlockChain {
	var lastHash []byte

	//return &BlockChain{
	//	// 使用创世块初始化区块链
	//	blocks: []*Block{genesisBlock},
	//}

	// 1. 打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil) // 读写权限
	if err != nil {
		log.Panic("数据库打开失败")
	}
	// defer db.Close()

	// 2. 写 - 使用update - 找到db或者创建db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			// 没有bucket - 需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket（blockBucket）失败")
			}

			// 创世块
			genesisBlock := GenesisBlock()

			// 写数据
			// key - hash
			// value - block的字节流
			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			bucket.Put([]byte(lastHashKey), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte(lastHashKey))
		}
		return nil
	})

	return &BlockChain{db, lastHash}
}

// 7.2 创世块
func GenesisBlock() *Block {
	return NewBlock("创世块", []byte{})
}

// 8. 添加区块
func (bc *BlockChain) AddBlock(data string) {
	/*
		// 0. 获取前区块hash
		// 0.1 获取最后一个区块
		lastBlock := bc.blocks[len(bc.blocks)-1]
		// 新增区块的前区块hash
		preHash := lastBlock.Hash

		// 1. 创建新的区块
		block := NewBlock(data, preHash)
		// 2. 添加到区块链数组中
		bc.blocks = append(bc.blocks, block)
	*/
}
