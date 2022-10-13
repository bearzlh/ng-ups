/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "upstream更新",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		weight, _ := cmd.Flags().GetInt("weight")
		maxFails, _ := cmd.Flags().GetInt("max_fails")
		failTimeout, _ := cmd.Flags().GetInt("fail_timeout")
		down, _ := cmd.Flags().GetBool("down")
		up, _ := cmd.Flags().GetBool("up")
		remove, _ := cmd.Flags().GetBool("remove")

		query := map[string]string{}
		if weight+maxFails+failTimeout > -3 {
			query = map[string]string{
				"server": server,
				"update": "",
			}
			if weight >= 0 {
				query["weight"] = strconv.Itoa(weight)
			}
			if maxFails >= 0 {
				query["max_fails"] = strconv.Itoa(maxFails)
			}
			if failTimeout >= 0 {
				query["fail_timeout"] = strconv.Itoa(failTimeout)
			}
			GetP(query, server)
			fmt.Println(GetServer(server))
		}
		query = map[string]string{}
		if down {
			query["down"] = ""
		} else if up {
			query["up"] = ""
		} else if remove {
			query["remove"] = ""
		}
		if len(query) > 0 {
			query["server"] = server
			GetP(query, server)
			if remove {
				fmt.Println(GetServer(""))
			} else {
				fmt.Println(GetServer(server))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.PersistentFlags().StringP("server", "s", "", "eg:192.168.0.104:6004")
	updateCmd.PersistentFlags().IntP("weight", "w", -1, "weight")
	updateCmd.PersistentFlags().IntP("max_fails", "f", -1, "max_fails")
	updateCmd.PersistentFlags().IntP("fail_timeout", "t", -1, "fail_timeout")
	updateCmd.PersistentFlags().BoolP("down", "d", false, "停用")
	updateCmd.PersistentFlags().BoolP("up", "u", false, "启用")
	updateCmd.PersistentFlags().BoolP("remove", "r", false, "删除")
}
