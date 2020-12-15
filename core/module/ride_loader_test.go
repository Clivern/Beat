// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"testing"
	"time"

	"bitbucket.org/clivern/beat/core/model"
	"bitbucket.org/clivern/beat/pkg"
)

// TestCSVLoader test cases
func TestCSVLoader(t *testing.T) {
	// TestLoadMethod
	t.Run("TestLoadMethod", func(t *testing.T) {
		// Verify that *CSVLoader implement RideLoader
		var _ RideLoader = CSVLoader{}

		ride := model.NewRide()
		loader := CSVLoader{}

		data := `1,37.966660,23.728308,1405594957
			1,37.966627,23.728263,1405594966
			1,37.966625,23.728263,1405594974
			1,37.966613,23.728375,1405594984
			1,37.966203,23.728597,1405594992
			1,37.966195,23.728613,1405595001
			1,37.966195,23.728613,1405595009
			1,37.966195,23.728613,1405595017
			1,37.966195,23.728613,1405595026
			1,37.966195,23.728613,1405595034
			1,37.966195,23.728613,1405595043
			1,37.966180,23.728662,1405595051
			1,37.965937,23.728440,1405595059
			1,37.965377,23.727717,1405595068
			1,37.965000,23.727242,1405595076
			1,37.964968,23.727183,1405595085
			1,37.964982,23.727120,1405595093
			1,37.964823,23.726953,1405595102`

		loader.Load(ride, data)

		pkg.Expect(t, ride.ID, 1)
		pkg.Expect(t, len(ride.Coordinates), 18)

		firstCoordinate := model.Coordinate{
			Latitude:  37.966660,
			Longitude: 23.728308,
			Timestamp: time.Unix(1405594957, 0),
		}

		lastCoordinate := model.Coordinate{
			Latitude:  37.964823,
			Longitude: 23.726953,
			Timestamp: time.Unix(1405595102, 0),
		}

		pkg.Expect(t, ride.Coordinates[0].Latitude, firstCoordinate.Latitude)
		pkg.Expect(t, ride.Coordinates[0].Longitude, firstCoordinate.Longitude)
		pkg.Expect(t, ride.Coordinates[0].Timestamp, firstCoordinate.Timestamp)

		pkg.Expect(t, ride.Coordinates[17].Latitude, lastCoordinate.Latitude)
		pkg.Expect(t, ride.Coordinates[17].Longitude, lastCoordinate.Longitude)
		pkg.Expect(t, ride.Coordinates[17].Timestamp, lastCoordinate.Timestamp)
	})
}
