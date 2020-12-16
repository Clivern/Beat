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
)

// TestGenerateData test cases
func TestGenerateData(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	// TestGenerateDataMissingFile
	t.Run("TestGenerateDataMissingFile", func(t *testing.T) {
		_, err := GenerateData(fmt.Sprintf("%s/not_found.csv", testDataDir))
		pkg.Expect(t, err != nil, true)
	})

	// TestGenerateDataExistentFile
	t.Run("TestGenerateDataExistentFile", func(t *testing.T) {
		channel, err := GenerateData(fmt.Sprintf("%s/test_paths_01.csv", testDataDir))
		pkg.Expect(t, err, nil)

		var output []string

		for elem := range channel {
			output = append(output, elem)
		}

		// Validate the data sent to the channel & it equals the data that was on the file
		pkg.Expect(t, len(output), 10)
		pkg.Expect(t, output[0], "1,37.966660,23.728308,1405594957")
		pkg.Expect(t, output[1], "2,37.946545,23.754918,1405591065\n2,37.946545,23.754918,1405591073")
		pkg.Expect(t, output[2], "3,37.946545,23.754918,1405591084\n3,37.946413,23.754767,1405591094\n3,37.946260,23.754830,1405591103")
		pkg.Expect(t, output[3], "4,37.946032,23.755347,1405591112\n4,37.946190,23.755707,1405591121\n4,37.946298,23.756495,1405591132\n4,37.946398,23.758092,1405591142")
		pkg.Expect(t, output[4], "5,37.946417,23.759267,1405591151\n5,37.945638,23.758867,1405591160\n5,37.945638,23.758867,1405591161\n5,37.945310,23.758720,1405591173\n5,37.945045,23.758625,1405591181")
		pkg.Expect(t, output[5], "6,37.944860,23.758528,1405591191\n6,37.944530,23.758438,1405591201\n6,37.944370,23.758412,1405591211\n6,37.944365,23.758407,1405591222\n6,37.944365,23.758407,1405591232\n6,37.944365,23.758407,1405591233")
		pkg.Expect(t, output[6], "7,37.944365,23.758407,1405591243\n7,37.944365,23.758407,1405591253\n7,37.944440,23.758473,1405591263\n7,37.944440,23.758473,1405591273\n7,37.944440,23.758473,1405591283\n7,37.944440,23.758473,1405591293\n7,37.944440,23.758473,1405591303")
		pkg.Expect(t, output[7], "8,37.944440,23.758473,1405591313\n8,37.944360,23.758402,1405591323\n8,37.944360,23.758402,1405591334\n8,37.944360,23.758402,1405591343\n8,37.944360,23.758402,1405591354\n8,37.944360,23.758402,1405591363\n8,37.944253,23.758287,1405591373\n8,37.944253,23.758287,1405591383")
		pkg.Expect(t, output[8], "9,37.944253,23.758287,1405591394\n9,37.944253,23.758287,1405591404\n9,37.944122,23.758543,1405591414\n9,37.944028,23.758845,1405591424\n9,37.943680,23.759372,1405591434\n9,37.943667,23.759413,1405591444\n9,37.943883,23.758887,1405591455\n9,37.944130,23.758447,1405591464\n9,37.944563,23.758408,1405591474")
		pkg.Expect(t, output[9], "10,37.945335,23.758682,1405591484\n10,37.946275,23.759078,1405591494\n10,37.946490,23.758197,1405591504\n10,37.946472,23.757032,1405591514\n10,37.946410,23.756332,1405591525\n10,37.946610,23.755890,1405591534\n10,37.946832,23.755435,1405591553\n10,37.946408,23.754733,1405591554\n10,37.946613,23.753868,1405591566\n10,37.947072,23.752240,1405591577")
	})
}

