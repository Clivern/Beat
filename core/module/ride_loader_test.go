// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"testing"

	"bitbucket.org/clivern/beat/core/model"
	"bitbucket.org/clivern/beat/pkg"
)

// TestCSVLoader test cases
func TestCSVLoader(t *testing.T) {

	t.Run("TestIfCSVLoaderImplementsRideLoader", func(t *testing.T) {
		// Verify that *CSVLoader implement RideLoader
		var _ RideLoader = CSVLoader{}
	})

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

	// At each iteration, we will add a new coordinate and calculate the price
	for _, tt := range tests {
		t.Run("TestLoadMethod", func(t *testing.T) {
			ride := model.NewRide()
			loader := CSVLoader{}
			_, err := loader.Load(ride, tt.data)

			pkg.Expect(t, ride.ID, tt.wantRideID)
			pkg.Expect(t, len(ride.Coordinates), tt.wantCoordinatesCount)

			if tt.wantErrorNil {
				// Verify the first element
				pkg.Expect(t, ride.Coordinates[0].Latitude, tt.wantFirstCoordinateLatitude)
				pkg.Expect(t, ride.Coordinates[0].Longitude, tt.wantFirstCoordinateLongitude)

				// Verify the last element
				pkg.Expect(t, ride.Coordinates[tt.wantCoordinatesCount-1].Latitude, tt.wantLastCoordinateLatitude)
				pkg.Expect(t, ride.Coordinates[tt.wantCoordinatesCount-1].Longitude, tt.wantLastCoordinateLongitude)
			}

			pkg.Expect(t, err == nil, tt.wantErrorNil)
		})
	}
}
