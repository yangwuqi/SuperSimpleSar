package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func httpGet(urlGet string, getDura *int, dataNewest *map[string]string, closeGetting chan bool) {
	makeDataToMap(dataNewest, urlGet)

	for {
		select { //Select is nice! I like it,
		case <-time.After(time.Duration((*getDura)*1000) * time.Millisecond):
			makeDataToMap(dataNewest, urlGet)

		case <-closeGetting:
			fmt.Println("shutdown ! ")
			return
		}
	}
}

func Get(url string) ([]byte, bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error!")
		return []byte{}, false
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("http read error!")
		return []byte{}, false
	}
	return body, true
}

func makeDataToMap(dataNewest *map[string]string, urlGet string) {
	if bodyData, ok := Get(urlGet); !ok {
		fmt.Println("http connection error!")
		return
	} else {
		dataToMap(bodyData, dataNewest)
	}
}
