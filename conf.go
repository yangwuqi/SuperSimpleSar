package main

type Conf struct { //the comfiguration
	MetricMeta     string
	Url            string
	FreshDura      int
	GetDura        int
	Persist        bool
	PersistDura    int
	MetricsShowing []string
	MetricsAlert   []int
	DB             database
}

type database struct {
	Url             string
	DataBaseEnabled bool
	User            string
	Password        string
}
