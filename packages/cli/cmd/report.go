package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ReportCommand(rootCmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("usage: report <startdate> <enddate?>")
		return
	}
}
