// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// Version buildinfo item
	Version = "dev"
	// Commit buildinfo item
	Commit = "none"
	// Date buildinfo item
	Date = "unknown"
	// BuiltBy buildinfo item
	BuiltBy = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run:   VersionHandler,
}

// VersionHandler runs the version command handler
func VersionHandler(_ *cobra.Command, args []string) {
	fmt.Println(versionHandler(args...))
}

func versionHandler(_ ...string) string {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("Version command got called.")

	return fmt.Sprintf(
		`Current Beat Version %v Commit %v, Built @%v By %v.`,
		Version,
		Commit,
		Date,
		BuiltBy,
	)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
