package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"sync"
)

type dataNew struct {
	mu   sync.RWMutex
	data map[string]string
}

func main() {
	var dataNewest dataNew
	var urlGet string
	var metricsDisplaying []string
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

	if _, err := toml.DecodeFile("configuration.toml", &newConf); err != nil {
		fmt.Println("configurations loading error! ")
		fmt.Println(err)
		return
	} else {
		urlGet = newConf.Url
		getDura = newConf.GetDura
		go httpGet(urlGet, &getDura, &dataNewest, closeGetting)

		metricsDisplaying = newConf.MetricsShowing
		displayingDura = newConf.FreshDura
		metricsMeta = newConf.MetricMeta
		go print(metricsDisplaying, urlGet, &displayingDura, metricsMeta, closeDisplaying, &dataNewest)

		databaseInfo = newConf.DB
		go toDataBase(databaseInfo)

		persistOk = newConf.Persist
		persistDura = newConf.PersistDura
		go persisting(persistOk, &persistDura, &dataNewest, closePersisting)
	}

	controlling(&dataNewest, closeGetting, closeDisplaying, closePersisting, metricsMeta, metricsDisplaying, urlGet)

}
