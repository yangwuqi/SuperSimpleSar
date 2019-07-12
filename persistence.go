package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"time"
)

type dataPersist struct {
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
			save("data", *dataNewest)

		case <-closePersisting:
			fmt.Println("the persisting process is shut down")
			return
		}
	}

}

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
