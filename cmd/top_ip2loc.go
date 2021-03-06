/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/cloudsark/loganalyzer/logalizer"
	"github.com/spf13/cobra"
)

// ip2locCmd represents the ip2loc command
var ip2locCmd = &cobra.Command{
	Use:   "ip2loc",
	Short: "Print Top 10 IP addresses accessing your web server with their location",
	Long: `Find Top 10 IP Addresses Accessing Your web server with their geo locations,
	it helps you to quicly identify abuse`,
	Run: func(cmd *cobra.Command, args []string) {
		logalizer.TopIP2LocCmd(&filename)
	},
}

func init() {
	topCmd.AddCommand(ip2locCmd)
}
