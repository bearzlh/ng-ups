/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "upstream查询",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		query := map[string]string{}

		if ok, _ := cmd.Flags().GetBool("verbose"); ok {
			query["verbose"] = ""
		}
		filter, _ := cmd.Flags().GetString("filter")
		content := GetP(query, filter)
		fmt.Print(content)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.PersistentFlags().BoolP("verbose", "v", true, "详细信息")
	lsCmd.PersistentFlags().StringP("filter", "f", "", "过滤")
}
