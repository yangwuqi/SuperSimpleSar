package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type AlertInfo struct {
	TimeStamp string
	Info      []int
}

func alertInfoHandle(alertInfoNodeData []int, timeStamp string) {
	//create another bucket int boltDB to save alert info
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err1 := encoder.Encode(alertInfoNodeData)
	if err1 != nil {
		fmt.Println("error when encoding!")
		panic(err1)
	}

	err2 := addDbBucketAsKeyIsTime(defaultBoltDB, alertInfoBucket, timeStamp, buffer.Bytes())
	if err2 != nil {
		panic(err2)
	}
}

func alertInfoRangeRead(timeStamp1, timeStamp2 string) []AlertInfo {
	timeStart := time.Now().UnixNano()
	err1, rangeData := getDbBucketRangeData(defaultBoltDB, alertInfoBucket, timeStamp1, timeStamp2)
	timeEnd := time.Now().UnixNano()
	timeUsed := timeEnd - timeStart
	fmt.Println("used ", timeUsed, " nano time to query")
	if err1 != nil {
		panic(err1)
	}
	return decodingDataToAlertInfo(rangeData)
}

func alertInfoAllRead() []AlertInfo {
	timeStart := time.Now().UnixNano()
	err1, rangeData := getDbBucketAllData2(defaultBoltDB, alertInfoBucket)
	//fmt.Println("len(rangeData): ",len(rangeData))
	timeEnd := time.Now().UnixNano()
	timeUsed := timeEnd - timeStart
	fmt.Println("used ", timeUsed, " nano time to query")
	if err1 != nil {
		panic(err1)
	}
	return decodingDataToAlertInfo(rangeData)
}

func decodingDataToAlertInfo(rangeData []BucketKeyValue) []AlertInfo {
	var result []AlertInfo
	for _, timeAndData := range rangeData {
		var alertInfo []int
		buffer := bytes.NewBuffer(timeAndData.Value)
		dec := gob.NewDecoder(buffer)
		err3 := dec.Decode(&alertInfo)
		if err3 != nil {
			fmt.Println("error the decoding to []int")
			break
			//panic(err3)
		}
		var resultSingle AlertInfo
		resultSingle.TimeStamp = string(timeAndData.Key)
		resultSingle.Info = alertInfo
		//fmt.Println("resultSingle in the decoding alert info: ",resultSingle)
		result = append(result, resultSingle)
	}
	return result
}

func alertInfoSumRangeQuery(timeStamp1, timeStamp2 string) {
	//alertInfoChosenData := alertInfoRead(timeStamp1,timeStamp2)
}

func alertInfoSumAllQuery() []int {
	alertInfoChosenData := alertInfoAllRead()
	//fmt.Println("len(alertInfoChosenData): ",len(alertInfoChosenData))
	var dataToSegmentTree [][]int
	var result []int
	for _, v := range alertInfoChosenData {
		dataToSegmentTree = append(dataToSegmentTree, v.Info)
	}
	alertInfoOperation := segmentTreeConstructor(dataToSegmentTree)
	result = alertInfoOperation.SumRange(0, len(dataToSegmentTree))
	for index, v := range result {
		fmt.Println("the index of alertInfoSum: ", index, ", the all sum result by segmentTree: ", v)
	} //prepare to change
	return result
}
