/*
Copyright © 2022 RBG-TUM (Joscha Henningsen)
Licensed under the MIT License
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start interactive configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
