package main

/*
func createLogTxt(name string) {
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
		_, _ = fmt.Fprintln(file)
	}
}
*/