// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ultiledger/go-ultiledger/api"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// read in config file
		if cfgFile == "" {
			log.Fatal(errors.New("config file not provided"))
		}
		v := viper.New()
		v.SetConfigFile(cfgFile)
		if err := v.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		// bootstrap a new ULTNode
		n := api.NewULTNode(v)
		n.Start()
	},
}

var cfgFile string

func init() {
	startCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Help message for toggle")
	startCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(startCmd)
}