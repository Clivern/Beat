// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"strings"
	"testing"

	"bitbucket.org/clivern/beat/core/util"
	"bitbucket.org/clivern/beat/pkg"

	"github.com/franela/goblin"
)

// TestGenerateData test cases
func TestGenerateData(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("GenerateData", func() {
		g.It("It should fail since file is missing", func() {
			_, err := GenerateData(fmt.Sprintf("%s/not_found.csv", testDataDir))
			g.Assert(err != nil).Equal(true)
		})

		g.It("It should satisfy all provided test cases", func() {
			channel, err := GenerateData(fmt.Sprintf("%s/test_paths_01.csv", testDataDir))
			g.Assert(err).Equal(nil)

			var output []string

			for elem := range channel {
				output = append(output, elem)
			}

			// Validate the data sent to the channel & it equals the data that was on the file (cache & testdata)
			g.Assert(len(output)).Equal(10)
			g.Assert(output[0]).Equal("1,37.966660,23.728308,1405594957")
			g.Assert(output[1]).Equal("2,37.946545,23.754918,1405591065\n2,37.946545,23.754918,1405591073")
			g.Assert(output[2]).Equal("3,37.946545,23.754918,1405591084\n3,37.946413,23.754767,1405591094\n3,37.946260,23.754830,1405591103")
			g.Assert(output[3]).Equal("4,37.946032,23.755347,1405591112\n4,37.946190,23.755707,1405591121\n4,37.946298,23.756495,1405591132\n4,37.946398,23.758092,1405591142")
			g.Assert(output[4]).Equal("5,37.946417,23.759267,1405591151\n5,37.945638,23.758867,1405591160\n5,37.945638,23.758867,1405591161\n5,37.945310,23.758720,1405591173\n5,37.945045,23.758625,1405591181")
			g.Assert(output[5]).Equal("6,37.944860,23.758528,1405591191\n6,37.944530,23.758438,1405591201\n6,37.944370,23.758412,1405591211\n6,37.944365,23.758407,1405591222\n6,37.944365,23.758407,1405591232\n6,37.944365,23.758407,1405591233")
			g.Assert(output[6]).Equal("7,37.944365,23.758407,1405591243\n7,37.944365,23.758407,1405591253\n7,37.944440,23.758473,1405591263\n7,37.944440,23.758473,1405591273\n7,37.944440,23.758473,1405591283\n7,37.944440,23.758473,1405591293\n7,37.944440,23.758473,1405591303")
			g.Assert(output[7]).Equal("8,37.944440,23.758473,1405591313\n8,37.944360,23.758402,1405591323\n8,37.944360,23.758402,1405591334\n8,37.944360,23.758402,1405591343\n8,37.944360,23.758402,1405591354\n8,37.944360,23.758402,1405591363\n8,37.944253,23.758287,1405591373\n8,37.944253,23.758287,1405591383")
			g.Assert(output[8]).Equal("9,37.944253,23.758287,1405591394\n9,37.944253,23.758287,1405591404\n9,37.944122,23.758543,1405591414\n9,37.944028,23.758845,1405591424\n9,37.943680,23.759372,1405591434\n9,37.943667,23.759413,1405591444\n9,37.943883,23.758887,1405591455\n9,37.944130,23.758447,1405591464\n9,37.944563,23.758408,1405591474")
			g.Assert(output[9]).Equal("10,37.945335,23.758682,1405591484\n10,37.946275,23.759078,1405591494\n10,37.946490,23.758197,1405591504\n10,37.946472,23.757032,1405591514\n10,37.946410,23.756332,1405591525\n10,37.946610,23.755890,1405591534\n10,37.946832,23.755435,1405591553\n10,37.946408,23.754733,1405591554\n10,37.946613,23.753868,1405591566\n10,37.947072,23.752240,1405591577")
		})
	})
}

