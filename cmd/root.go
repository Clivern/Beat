// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Verbose var
var Verbose bool

// Config var
var Config string

var rootCmd = &cobra.Command{
	Use: "beat",
	Short: `A Command Line Tool to do Fare Estimation for a big set of rides.

If you have any suggestions, bug reports, or annoyances please report
them to our issue tracker at <https://bitbucket.org/clivern/beat>`,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(
		&Config,
		"config",
		"c",
		"",
		"config file",
	)
}

// Execute runs cmd tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
