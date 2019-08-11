package main

import (
	"bufio"
	"fmt"
	"os"
)

type TimeAndData struct {
	Timestamp string
	Data      map[string]string
}

func createLogTxtGetTimestamps() {
	fmt.Println("input the timestamps between which you want to query")

	var timeStamp1, timeStamp2 string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("timestamp1: ")
	stdIn1, _, _ := reader.ReadLine()
	timeStamp1 = string(stdIn1)

	fmt.Println("timestamp2: ")
	stdIn2, _, _ := reader.ReadLine()
	timeStamp2 = string(stdIn2)

	var name string
	name = createTimestampName(timeStamp1) + "TO" + createTimestampName(timeStamp2) + ".txt"
	createLogTxt(name, timeStamp1, timeStamp2)
}

func createLogTxt(name, timeStamp1, timeStamp2 string) {
	result := rangeQuery(timeStamp1, timeStamp2)
	file, err1 := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("error when creating dataLog.txt!")
		panic(err1)
	}
	defer file.Close()

	fmt.Println("making the txt log...")

	for i, v := range result {
		_, _ = fmt.Fprintln(file, "This is the ", i, " dataRecord,", " this TIME is ", v.Timestamp)
		_, _ = fmt.Fprintln(file)
		for v1, v2 := range v.Data {
			_, _ = fmt.Fprintln(file, v.Timestamp, "   ", v1, " : ", v2)
		}
		_, _ = fmt.Fprintln(file)
		_, _ = fmt.Fprintln(file)
		_, _ = fmt.Fprintln(file)
	}

	fmt.Println("the txt log is ok for use now")
}

func createTimestampName(timeStamp string) string {
	var name string
	word1, indexEnd := getWord(timeStamp)
	word1 = makeTheCharTo_(word1)
	word2, _ := getWord(timeStamp[indexEnd:])
	word2 = makeTheCharTo_(word2)
	name = word1 + "-" + word2
	return name
}

func makeTheCharTo_(word string) string { //this function change the char not number to '_'
	wordBytes := []byte(word) //or the log file's name will be invalid because of chars like ':'
	for i := 0; i < len(word); i++ {
		if int(wordBytes[i])-int('0') < 0 || int(wordBytes[i])-int('9') > 0 {
			wordBytes[i] = '_'
		}
	}
	return string(wordBytes)
}
