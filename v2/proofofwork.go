package main

import "math/big"

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
