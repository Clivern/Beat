// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"
	"testing"

	"bitbucket.org/clivern/beat/core/util"
	"bitbucket.org/clivern/beat/pkg"

	"github.com/franela/goblin"
)

// TestCalculateCommand test cases
func TestCalculateCommand(t *testing.T) {
	g := goblin.Goblin(t)

	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	DatasetFile = fmt.Sprintf("%s/test_paths_02.csv", testDataDir)
	OutputFile = fmt.Sprintf("%s/cache/calculate_command_test_01.csv", baseDir)
	Config = fmt.Sprintf("%s/config.dist.yml", baseDir)

	g.Describe("CalculateCommand", func() {
		g.It("It should run and calculate the ride fare inside cache/test_paths_02.csv file", func() {
			// Run command
			result, err := calculateHandler()

			g.Assert(err).Equal(nil)
			g.Assert(result).Equal("Ride data processed successfully!")

			// Validate command output
			fileContent, err := util.ReadFile(OutputFile)
			g.Assert(err).Equal(nil)
			g.Assert(strings.Contains(fileContent, "2,58.30")).Equal(true)
		})

		g.It("It should fail since dataset file doesn't exist", func() {
			// Override with non existent dataset file
			DatasetFile = fmt.Sprintf("%s/not_found_test_paths_02.csv", testDataDir)

			// Run command
			result, err := calculateHandler()

			g.Assert(err != nil).Equal(true)
			g.Assert(result).Equal("")
		})

		g.It("It should fail since config file doesn't exist", func() {
			// Add existent dataset file
			DatasetFile = fmt.Sprintf("%s/test_paths_02.csv", testDataDir)

			// Override with non existent config file
			Config = fmt.Sprintf("%s/not_found_config.dist.yml", baseDir)

			// Run command
			result, err := calculateHandler()

			g.Assert(err != nil).Equal(true)
			g.Assert(result).Equal("")
		})
	})
}

// BenchmarkCalculateCommand benchmark
func BenchmarkCalculateCommand(b *testing.B) {
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	DatasetFile = fmt.Sprintf("%s/test_paths_02.csv", testDataDir)
	OutputFile = fmt.Sprintf("%s/cache/calculate_command_bench_01.csv", baseDir)
	Config = fmt.Sprintf("%s/config.dist.yml", baseDir)

	for n := 0; n < b.N; n++ {
		calculateHandler()
	}
}
