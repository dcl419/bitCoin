package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// strings test.db   查看文件内容
	var dbName = "./test.db"
	var bucketName = []byte("b1")

	// 1. 打开数据库
	db, err := bolt.Open(dbName, 0600, nil) // 读写权限
	if err != nil {
		log.Panic("数据库打开失败")
	}

	// 2. 写 - 使用update - 找到db或者创建db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			// 没有bucket - 需要创建
			bucket, err = tx.CreateBucket(bucketName)
			if err != nil {
				log.Panic("创建bucket（b1）失败")
			}
		}

		bucket.Put([]byte("name"), []byte("dcl"))
		bucket.Put([]byte("sb"), []byte("wanger"))
		return nil
	})

	// 3. 读 - view
	db.View(func(tx *bolt.Tx) error {
		// 1. 找到抽屉 - 没有直接报错退出
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			log.Panic("未找到bucket，请检查")
		}
		// 2. 有的话直接读取数据
		value := bucket.Get([]byte("name"))
		fmt.Printf("name: %s", value)
		return nil
	})

}