// TestStoreData test cases
func TestStoreData(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	// TestStoreDataMissingFile
	t.Run("TestStoreDataMissingFile", func(t *testing.T) {
		_, err := GenerateData(fmt.Sprintf("%s/not_found.csv", testDataDir))
		pkg.Expect(t, err != nil, true)
	})

	// TestStoreDataExistentFile
	t.Run("TestStoreDataExistentFile", func(t *testing.T) {
		channel, err := GenerateData(fmt.Sprintf("%s/test_paths_01.csv", testDataDir))
		pkg.Expect(t, err, nil)

		err = StoreData(fmt.Sprintf("%s/store_data_test01.csv", cacheDir), channel)
		pkg.Expect(t, err, nil)

		// Validate written data using generate method
		channel, err = GenerateData(fmt.Sprintf("%s/store_data_test01.csv", cacheDir))
		pkg.Expect(t, err, nil)

		var output []string

		for elem := range channel {
			output = append(output, elem)
		}

		// Validate the data sent to the channel & it equals the data that was on the file (cache & testdata)
		pkg.Expect(t, len(output), 10)
		pkg.Expect(t, output[0], "1,37.966660,23.728308,1405594957")
		pkg.Expect(t, output[1], "2,37.946545,23.754918,1405591065\n2,37.946545,23.754918,1405591073")
		pkg.Expect(t, output[2], "3,37.946545,23.754918,1405591084\n3,37.946413,23.754767,1405591094\n3,37.946260,23.754830,1405591103")
		pkg.Expect(t, output[3], "4,37.946032,23.755347,1405591112\n4,37.946190,23.755707,1405591121\n4,37.946298,23.756495,1405591132\n4,37.946398,23.758092,1405591142")
		pkg.Expect(t, output[4], "5,37.946417,23.759267,1405591151\n5,37.945638,23.758867,1405591160\n5,37.945638,23.758867,1405591161\n5,37.945310,23.758720,1405591173\n5,37.945045,23.758625,1405591181")
		pkg.Expect(t, output[5], "6,37.944860,23.758528,1405591191\n6,37.944530,23.758438,1405591201\n6,37.944370,23.758412,1405591211\n6,37.944365,23.758407,1405591222\n6,37.944365,23.758407,1405591232\n6,37.944365,23.758407,1405591233")
		pkg.Expect(t, output[6], "7,37.944365,23.758407,1405591243\n7,37.944365,23.758407,1405591253\n7,37.944440,23.758473,1405591263\n7,37.944440,23.758473,1405591273\n7,37.944440,23.758473,1405591283\n7,37.944440,23.758473,1405591293\n7,37.944440,23.758473,1405591303")
		pkg.Expect(t, output[7], "8,37.944440,23.758473,1405591313\n8,37.944360,23.758402,1405591323\n8,37.944360,23.758402,1405591334\n8,37.944360,23.758402,1405591343\n8,37.944360,23.758402,1405591354\n8,37.944360,23.758402,1405591363\n8,37.944253,23.758287,1405591373\n8,37.944253,23.758287,1405591383")
		pkg.Expect(t, output[8], "9,37.944253,23.758287,1405591394\n9,37.944253,23.758287,1405591404\n9,37.944122,23.758543,1405591414\n9,37.944028,23.758845,1405591424\n9,37.943680,23.759372,1405591434\n9,37.943667,23.759413,1405591444\n9,37.943883,23.758887,1405591455\n9,37.944130,23.758447,1405591464\n9,37.944563,23.758408,1405591474")
		pkg.Expect(t, output[9], "10,37.945335,23.758682,1405591484\n10,37.946275,23.759078,1405591494\n10,37.946490,23.758197,1405591504\n10,37.946472,23.757032,1405591514\n10,37.946410,23.756332,1405591525\n10,37.946610,23.755890,1405591534\n10,37.946832,23.755435,1405591553\n10,37.946408,23.754733,1405591554\n10,37.946613,23.753868,1405591566\n10,37.947072,23.752240,1405591577")
	})
}

// TestProcessData test cases
func TestProcessData(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	testDataDir := fmt.Sprintf("%s/%s", baseDir, "testdata")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	// TestProcessData
	t.Run("TestProcessData", func(t *testing.T) {
		channel, err := GenerateData(fmt.Sprintf("%s/test_paths_01.csv", testDataDir))
		pkg.Expect(t, err, nil)

		outChannel := ProcessData(channel)

		err = StoreData(fmt.Sprintf("%s/process_data_test01.csv", cacheDir), outChannel)
		pkg.Expect(t, err, nil)

		fileContent, err := util.ReadFile(fmt.Sprintf("%s/process_data_test01.csv", cacheDir))
		pkg.Expect(t, err, nil)

		pkg.Expect(t, strings.Contains(fileContent, "1,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "2,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "3,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "4,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "5,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "6,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "7,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "8,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "9,3.47"), true)
		pkg.Expect(t, strings.Contains(fileContent, "10,3.47"), true)
	})

	// TestProcessData
	t.Run("TestProcessData", func(t *testing.T) {
		channel, err := GenerateData(fmt.Sprintf("%s/test_paths_02.csv", testDataDir))
		pkg.Expect(t, err, nil)

		outChannel := ProcessData(channel)

		err = StoreData(fmt.Sprintf("%s/process_data_test02.csv", cacheDir), outChannel)
		pkg.Expect(t, err, nil)

		fileContent, err := util.ReadFile(fmt.Sprintf("%s/process_data_test02.csv", cacheDir))
		pkg.Expect(t, err, nil)

		pkg.Expect(t, strings.Contains(fileContent, "2,58.30"), true)
	})
}
