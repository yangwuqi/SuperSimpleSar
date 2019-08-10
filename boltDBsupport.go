package main

import (
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

type BucketKeyValue struct{
	Key string
	Value []byte
}

func openBoltDB(dbName string)(*bolt.DB,error){
	db,err := bolt.Open(dbName, 0600, nil)
	if err != nil{
		return nil, err
	}
	fmt.Println("open the boltDB named ",dbName)
	return db, nil
}

func createDbBucket(db *bolt.DB, bucketName string)error{
	err := db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(bucketName))
		if b == nil{
			_,err1 := tx.CreateBucket([]byte(bucketName))
			if err1!=nil{
				panic(err1)
			}
		}
		return nil
	})
	if err!=nil{
		return err
	}
	fmt.Println("open the boltDB Bucket named ", bucketName)
	return nil
}

func updateDbBucketAsKeyIsTime(db *bolt.DB, bucketName string, data []byte)error{
	err := db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(bucketName))
		if b == nil{
			_,err1 := tx.CreateBucket([]byte(bucketName))
			if err1!=nil{
				panic(err1)
			}
			b = tx.Bucket([]byte(bucketName))
		}
		err2 := b.Put([]byte(time.Now().Format("2006/1/2 15:04:05")),data)
		if err2 != nil{
			panic(err2)
		}
		return nil
	})
	if err!=nil{
		return err
	}
	return nil
}

func getDbBucketDataByKey(db *bolt.DB, bucketName string, key []byte)(error,[]byte){
	var data []byte
	err := db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(bucketName))
		if b!=nil{
			data = b.Get(key)
		}else{
			panic("bucket does not exist!")
		}
		return nil
	})
	if err!=nil{
		return err, data
	}
	return nil, data
}

func getDbBucketAllData(db *bolt.DB, bucketName string)(error,[]BucketKeyValue){
	var result []BucketKeyValue
	err := db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(bucketName))
		err1 := b.ForEach(func(k, v []byte)error{
			result = append(result, BucketKeyValue{string(k), v})
			return nil
		})
		if err1 != nil{
			panic(err1)
		}
		return nil
	})
	if err!=nil{
		return err, result
	}
	return nil, result
}

func getDbBucketAllData2(db *bolt.DB, bucketName string)(error,[]BucketKeyValue){
	var result []BucketKeyValue
	err := db.View(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k!=nil; k,v = c.Next(){
			result=append(result, BucketKeyValue{string(k), v})
		}
		return nil
	})
	if err!=nil{
		return err, result
	}
	return nil, result
}

func getDbBucketRangeData(db *bolt.DB, bucketName, keyStart ,keyEnd string)(error, []BucketKeyValue){
	var result []BucketKeyValue
	err := db.View(func(tx *bolt.Tx)error{
		c := tx.Bucket([]byte(bucketName)).Cursor()
		for k,v := c.Seek([]byte(keyStart)); k!=nil&&bytes.Compare(k, []byte(keyEnd))<=0; k,v=c.Next(){
			result=append(result,BucketKeyValue{string(k),v})
		}
		return nil
	})
	if err!=nil{
		return err, result
	}
	return nil, result
}