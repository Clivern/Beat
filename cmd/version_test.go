// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"testing"

	"github.com/franela/goblin"
)

// TestVersionCommand test cases
func TestVersionCommand(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("VersionCommand", func() {
		g.It("It should run and return the expected string", func() {
			g.Assert(versionHandler()).Equal("Current Beat Version dev Commit none, Built @unknown By unknown.")
		})
	})
}

// BenchmarkVersionCommand benchmark
func BenchmarkVersionCommand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		versionHandler()
	}
}
