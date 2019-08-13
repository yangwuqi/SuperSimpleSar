package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

func rangeQueryPrint() {

	fmt.Println("input the timestamps between which you want to query")

	var timeStamp1, timeStamp2 string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("timestamp1: ")
	stdIn1, _, _ := reader.ReadLine()
	timeStamp1 = string(stdIn1)

	fmt.Println("timestamp2: ")
	stdIn2, _, _ := reader.ReadLine()
	timeStamp2 = string(stdIn2)

	timeStart := time.Now().UnixNano()
	err3, rangeData := getDbBucketRangeData(defaultBoltDB, defaultBucket, timeStamp1, timeStamp2)
	timeEnd := time.Now().UnixNano()
	timeUsed := timeEnd - timeStart
	fmt.Println("used ", timeUsed, " nano time to query")
	if err3 != nil {
		panic(err3)
	}

	//fmt.Println(len(rangeData))

	for _, timeAndData := range rangeData {
		metrics := make(map[string]string)
		buffer := bytes.NewBuffer(timeAndData.Value)
		dec := gob.NewDecoder(buffer)
		err4 := dec.Decode(&metrics)
		if err4 != nil {
			break
			//panic(err4)
		}
		//fmt.Println()
		//fmt.Println(string(timeAndData.Key))
		//fmt.Println()
		//for k, v := range metrics {
		//	fmt.Println("key: ", k, "value: ", v)
		//}
	}
	timeEnd2 := time.Now().UnixNano()
	timeUsed2 := timeEnd2 - timeStart
	fmt.Println("used ", timeUsed2, " nano time to query and decoding")
}

func rangeQuery(timeStamp1, timeStamp2 string) []TimeAndData {

	timeStart := time.Now().UnixNano()
	err2, rangeData := getDbBucketRangeData(defaultBoltDB, defaultBucket, timeStamp1, timeStamp2)
	timeEnd := time.Now().UnixNano()
	timeUsed := timeEnd - timeStart
	fmt.Println("used ", timeUsed, " nano time to query range data")
	if err2 != nil {
		panic(err2)
	}

	return decodingDataToTimeAndMap(rangeData, timeStart)
}

func allQuery() []TimeAndData {

	timeStart := time.Now().UnixNano()
	err2, rangeData := getDbBucketAllData2(defaultBoltDB, defaultBucket)
	timeEnd := time.Now().UnixNano()
	timeUsed := timeEnd - timeStart
	fmt.Println("used ", timeUsed, " nano time to query all data")
	if err2 != nil {
		panic(err2)
	}

	return decodingDataToTimeAndMap(rangeData, timeStart)
}

func decodingDataToTimeAndMap(rangeData []BucketKeyValue, timeStart int64) []TimeAndData {
	var result []TimeAndData

	for _, timeAndData := range rangeData {
		metrics := make(map[string]string)
		buffer := bytes.NewBuffer(timeAndData.Value)
		dec := gob.NewDecoder(buffer)
		err3 := dec.Decode(&metrics)
		if err3 != nil {
			break
			//panic(err3)
		}
		var resultSingle TimeAndData
		resultSingle.Data = make(map[string]string)
		resultSingle.Timestamp = string(timeAndData.Key)
		resultSingle.Data = metrics
		result = append(result, resultSingle)
	}
	timeEnd := time.Now().UnixNano()
	timeUsed := timeEnd - timeStart
	fmt.Println("used ", timeUsed, " nano time to query range data and decoding")
	return result
}
