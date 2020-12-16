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
)

// TestCalculateCommand test cases
func TestCalculateCommand(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	// Override command args
	DatasetFile = fmt.Sprintf("%s/test_paths_02.csv", testDataDir)
	OutputFile = fmt.Sprintf("%s/cache/calculate_command_test_01.csv", baseDir)
	Config = fmt.Sprintf("%s/config.dist.yml", baseDir)

	// TestCalculateCommand
	t.Run("TestCalculateCommand", func(t *testing.T) {
		// Run command
		calculateHandler()

		// Validate command output
		fileContent, err := util.ReadFile(OutputFile)
		pkg.Expect(t, err, nil)
		pkg.Expect(t, strings.Contains(fileContent, "2,58.30"), true)
	})
}
