/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	time2 "time"

	"github.com/spf13/cobra"
)

// weightIncreaseCmd represents the weightIncrease command
var weight = &cobra.Command{
	Use:   "weight",
	Short: "upstream权重变更",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		query := map[string]string{}
		weightInit, _ := cmd.Flags().GetInt("weight_init")
		weight, _ := cmd.Flags().GetInt("weight")
		time, _ := cmd.Flags().GetInt("time")
		span, _ := cmd.Flags().GetInt("span")

		server, _ := cmd.Flags().GetString("server")
		query["server"] = server

		// 启用，设置权重为0
		query["weight"] = fmt.Sprintf("%d", 0)
		query["up"] = ""
		GetP(query, "")
		content := GetP(map[string]string{"verbose": ""}, server)
		Debug(content)
		for i := weightInit; i < weight+span; i += span {
			if i > weight {
				i = weight
			}
			query["weight"] = fmt.Sprintf("%d", i)
			GetP(query, "")
			content := GetP(map[string]string{"verbose": ""}, server)
			Debug(content)
			time2.Sleep(time2.Duration(int64(time2.Second) * int64(time)))
		}
	},
}

func init() {
	rootCmd.AddCommand(weight)

	weight.PersistentFlags().StringP("server", "s", "", "eg:192.168.0.104:6004")
	weight.PersistentFlags().IntP("weight_init", "i", 1, "weight初值")
	weight.PersistentFlags().IntP("weight", "w", 10, "weight终值")
	weight.PersistentFlags().IntP("time", "t", 1, "time秒更新一次")
	weight.PersistentFlags().IntP("span", "p", 1, "每次增加的权重")
}
