/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "버전 출력",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.3")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
