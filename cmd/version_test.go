// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"testing"

	"bitbucket.org/clivern/beat/pkg"
)

// TestVersionCommand test cases
func TestVersionCommand(t *testing.T) {
	// TestVersionCommand
	t.Run("TestVersionCommand", func(t *testing.T) {
		pkg.Expect(t, versionHandler(), "Current Beat Version dev Commit none, Built @unknown By unknown.")
	})
}
