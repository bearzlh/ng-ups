/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"ng-ups/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "修改配置",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		zone, _ := cmd.Flags().GetString("zone")
		url, _ := cmd.Flags().GetString("url")
		debug, _ := cmd.Flags().GetString("debug")

		if zone != "" || url != "" || debug != "" {
			if zone != "" {
				config.Config.Set("zone", zone)
			}
			if url != "" {
				config.Config.Set("url", url)
			}
			if debug != "" {
				if debug == "true" {
					config.Config.Set("debug", true)
				} else {
					config.Config.Set("debug", false)
				}
			}
			err := config.Config.WriteConfig()
			if err != nil {
				log.Fatalf("更新失败:%v", err)
			}
		}

		conf, _ := json.Marshal(config.Config.AllSettings())
		fmt.Println(string(conf))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringP("zone", "z", "", "upstream zone config")
	configCmd.PersistentFlags().StringP("url", "u", "", "example:http://127.0.0.1:6000/ups")
	configCmd.PersistentFlags().StringP("debug", "d", "", "debug信息")
}
