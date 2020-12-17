// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"time"

	"bitbucket.org/clivern/beat/core/module"
	"bitbucket.org/clivern/beat/core/util"

	"github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DatasetFile var
var DatasetFile string

// OutputFile var
var OutputFile string

// Config var
var Config string

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate fare for a big set of rides",
	Run:   CalculateHandler,
}

// CalculateHandler runs the calculate command handler
func CalculateHandler(_ *cobra.Command, args []string) {
	spin := spinner.New(spinner.CharSets[27], 100*time.Millisecond)
	spin.Color("green")
	spin.Start()

	result, err := calculateHandler(args...)

	spin.Stop()

	// If command failed due to wrong command flags or non existent files,
	// stop execution
	if err != nil {
		panic(err)
	}

	fmt.Println(aurora.Green(result))
}

func calculateHandler(_ ...string) (string, error) {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("calculate command got called.")

	content, err := util.ReadFile(Config)

	if err != nil {
		return "", fmt.Errorf(
			"Error while loading config file %s: %s",
			Config,
			err.Error(),
		)
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(content)))

	if err != nil {
		return "", fmt.Errorf(
			"Error while loading config file content %s: %s",
			Config,
			err.Error(),
		)
	}

	log.Debug(fmt.Sprintf("Config file %s loaded successfully", Config))

	channel, err := module.GenerateData(DatasetFile)

	if err != nil {
		return "", fmt.Errorf(
			"Error while reading dataset file %s: %s",
			DatasetFile,
			err.Error(),
		)
	}

	outChannel := module.ProcessData(channel)

	err = module.StoreData(OutputFile, outChannel)

	if err != nil {
		return "", fmt.Errorf(
			"Error while storing date to file %s: %s",
			OutputFile,
			err.Error(),
		)
	}

	return "Ride data processed successfully!", nil
}

func init() {
	calculateCmd.Flags().StringVarP(
		&Config,
		"config_file",
		"c",
		"config.dist.yml",
		"Absolute path to config file (required)",
	)
	calculateCmd.Flags().StringVarP(
		&DatasetFile,
		"dataset_file",
		"i",
		"",
		"Absolute path to dataset CSV file (required)",
	)
	calculateCmd.Flags().StringVarP(
		&OutputFile,
		"output_file",
		"o",
		"",
		"Absolute path to output CSV file (required)",
	)
	calculateCmd.MarkFlagRequired("dataset_file")
	calculateCmd.MarkFlagRequired("output_file")
	rootCmd.AddCommand(calculateCmd)
}
