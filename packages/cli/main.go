package main

import "github.com/kletskovg/accounting/packages/cli/cmd"

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
