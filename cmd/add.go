/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加upstream",
	Long:  `添加upstream`,
	Run: func(cmd *cobra.Command, args []string) {

		query := map[string]string{}

		server, _ := cmd.Flags().GetString("server")
		weight, _ := cmd.Flags().GetInt("weight")
		maxFails, _ := cmd.Flags().GetInt("max_fails")
		failTimeout, _ := cmd.Flags().GetInt("fail_timeout")
		query["server"] = server
		query["weight"] = strconv.Itoa(weight)
		query["max_fails"] = strconv.Itoa(maxFails)
		query["fail_timeout"] = strconv.Itoa(failTimeout)
		query["add"] = ""
		GetP(query, "")
		content := GetServer(server)
		fmt.Println(content)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringP("server", "s", "", "eg:192.168.0.104:6004")
	addCmd.PersistentFlags().IntP("weight", "w", 0, "weight")
	addCmd.PersistentFlags().IntP("max_fails", "f", 1, "max_fails")
	addCmd.PersistentFlags().IntP("fail_timeout", "t", 10, "fail_timeout")
}
