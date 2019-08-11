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

	fmt.Println(len(rangeData))

	for _, timeAndData := range rangeData {
		metrics := make(map[string]string)
		buffer := bytes.NewBuffer(timeAndData.Value)
		dec := gob.NewDecoder(buffer)
		err4 := dec.Decode(&metrics)
		if err4 != nil {
			panic(err4)
		}
		fmt.Println()
		fmt.Println(timeAndData.Key)
		fmt.Println()
		for k, v := range metrics {
			fmt.Println("key: ", k, "value: ", v)
		}
	}
}

func rangeQuery(timeStamp1, timeStamp2 string) []TimeAndData {

	timeStart := time.Now().UnixNano()
	err2, rangeData := getDbBucketRangeData(defaultBoltDB, defaultBucket, timeStamp1, timeStamp2)
	timeEnd := time.Now().UnixNano()
	timeUsed := timeEnd - timeStart
	fmt.Println("used ", timeUsed, " nano time to query")
	if err2 != nil {
		panic(err2)
	}

	var result []TimeAndData

	for _, timeAndData := range rangeData {
		metrics := make(map[string]string)
		buffer := bytes.NewBuffer(timeAndData.Value)
		dec := gob.NewDecoder(buffer)
		err3 := dec.Decode(&metrics)
		if err3 != nil {
			panic(err3)
		}
		var resultSingle TimeAndData
		resultSingle.Data = make(map[string]string)
		resultSingle.Timestamp = timeAndData.Key
		resultSingle.Data = metrics
		result = append(result, resultSingle)
	}
	return result
}