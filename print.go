package main

import (
	"fmt"
	"time"
)

//

func print(metricsDisplaying []string, url string, displayingDura *int, metricsMeta string, closeDisplaying chan bool, dataNewest *dataNew, metricsAlertLine []string) {
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
			simplePrint(metricsDisplaying, dataNewest, url, metricLengths, metricsAlertLine)
			fmt.Println()

		case <-closeDisplaying:
			fmt.Println("shutdown the dynamic output...")
			return
		}
	}
}

func firstPrintSingle(metrics []string, metricMeta string, url string) {
	fmt.Println()
	fmt.Printf("*********************************************************%v - %v****************************************************************\n", url, metricMeta)
	fmt.Println()
	//fmt.Printf("TIME %v   ", time.Now().Format("2006/1/2 15:04:05"))
	fmt.Printf("TIME                      ")
	for i := 0; i < len(metrics); i++ {
		fmt.Printf("%v   ", metrics[i])
	}
	fmt.Println()
}

func simplePrint(metricsDisplaying []string, dataNewest *dataNew, url string, metricLengths []int, metricsAlertLine []string) {
	dataNewest.mu.RLock()
	defer dataNewest.mu.RUnlock()

	var alertInfoNodeData []int
	timeStamp := (*dataNewest).timeStamp
	fmt.Printf("TIME %v   ", timeStamp)

	for i := 0; i < len(metricsDisplaying); i++ {
		printblank(metricLengths[i] - len((*dataNewest).data[metricsDisplaying[i]]))
		fmt.Printf("%v", (*dataNewest).data[metricsDisplaying[i]])
		alertInfoNodeData = append(alertInfoNodeData, 0)
		if metricsAlertLine[i] < (*dataNewest).data[metricsDisplaying[i]] {
			alertInfoNodeData[i] = 1
		}
	}
	alertInfoHandle(alertInfoNodeData, timeStamp)
}

func printblank(number int) {
	for i := 0; i < number; i++ {
		fmt.Printf(" ")
	}
}
