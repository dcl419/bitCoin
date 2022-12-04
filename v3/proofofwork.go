package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 工作量证明

type ProofOfWork struct {
	// a. block
	block *Block
	// b. 目标值
	// 一个非常大的数字 他有丰富的方法：比较，赋值
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	// 指定一个难度值
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	// 辅助变量，将难度值转化为bigInt，指定16进制格式
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	// 1. 拼装数据（区块数据+随机数）
	// 2. 做hash运算
	// 3. 与pow的target进行比较
	//  3.1 找到了，退出返回
	//  3.2 没找到，继续找，随机数+1
	block := pow.block
	var nonce uint64
	var hash [32]byte
	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PreHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		// 将二维切片数组连接起来，返回一个一维切片
		blockInfo := bytes.Join(tmp, []byte{})
		// hash运算
		hash = sha256.Sum256(blockInfo)
		// 与pow中的target进行比较
		temInt := big.Int{}
		// 将hash数转化为big.int
		temInt.SetBytes(hash[:])
		// 比较

		// 当前 < 目标  就是满足条件
		//	-1 if x <  y
		//	 0 if x == y
		//	+1 if x >  y
		if temInt.Cmp(pow.target) == -1 {
			// 找到了
			fmt.Printf("挖矿成功: hash : %x, nonce : %d\n", hash, nonce)
			break
		} else {
			// 没找到 - 随机数+1
			if nonce%100000 == 0 {
				// fmt.Printf("挖矿失败: hash : %x, nonce : %d\n", hash, nonce)
			}
			nonce++
		}

	}
	// return []byte("helloworld"), 10
	return hash[:], nonce
}
