// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"bitbucket.org/clivern/beat/pkg"
)

// TestHelpers test cases
func TestHelpers(t *testing.T) {
	// Get Base DIR and Cache DIR
	baseDir, _ := os.Getwd()
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")

	for {
		if DirExists(cacheDir) {
			break
		}
		baseDir = filepath.Dir(baseDir)
		cacheDir = fmt.Sprintf("%s/%s", baseDir, "cache")
	}

	// TestStringToIntFunc
	t.Run("TestStringToIntFunc", func(t *testing.T) {
		value, err := StringToInt("23")

		pkg.Expect(t, value, 23)
		pkg.Expect(t, err, nil)

		value, err = StringToInt("inv")

		pkg.Expect(t, value, 0)
		pkg.Expect(t, strings.Contains(err.Error(), `Unable to convert string value inv to int`), true)
	})

	// TestStringToFloat64Func
	t.Run("TestStringToFloat64Func", func(t *testing.T) {
		value, err := StringToFloat64("23.728613")
		testFloat := 23.728613

		pkg.Expect(t, value, testFloat)
		pkg.Expect(t, err, nil)

		value, err = StringToFloat64("inv")
		var emptyFloat float64

		pkg.Expect(t, value, emptyFloat)
		pkg.Expect(t, strings.Contains(err.Error(), `Unable to convert string value inv to float64`), true)
	})

	// TestStringToTimestampFunc
	t.Run("TestStringToTimestampFunc", func(t *testing.T) {
		value, err := StringToTimestamp("1405594992")

		pkg.Expect(t, value, time.Unix(1405594992, 0))
		pkg.Expect(t, err, nil)

		value, err = StringToTimestamp("inv")
		var emptyTimestamp time.Time

		pkg.Expect(t, value, emptyTimestamp)
		pkg.Expect(t, strings.Contains(err.Error(), `Unable to convert string value inv to timestamp`), true)
	})

	// TestFileExistsFunc
	t.Run("TestFileExistsFunc", func(t *testing.T) {
		pkg.Expect(t, FileExists("helpers.go"), true)
		pkg.Expect(t, FileExists("not_found.go"), false)
	})

	// TestPathExistsFunc
	t.Run("TestPathExistsFunc", func(t *testing.T) {
		pkg.Expect(t, PathExists("helpers.go"), true)
		pkg.Expect(t, PathExists("not_found.go"), false)
	})

	// TestDirExistsFunc
	t.Run("TestDirExistsFunc", func(t *testing.T) {
		pkg.Expect(t, DirExists(cacheDir), true)
		pkg.Expect(t, DirExists(fmt.Sprintf("%s/notFound", cacheDir)), false)
	})

	t.Run("TestDeleteFileFunc", func(t *testing.T) {
		// Create test file
		tmpFilePath := fmt.Sprintf("%s/helpers_test_file_01.txt", cacheDir)
		os.Create(tmpFilePath)
		pkg.Expect(t, DeleteFile(tmpFilePath), nil)
		pkg.Expect(t, DeleteFile(tmpFilePath) != nil, true)
	})

	t.Run("TestReadFileFunc", func(t *testing.T) {
		tmpFilePath := fmt.Sprintf("%s/helpers_test_file_02.txt", cacheDir)

		file, _ := os.Create(tmpFilePath)
		file.WriteString("Hello World")

		content, err := ReadFile(tmpFilePath)
		pkg.Expect(t, content, "Hello World")
		pkg.Expect(t, err, nil)

		nonExistentFile := fmt.Sprintf("%s/helpers_not_exist_file.txt", cacheDir)
		_, err = ReadFile(nonExistentFile)
		pkg.Expect(t, err != nil, true)
	})
}