// TestStoreData test cases
func TestStoreData(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("StoreData", func() {
		g.It("It should fail since file is missing", func() {
			_, err := GenerateData(fmt.Sprintf("%s/not_found.csv", testDataDir))
			g.Assert(err != nil).Equal(true)
		})

		g.It("It should satisfy all provided test cases", func() {
			channel, err := GenerateData(fmt.Sprintf("%s/test_paths_01.csv", testDataDir))
			g.Assert(err).Equal(nil)

			err = StoreData(fmt.Sprintf("%s/store_data_test01.csv", cacheDir), channel)
			g.Assert(err).Equal(nil)

			// Validate written data using generate method
			channel, err = GenerateData(fmt.Sprintf("%s/store_data_test01.csv", cacheDir))
			g.Assert(err).Equal(nil)

			var output []string

			for elem := range channel {
				output = append(output, elem)
			}

			// Validate the data sent to the channel & it equals the data that was on the file (cache & testdata)
			g.Assert(len(output)).Equal(10)
			g.Assert(output[0]).Equal("1,37.966660,23.728308,1405594957")
			g.Assert(output[1]).Equal("2,37.946545,23.754918,1405591065\n2,37.946545,23.754918,1405591073")
			g.Assert(output[2]).Equal("3,37.946545,23.754918,1405591084\n3,37.946413,23.754767,1405591094\n3,37.946260,23.754830,1405591103")
			g.Assert(output[3]).Equal("4,37.946032,23.755347,1405591112\n4,37.946190,23.755707,1405591121\n4,37.946298,23.756495,1405591132\n4,37.946398,23.758092,1405591142")
			g.Assert(output[4]).Equal("5,37.946417,23.759267,1405591151\n5,37.945638,23.758867,1405591160\n5,37.945638,23.758867,1405591161\n5,37.945310,23.758720,1405591173\n5,37.945045,23.758625,1405591181")
			g.Assert(output[5]).Equal("6,37.944860,23.758528,1405591191\n6,37.944530,23.758438,1405591201\n6,37.944370,23.758412,1405591211\n6,37.944365,23.758407,1405591222\n6,37.944365,23.758407,1405591232\n6,37.944365,23.758407,1405591233")
			g.Assert(output[6]).Equal("7,37.944365,23.758407,1405591243\n7,37.944365,23.758407,1405591253\n7,37.944440,23.758473,1405591263\n7,37.944440,23.758473,1405591273\n7,37.944440,23.758473,1405591283\n7,37.944440,23.758473,1405591293\n7,37.944440,23.758473,1405591303")
			g.Assert(output[7]).Equal("8,37.944440,23.758473,1405591313\n8,37.944360,23.758402,1405591323\n8,37.944360,23.758402,1405591334\n8,37.944360,23.758402,1405591343\n8,37.944360,23.758402,1405591354\n8,37.944360,23.758402,1405591363\n8,37.944253,23.758287,1405591373\n8,37.944253,23.758287,1405591383")
			g.Assert(output[8]).Equal("9,37.944253,23.758287,1405591394\n9,37.944253,23.758287,1405591404\n9,37.944122,23.758543,1405591414\n9,37.944028,23.758845,1405591424\n9,37.943680,23.759372,1405591434\n9,37.943667,23.759413,1405591444\n9,37.943883,23.758887,1405591455\n9,37.944130,23.758447,1405591464\n9,37.944563,23.758408,1405591474")
			g.Assert(output[9]).Equal("10,37.945335,23.758682,1405591484\n10,37.946275,23.759078,1405591494\n10,37.946490,23.758197,1405591504\n10,37.946472,23.757032,1405591514\n10,37.946410,23.756332,1405591525\n10,37.946610,23.755890,1405591534\n10,37.946832,23.755435,1405591553\n10,37.946408,23.754733,1405591554\n10,37.946613,23.753868,1405591566\n10,37.947072,23.752240,1405591577")
		})
	})
}

// TestProcessData test cases
func TestProcessData(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("ProcessData", func() {
		g.It("It should satisfy all provided test cases", func() {
			channel, err := GenerateData(fmt.Sprintf("%s/test_paths_01.csv", testDataDir))
			g.Assert(err).Equal(nil)

			outChannel := ProcessData(channel)

			err = StoreData(fmt.Sprintf("%s/process_data_test01.csv", cacheDir), outChannel)
			g.Assert(err).Equal(nil)

			fileContent, err := util.ReadFile(fmt.Sprintf("%s/process_data_test01.csv", cacheDir))
			g.Assert(err).Equal(nil)

			g.Assert(strings.Contains(fileContent, "1,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "2,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "3,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "4,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "5,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "6,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "7,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "8,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "9,3.47")).Equal(true)
			g.Assert(strings.Contains(fileContent, "10,3.47")).Equal(true)
		})

		g.It("It should satisfy all provided test cases", func() {
			channel, err := GenerateData(fmt.Sprintf("%s/test_paths_02.csv", testDataDir))
			g.Assert(err).Equal(nil)

			outChannel := ProcessData(channel)

			err = StoreData(fmt.Sprintf("%s/process_data_test02.csv", cacheDir), outChannel)
			g.Assert(err).Equal(nil)

			fileContent, err := util.ReadFile(fmt.Sprintf("%s/process_data_test02.csv", cacheDir))
			g.Assert(err).Equal(nil)

			g.Assert(strings.Contains(fileContent, "2,58.30")).Equal(true)
		})
	})
}
