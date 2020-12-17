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

	"github.com/franela/goblin"
)

// TestStringToIntFunc test cases
func TestStringToIntFunc(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("StringToInt", func() {
		g.It("It should satisfy all provided test cases", func() {
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
				value, err := StringToInt(tt.value)
				g.Assert(value).Equal(tt.wantValue)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})

	})
}

// BenchmarkStringToInt benchmark
func BenchmarkStringToInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringToInt("10")
	}
}

// TestStringToFloat64Func test cases
func TestStringToFloat64Func(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("StringToFloat64", func() {
		g.It("It should satisfy all provided test cases", func() {
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
				value, err := StringToFloat64(tt.value)
				g.Assert(value).Equal(tt.wantValue)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})

	})
}

// BenchmarkStringToFloat64 benchmark
func BenchmarkStringToFloat64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringToFloat64("10.23")
	}
}

// TestStringToTimestampFunc test cases
func TestStringToTimestampFunc(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("StringToTimestamp", func() {
		g.It("It should satisfy all provided test cases", func() {
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
				value, err := StringToTimestamp(tt.value)
				g.Assert(value).Equal(tt.wantValue)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})

	})
}

// BenchmarkStringToTimestamp benchmark
func BenchmarkStringToTimestamp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringToTimestamp("1405594992")
	}
}

// TestFileExistsFunc test cases
func TestFileExistsFunc(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("FileExists", func() {
		g.It("It should find the file and return true", func() {
			g.Assert(FileExists("helpers.go")).Equal(true)
		})

		g.It("It shouldn't find the file and return false", func() {
			g.Assert(FileExists("not_found.go")).Equal(false)
		})
	})
}

// BenchmarkFileExists benchmark
func BenchmarkFileExists(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FileExists("helpers.go")
	}
}

// TestDeleteFileFunc test cases
func TestDeleteFileFunc(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")
	tmpFilePath := fmt.Sprintf("%s/helpers_test_file_01.txt", cacheDir)
	os.Create(tmpFilePath)

	g := goblin.Goblin(t)

	g.Describe("DeleteFile", func() {
		g.It("It should delete the file and return no errors", func() {
			g.Assert(DeleteFile(tmpFilePath)).Equal(nil)
		})

		g.It("It should return an error since file got deleted before", func() {
			g.Assert(DeleteFile(tmpFilePath) != nil).Equal(true)
		})
	})
}

// BenchmarkDeleteFile benchmark
func BenchmarkDeleteFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DeleteFile("not_found_file.go")
	}
}

// TestReadFileFunc test cases
func TestReadFileFunc(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	cacheDir := fmt.Sprintf("%s/%s", baseDir, "cache")

	// Create a tmp file
	tmpFilePath := fmt.Sprintf("%s/helpers_test_file_02.txt", cacheDir)
	file, _ := os.Create(tmpFilePath)
	file.WriteString("Hello World")

	nonExistentFile := fmt.Sprintf("%s/helpers_not_exist_file.txt", cacheDir)

	g := goblin.Goblin(t)

	g.Describe("ReadFile", func() {
		g.It("It should read the file, return content and no errors", func() {
			content, err := ReadFile(tmpFilePath)

			g.Assert(content).Equal("Hello World")
			g.Assert(err).Equal(nil)
		})

		g.It("It should return an error since file not exist", func() {
			_, err := ReadFile(nonExistentFile)

			g.Assert(err != nil).Equal(true)
		})
	})
}

// BenchmarkReadFile benchmark
func BenchmarkReadFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadFile("helpers.go")
	}
}
