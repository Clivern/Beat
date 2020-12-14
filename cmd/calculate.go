// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate fare for a big set of rides",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		spin := spinner.New(spinner.CharSets[26], 100*time.Millisecond)
		spin.Color("green")
		spin.Start()

		log.Debug("calculate command got called.")

		spin.Stop()

		fmt.Println(aurora.Green("Fare got calculated for the provided sheet!"))
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)
}
