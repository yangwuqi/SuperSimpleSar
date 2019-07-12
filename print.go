package main

import (
	"fmt"
	"time"
)

//

func print(metricsDisplaying []string, url string, displayingDura *int, metricsMeta string, closeDisplaying chan bool, dataNewest *map[string]string) {
	if len(metricsDisplaying) == 0 {
		fmt.Println("no metric chosen to displaying! ")
		return
	}
	var metricLengths []int
	for i := 0; i < len(metricsDisplaying); i++ {
		length := len(metricsDisplaying[i])
		if i > 0 {
			length += 3
		}
		metricLengths = append(metricLengths, length)
	}
	firstPrintSingle(metricsDisplaying, metricsMeta, url)

	for {
		select {
		case <-time.After(time.Duration((*displayingDura)*1000) * time.Millisecond):
			simplePrint(metricsDisplaying, dataNewest, url, metricLengths)
			fmt.Println()

		case <-closeDisplaying:
			fmt.Println("shutdown the dynamic output...")
			return
		}
	}
}

func firstPrintSingle(metrics []string, metricMeta string, url string) {
	fmt.Println()
	fmt.Printf("*********************************************************%v - %v*************************************************************\n", url, metricMeta)
	fmt.Println()
	//fmt.Printf("TIME %v   ", time.Now().Format("2006/1/2 15:04:05"))
	fmt.Printf("TIME                      ")
	for i := 0; i < len(metrics); i++ {
		fmt.Printf("%v   ", metrics[i])
	}
	fmt.Println()
}

func simplePrint(metricsDisplaying []string, dataNest *map[string]string, url string, metricLengths []int) {
	fmt.Printf("TIME %v   ", time.Now().Format("2006/1/2 15:04:05"))
	for i := 0; i < len(metricsDisplaying); i++ {
		printblank(metricLengths[i] - len((*dataNest)[metricsDisplaying[i]]))
		fmt.Printf("%v", (*dataNest)[metricsDisplaying[i]])
	}
}

func printblank(number int) {
	for i := 0; i < number; i++ {
		fmt.Printf(" ")
	}
}
