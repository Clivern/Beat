// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

// Expect compare two values for testing
func Expect(t *testing.T, got, want interface{}) {
	t.Logf(`Comparing values %v, %v`, got, want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf(`ERROR! got %v, want %v`, got, want)
	}
}

// GetBaseDir returns the command line tool base dir
// Base dir idenfied if dirName found
// This function for testing purposes only
func GetBaseDir(dirName string) string {
	baseDir, _ := os.Getwd()
	cacheDir := fmt.Sprintf("%s/%s", baseDir, dirName)

	for {
		if fi, err := os.Stat(cacheDir); err == nil {
			if fi.Mode().IsDir() {
				return baseDir
			}
		}
		baseDir = filepath.Dir(baseDir)
		cacheDir = fmt.Sprintf("%s/%s", baseDir, dirName)
	}
}

// LoadConfigs load configs for testing purposes
func LoadConfigs(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	viper.SetConfigType("yaml")
	return viper.ReadConfig(bytes.NewBuffer([]byte(data)))
}
