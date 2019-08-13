package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/boltdb/bolt"
	"sync"
	"time"
)

var (
	defaultBoltDB   *bolt.DB
	defaultBucket   string
	alertInfoBucket string
)

type dataNew struct {
	mu        sync.RWMutex
	timeStamp string
	data      map[string]string
}

func main() {
	var dataNewest dataNew
	var urlGet string
	var metricsDisplaying []string
	var metricsAlertLine []string
	var databaseInfo database
	var getDura int
	var displayingDura int
	var metricsMeta string
	var persistOk bool
	var persistDura int
	var newConf Conf

	mapInitial := make(map[string]string)
	dataNewest.mu.Lock()
	dataNewest.data = mapInitial
	dataNewest.mu.Unlock()
	closeGetting := make(chan bool, 1)
	closeDisplaying := make(chan bool, 1)
	closePersisting := make(chan bool, 1)

	db, err := openBoltDB("defaultBoltDB")
	if err != nil {
		fmt.Println("cannot open the default BoltDB!")
	}
	defaultBoltDB = db
	err1 := createDbBucket(defaultBoltDB, "defaultBucket")
	if err1 != nil {
		fmt.Println("cannot open the default BoltDB Bucket!")
	}
	defaultBucket = "defaultBucket"
	err2 := createDbBucket(defaultBoltDB, "alertInfoBucket")
	if err2 != nil {
		fmt.Println("cannot open the alertInfo BoltDB Bucket!")
	}
	alertInfoBucket = "alertInfoBucket"

	if _, err := toml.DecodeFile("configuration.toml", &newConf); err != nil {
		fmt.Println("configurations loading error! ")
		fmt.Println(err)
		return
	} else {
		urlGet = newConf.Url
		getDura = newConf.GetDura
		go httpGet(urlGet, &getDura, &dataNewest, closeGetting)

		metricsDisplaying = newConf.MetricsShowing
		metricsAlertLine = newConf.MetricsAlertLine
		displayingDura = newConf.FreshDura
		metricsMeta = newConf.MetricMeta
		timeStart := time.Now().UnixNano()
		for len(dataNewest.data) == 0 {
			if time.Now().UnixNano()-timeStart > 10000000000 {
				fmt.Println("time out, quit")
				_ = defaultBoltDB.Close()
				fmt.Println("shutdown all!")
				return
			}
		}
		go print(metricsDisplaying, urlGet, &displayingDura, metricsMeta, closeDisplaying, &dataNewest, metricsAlertLine)

		databaseInfo = newConf.DB
		go toDataBase(databaseInfo)

		persistOk = newConf.Persist
		persistDura = newConf.PersistDura
		go persisting(persistOk, &persistDura, &dataNewest, closePersisting)
	}

	controlling(&dataNewest, closeGetting, closeDisplaying, closePersisting)

}
