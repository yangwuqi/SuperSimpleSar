package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func controlling(dataNewest *dataNew, closeGetting chan bool, closeDisplaying chan bool, closePersisting chan bool) {
	for {
		reader := bufio.NewReader(os.Stdin)
		stdIn, _, _ := reader.ReadLine()
		inputLine := string(stdIn)
		startIndex := 0
		word, length := getWord(inputLine[startIndex:])
		for word != "" {
			switch word {
			case "-q": //quit all
				_ = defaultBoltDB.Close()
				fmt.Println("shutdown all!")
				return
			//case "-save":
			//save("data", dataNewest)
			case "-s":
				closeDisplaying <- true
			case "-cp":
				closePersisting <- true
			case "-log":
				go createLogTxtGetTimestamps()
				time.Sleep(time.Second * 15)
			case "-la":
				go loadAll()
			case "-tl": //show timestamps list, timestamp is the key in boltDB
				go showTimestamps()
			case "-rq": //range query, timestamp to timestamp
				go rangeQueryPrint()
				time.Sleep(time.Second * 15)
			case "-as":
				go alertInfoSumAllQuery()
			case "-al":
				go createLogTxtAll("all.data")
			}
			startIndex += length
			if startIndex >= len(inputLine) {
				break
			}
			word, length = getWord(inputLine[startIndex:])
		}
	}
}
