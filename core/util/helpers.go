// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// StringToInt converts a string to integer
func StringToInt(value string) (int, error) {
	var result int
	var err error

	result, err = strconv.Atoi(value)

	if err != nil {
		return result, fmt.Errorf(
			"Unable to convert string value %s to int: %s",
			value,
			err.Error(),
		)
	}

	return result, nil
}

// StringToFloat64 converts a string to float64
func StringToFloat64(value string) (float64, error) {
	var result float64
	var err error

	result, err = strconv.ParseFloat(value, 64)

	if err != nil {
		return result, fmt.Errorf(
			"Unable to convert string value %s to float64: %s",
			value,
			err.Error(),
		)
	}

	return result, nil
}

// StringToTimestamp converts a string to timestamp
func StringToTimestamp(value string) (time.Time, error) {
	var result time.Time

	intValue, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return result, fmt.Errorf(
			"Unable to convert string value %s to timestamp: %s",
			value,
			err.Error(),
		)
	}

	result = time.Unix(intValue, 0)

	return result, nil
}

// FileExists reports whether the named file exists
func FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}

	return false
}

// PathExists reports whether the path exists
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// DirExists reports whether the dir exists
func DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

// DeleteFile deletes a file
func DeleteFile(path string) error {
	return os.Remove(path)
}

// ReadFile get the file content
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
