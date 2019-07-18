package main

import "fmt"

/*
this file is used to parse the line and return two string.
*/

type Data struct {
	data      []byte
	lineStart int
	lineEnd   int
}

func dataToMap(bodyData []byte, dataNewest *dataNew) {
	dataNewest.mu.Lock()

	var dataTemp Data
	dataTemp.data = bodyData
	dataTemp.handleMap(dataNewest)

	dataNewest.mu.Unlock()
}

func (d *Data) handleMap(dataNewest *dataNew) {

	for d.lineEnd < len(d.data) && d.lineStart < len(d.data) {
		writeLine1, writeLine2 := d.nextLine()
		if writeLine1 != "" && writeLine1 != "#2" && len(writeLine1) > 0 && len(writeLine2) > 0 {
			(*dataNewest).data[writeLine1] = writeLine2
		}
	}
}

func (d *Data) nextLine() (result1, result2 string) { //get one line and it is parsed

	for i := d.lineStart; i < len(d.data); i++ {
		if (d.data)[i] == '\n' || (d.data)[i] == '\r' || i == len(d.data)-1 {
			//fmt.Println("here is a line !")
			d.lineEnd = i
			result1, result2 = ParseLine(d.lineStart, d.lineEnd, d.data)
			d.lineStart = i + 1
			break
		}
	}
	//fmt.Println("by nextLine() ",result1,result2,"d.lineStart and d.lineEnd: ",d.lineStart,d.lineEnd)
	return result1, result2
}

func ParseLine(lineStart, lineEnd int, body []byte) (line1, line2 string) {
	line := string(body[lineStart:lineEnd])

	//fmt.Println(string(line))

	if len(line) < 6 {
		fmt.Println("too short! the len of line is only ", len(line))
		panic("the line is " + line)
	}

	if line[0] == '#' { //"TYPE" OR "HELP"
		word, _ := getWord(line[1:])
		if word == "TYPE" {
			line1 = "#2"
			line2, _ = getWord(line[6:])
		} else if word == "HELP" {
			line1 = "#1"
			_, lengthBefore := getWord(line[6:])
			line2 = line[6+lengthBefore:]
		}

	} else {
		for i := 0; i < len(line); i++ {
			if line[i] == ' ' {
				line1 = string(line[0:i])
				line2 = string(line[i+1:])
			}
		}
	}
	return line1, line2
}

func getWord(line string) (word string, lengthTotal int) { //find the first single word
	start1 := 0
	start2 := 0
	flag := 0
	for i := start1; i < len(line); i++ {
		if line[i] != ' ' && flag == 0 { //the first character
			start1 = i
			flag = 1
		}
		if flag == 1 && (line[i] == ' ' || i == len(line)-1) {
			if line[i] == ' ' {
				start2 = i
			} else if i == len(line)-1 {
				start2 = i + 1
			}
			break
		}
	}
	word = string(line[start1:start2])
	return word, start2 //start2 is also the length, including the prefix ' '
}
