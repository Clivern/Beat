// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"os"
	"testing"
	"time"

	"bitbucket.org/clivern/beat/pkg"
)

// TestStringToIntFunc test cases
func TestStringToIntFunc(t *testing.T) {
	var tests = []struct {
		value     string
		wantValue int
		wantError bool
	}{
		{"23", 23, false},
		{"237888", 237888, false},
		{"f456666", 0, true},
		{"456666hyy", 0, true},
		{"df", 0, true},
		{"344552", 344552, false},
	}

	for _, tt := range tests {
		t.Run("TestStringToIntFunc", func(t *testing.T) {
			value, err := StringToInt(tt.value)
			pkg.Expect(t, value, tt.wantValue)
			pkg.Expect(t, err != nil, tt.wantError)
		})
	}
}

// TestStringToFloat64Func test cases
func TestStringToFloat64Func(t *testing.T) {
	var tests = []struct {
		value     string
		wantValue float64
		wantError bool
	}{
		{"23.728613", 23.728613, false},
		{"fr", 0, true},
		{"f456666", 0, true},
		{"456666hyy", 0, true},
		{"45666.6hyy", 0, true},
		{"344552", 344552, false},
	}

	for _, tt := range tests {
		t.Run("TestStringToFloat64Func", func(t *testing.T) {
			value, err := StringToFloat64(tt.value)
			pkg.Expect(t, value, tt.wantValue)
			pkg.Expect(t, err != nil, tt.wantError)
		})
	}
}

// TestStringToTimestampFunc test cases
func TestStringToTimestampFunc(t *testing.T) {
	var tests = []struct {
		value     string
		wantValue time.Time
		wantError bool
	}{
		{"1405594992", time.Unix(1405594992, 0), false},
		{value: "fr", wantError: true},
		{value: "g6738183", wantError: true},
		{value: "1405594992hy", wantError: true},
		{value: "14055.94992hy", wantError: true},
		{"344552", time.Unix(344552, 0), false},
	}

	for _, tt := range tests {
		t.Run("TestStringToTimestampFunc", func(t *testing.T) {
			value, err := StringToTimestamp(tt.value)
			pkg.Expect(t, value, tt.wantValue)
			pkg.Expect(t, err != nil, tt.wantError)
		})
	}
}

// TestFileExistsFunc test cases
func TestFileExistsFunc(t *testing.T) {
	// TestFileExistsFunc
	t.Run("TestFileExistsFunc", func(t *testing.T) {
		pkg.Expect(t, FileExists("helpers.go"), true)
		pkg.Expect(t, FileExists("not_found.go"), false)
	})
}

// TestDeleteFileFunc test cases
func TestDeleteFileFunc(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")

	t.Run("TestDeleteFileFunc", func(t *testing.T) {
		// Create test file
		tmpFilePath := fmt.Sprintf("%s/helpers_test_file_01.txt", cacheDir)
		os.Create(tmpFilePath)
		pkg.Expect(t, DeleteFile(tmpFilePath), nil)
		pkg.Expect(t, DeleteFile(tmpFilePath) != nil, true)
	})
}

// TestReadFileFunc test cases
func TestReadFileFunc(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")

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
