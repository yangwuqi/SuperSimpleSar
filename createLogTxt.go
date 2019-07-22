package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

func createLogTxt(name string) {
	fmt.Println("the log txt file is creating... you can use snaptext to open it if it is too big")
	done := make(chan bool, 1)
	go doing(done)
	result, _ := load("data")
	file, err1 := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("error when creating dataLog.txt!")
		panic(err1)
	}
	defer file.Close()

	for i, v := range result {
		_, _ = fmt.Fprintln(file, "This is the ", i, " dataRecord,", " this TIME is ", v.Time)
		_, _ = fmt.Fprintln(file)
		for v1, v2 := range v.DataSaved {
			_, _ = fmt.Fprintln(file, v.Time, "   ", v1, " : ", v2)
		}
		_, _ = fmt.Fprintln(file)
		_, _ = fmt.Fprintln(file)
	}
	done <- true
	fmt.Println()
	fmt.Println("the log txt file is ok now")
}


func createLogTxtChosen(name string, metricsMeta string, metricsDisplaying []string, urlGet string){
	sort.Strings(metricsDisplaying)//for order
	fmt.Println("the log txt file for chosen metrics is creating... you can use snaptext to open it if it is too big")
	done := make(chan bool, 1)
	go doing(done)
	result, _ := load("data")
	file, err1 := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("error when creating dataLog.txt!")
		panic(err1)
	}
	defer file.Close()

	_, _ = fmt.Fprintln(file, "This is the txt log of the chosen metrics ", metricsDisplaying,", the information word is ", metricsMeta)
	_,_=fmt.Fprintln(file,)
	for i, v := range result {
		_, _ = fmt.Fprintln(file, "This is the ", i, " dataRecord,", " this TIME is ", v.Time)
		_, _ = fmt.Fprintln(file)
		for i:=0;i<len(metricsDisplaying);i++{//for order
			word1:=metricsDisplaying[i]
			if _,ok:=(v.DataSaved)[metricsDisplaying[i]];ok {
				word2 := (v.DataSaved)[metricsDisplaying[i]]
				_, _ = fmt.Fprintln(file, v.Time, "   ", word1, " : ", word2)
			}else{
				_, _ = fmt.Fprintln(file, v.Time, "   ", word1, " : ", "no value")
			}
		}
		_, _ = fmt.Fprintln(file)
		_, _ = fmt.Fprintln(file)
	}
	done <- true
	fmt.Println()
	fmt.Println("the log txt for chosen metrics is ok now")
}

func doing(done chan bool) {
	usingTime := 0
	for {
		select {
		case <-done:
			return
		case <-time.After(time.Duration(5000000000)):
			usingTime += 5
			fmt.Println("**having used ", usingTime, " s")
		}
	}
}
