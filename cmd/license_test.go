// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"testing"

	"bitbucket.org/clivern/beat/pkg"
)

// TestLicenseCommand test cases
func TestLicenseCommand(t *testing.T) {
	// TestLicenseCommand
	t.Run("TestLicenseCommand", func(t *testing.T) {
		pkg.Expect(t, licenseHandler(), "MIT License, Copyright (c) 2020 Clivern")
	})
}
