// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Print the license",
	Run:   LicenseHandler,
}

// LicenseHandler runs the license command handler
func LicenseHandler(_ *cobra.Command, args []string) {
	fmt.Println(licenseHandler(args...))
}

func licenseHandler(_ ...string) string {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("License command got called.")

	return "MIT License, Copyright (c) 2020 Clivern"
}

func init() {
	rootCmd.AddCommand(licenseCmd)
}
