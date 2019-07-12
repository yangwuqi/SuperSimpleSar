package main

import (
	"bufio"
	"fmt"
	"os"
)

func controlling(dataNewest *map[string]string, closeGetting chan bool, closeDisplaying chan bool, closePersisting chan bool) {
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
				save("data", *dataNewest)
			case "-s":
				closeDisplaying <- true
			case "-cp":
				closePersisting <- true
			}
			startIndex += length
			if startIndex >= len(inputLine) {
				break
			}
			word, length = getWord(inputLine[startIndex:])
		}
	}
}
