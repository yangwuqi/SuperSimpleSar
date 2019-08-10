package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type dataPersist struct {
	Index int
	Time      string
	DataSaved map[string]string
}

func persisting(persistOK bool, persistDura *int, dataNewest *dataNew, closePersisting chan bool) {
	if !persistOK {
		fmt.Println("the persistence option is not open")
		return
	}

	for {
		select {
		case <-time.After(time.Duration((*persistDura)*1000) * time.Millisecond):
			if dataNewest!=nil {
				save(dataNewest)
			}
		case <-closePersisting:
			fmt.Println("the persisting process is shut down")
			return
		}
	}

}

/*
func save(path string, dataNewest dataNew) {
	dataNewest.mu.RLock()
	defer dataNewest.mu.RUnlock()

	var newPersist dataPersist
	newPersist.Time = time.Now().Format("2006/1/2 15:04:05")
	newPersist.DataSaved = dataNewest.data

	dataRead, _ := load(path)
	dataRead = append(dataRead, newPersist)

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err1 := encoder.Encode(dataRead)
	if err1 != nil {
		fmt.Println("error when encoding the []byte")
	}
	err3 := ioutil.WriteFile(path, buffer.Bytes(), 0644)
	if err3 != nil {
		fmt.Println("error when writing the file!")
	}
	//seesee()
}

func load(path string) ([]dataPersist, error) {
	var result []dataPersist
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return []dataPersist{}, err
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(&result)
	return result, nil
}

func seesee() {
	xx, _ := load("data")
	for i := 0; i < len(xx); i++ {
		fmt.Println(xx[i].Time, ":   ", xx[i].DataSaved)
	}
}

*/

func save11(path string, dataNewest *dataNew) {
	dataNewest.mu.RLock()
	defer dataNewest.mu.RUnlock()

	var newPersist dataPersist
	newPersist.Time = time.Now().Format("2006/1/2 15:04:05")
	newPersist.DataSaved = dataNewest.data
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	_ = encoder.Encode(newPersist)
	file, err1 := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Fatalf("failed creating file: %s", err1)
	}

	toBeWritten := append(buffer.Bytes(), '!', 'p', 'g', 'n', 'b')
	dataWriter := bufio.NewWriter(file)
	_, err2 := dataWriter.Write(toBeWritten)
	if err2 != nil {
		fmt.Printf("error when writing to the file!\n")
		panic(err2)
	}
	err3 := dataWriter.Flush()
	if err3 != nil {
		fmt.Printf("error when Flush() the file!\n")
		panic(err3)
	}
	err4 := file.Close()
	if err4 != nil {
		fmt.Printf("error when Close() the file!\n")
		panic(err4)
	}
}

func save(dataNewest *dataNew) {
	dataNewest.mu.RLock()
	defer dataNewest.mu.RUnlock()

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(dataNewest.data)
	if err!=nil{
		fmt.Println("error when encoding!")
		panic(err)
	}
	err1:=updateDbBucketAsKeyIsTime(defaultBoltDB, defaultBucket, buffer.Bytes())
	if err1!=nil{
		panic(err1)
	}
}

func loadAll()[]dataPersist{
	var result []dataPersist
	err,allDataFromBucket := getDbBucketAllData2(defaultBoltDB, defaultBucket)
	if err!=nil{
		panic(err)
	}
	for index,data:=range allDataFromBucket{
		if len(data.Value)==0{
			break
		}
		metrics:=make(map[string]string)
		buffer := bytes.NewBuffer(data.Value)
		dec := gob.NewDecoder(buffer)
		err1 := dec.Decode(&metrics)
		//fmt.Println()//
		//fmt.Println(index," ",string(data.Key))//
		//fmt.Println()//
		//fmt.Println(metrics)//
		if err1 != nil {
			fmt.Printf("error when Decoding!\n")
			panic(err1)
		}
		result=append(result, dataPersist{index, string(data.Key), metrics})
	}
	return result
}

func load11(path string) ([]dataPersist, error) {
	var result []dataPersist
	flag := 0
	file, _ := os.Open(path)
	fileBytes, _ := ioutil.ReadAll(file)
	if len(fileBytes) < 5 {
		fmt.Println("wrong! no valid data")
		return []dataPersist{}, nil
	}
	start := 0

	for i := 0; i < len(fileBytes)-4; i++ {
		if fileBytes[i] == '!' && fileBytes[i+1] == 'p' && fileBytes[i+2] == 'g' && fileBytes[i+3] == 'n' && fileBytes[i+4] == 'b' {
			var resultLine dataPersist
			var resultLineBytes []byte
			resultLineBytes = fileBytes[start:i]
			buffer := bytes.NewBuffer(resultLineBytes)
			dec := gob.NewDecoder(buffer)
			err1 := dec.Decode(&resultLine)
			if err1 != nil {
				fmt.Printf("error when Decoding!\n")
				return result, err1
			}
			//fmt.Println("resultLine: ",resultLine)
			result = append(result, resultLine)
			flag++
			start = i + 5
			i = i + 4
		}
	}

	err2 := file.Close()
	if err2 != nil {
		fmt.Printf("error when Close()!\n")
		return result, err2
	}
	return result, nil
}
