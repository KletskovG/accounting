package main

import "github.com/kletskovg/accounting/packages/cli/cmd"

const DateLayout = "2006-01-02"

func main() {
	cmd.Execute()
	// for i := 1; i < 30; i++ {
	// 	var dayStr string
	// 	if i < 10 {
	// 		dayStr = "0" + strconv.Itoa(i)
	// 	} else {
	// 		dayStr = strconv.Itoa(i)
	// 	}

	// 	date := "2023-12-" + dayStr
	// 	db.InsertTransaction(date, 1000, "test", "test")
	// }
}
