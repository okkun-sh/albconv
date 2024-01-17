package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "albconv",
	Short: "albconv is a CLI tool for converting AWS ALB access logs.",
	Long:  "albconv is a CLI tool that converts AWS ALB access logs, designed to be used effectively in combination with UNIX commands.",
	Run:   runRoot,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runRoot(cmd *cobra.Command, args []string) {
	entries, err := Run(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(entries)
}
