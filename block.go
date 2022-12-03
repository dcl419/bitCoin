package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

// 1. 定义结构
type Block struct {
	// 1. 版本号
	Version uint64
	//2. 前区块哈希
	PreHash []byte
	// 3. 梅克尔根 hash值
	MerkelRoot []byte
	// 4. 时间戳
	TimeStamp uint64
	// 5. 难度值
	Difficulty uint64
	// 6. 随机数
	Nonce uint64

	//7. 当前区块哈希 - 正常区块链中不存在该字段，为了方便，做了简化
	Hash []byte
	//8. 数据
	Data []byte
}

// 5. 创建区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PreHash:    preBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, // 暂时随便先填充，无效值
		Nonce:      0, // 暂时随便先填充，无效值
		Hash:       []byte{},
		Data:       []byte(data),
	}

	block.SetHash()
	return &block
}

// uint64 -> []byte
func Uint64ToByte(number uint64) []byte {
	var buffer bytes.Buffer
	// binary.BigEndian 大端对齐
	// 大端 - 高位在前 低位在后
	// 小端 - 低位在前 高位在后
	err := binary.Write(&buffer, binary.BigEndian, number)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// 6. 生成哈希
func (block *Block) SetHash() {
	var blockInfo []byte
	// 1. 拼装数据
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PreHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	// block.Data...  将block.Data打散成byte
	blockInfo = append(blockInfo, block.Data...)
	// 2. sha256
	hash := sha256.Sum256(blockInfo)

	block.Hash = hash[:]
}
