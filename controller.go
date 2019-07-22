package main

import (
	"bufio"
	"fmt"
	"os"
)

func controlling(dataNewest *dataNew, closeGetting chan bool, closeDisplaying chan bool, closePersisting chan bool, metricsMeta string, metricsDisplaying []string, urlGet string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		stdIn, _, _ := reader.ReadLine()
		inputLine := string(stdIn)
		startIndex := 0
		word, length := getWord(inputLine[startIndex:])
		for word != "" {
			switch word {
			case "-q":
				fmt.Println("shutdown all!")
				return
			case "-save":
				save("data", dataNewest)
			case "-sd":
				closeDisplaying <- true
			case "-cp":
				closePersisting <- true
			case "-log":
				createLogTxt("dataLog.txt")
			case "-logC"://create logTxt for the chosen metrics
				createLogTxtChosen("dataLogChosen.txt",metricsMeta, metricsDisplaying, urlGet)
			}
			startIndex += length
			if startIndex >= len(inputLine) {
				break
			}
			word, length = getWord(inputLine[startIndex:])
		}
	}
}
