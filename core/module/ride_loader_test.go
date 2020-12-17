// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"testing"

	"bitbucket.org/clivern/beat/core/model"

	"github.com/franela/goblin"
)

// TestCSVLoader test cases
func TestCSVLoader(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("CSVLoader", func() {

		g.It("CSVLoader should implement RideLoader", func() {
			var _ RideLoader = CSVLoader{}
		})

		g.It("It should satisfy all provided test cases", func() {

			var tests = []struct {
				data string

				wantRideID                   int
				wantCoordinatesCount         int
				wantFirstCoordinateLatitude  float64
				wantFirstCoordinateLongitude float64
				wantLastCoordinateLatitude   float64
				wantLastCoordinateLongitude  float64
				wantErrorNil                 bool
			}{
				// Valid data
				{"1,37.966660,23.728308,1405594957\n1,37.966627,23.728263,1405594966\n1,37.966625,23.728263,1405594974", 1, 3, 37.966660, 23.728308, 37.966625, 23.728263, true},

				// Invalid timestamp in data
				{"1,37.966660,23.728308,1405594957\n1,37.966627,23.728263,ss", 1, 1, 0, 0, 0, 0, false},

				// Invalid ride id in data
				{"1,37.966660,23.728308,1405594957\ner,37.966627,23.728263,1405594957", 1, 1, 0, 0, 0, 0, false},

				// Invalid ride id in data
				{"rt,37.966660,23.728308,1405594957\ner,37.966627,23.728263,1405594957", 0, 0, 0, 0, 0, 0, false},

				// Invalid latitude in data
				{"1,37.966660,23.728308,1405594957\n1,ji,23.728263,1405594957", 1, 1, 0, 0, 0, 0, false},

				// Invalid longitude in data
				{"1,37.966660,23.728308,1405594957\n1,37.966660,gh,1405594957", 1, 1, 0, 0, 0, 0, false},
			}

			for _, tt := range tests {
				ride := model.NewRide()
				loader := CSVLoader{}
				_, err := loader.Load(ride, tt.data)

				g.Assert(ride.ID).Equal(tt.wantRideID)
				g.Assert(len(ride.Coordinates)).Equal(tt.wantCoordinatesCount)

				if tt.wantErrorNil {
					// Verify the first element
					g.Assert(ride.Coordinates[0].Latitude).Equal(tt.wantFirstCoordinateLatitude)
					g.Assert(ride.Coordinates[0].Longitude).Equal(tt.wantFirstCoordinateLongitude)

					// Verify the last element
					g.Assert(ride.Coordinates[tt.wantCoordinatesCount-1].Latitude).Equal(tt.wantLastCoordinateLatitude)
					g.Assert(ride.Coordinates[tt.wantCoordinatesCount-1].Longitude).Equal(tt.wantLastCoordinateLongitude)
				}

				g.Assert(err == nil).Equal(tt.wantErrorNil)
			}
		})

	})
}

// BenchmarkCSVLoad benchmark
func BenchmarkCSVLoad(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ride := model.NewRide()
		loader := CSVLoader{}

		loader.Load(ride, "1,37.966660,23.728308,1405594957\n1,37.966627,23.728263,1405594966\n1,37.966625,23.728263,1405594974")
	}
}
