// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"testing"
	"time"

	"bitbucket.org/clivern/beat/pkg"
)

// TestRideType test cases
func TestRideType(t *testing.T) {
	t.Run("TestRideType", func(t *testing.T) {
		var fare float64
		ride := NewRide()

		pkg.Expect(t, ride.ID, 0)
		pkg.Expect(t, ride.Fare, fare)

		fare = 22.33

		ride.SetID(1)
		ride.SetFare(fare)
		pkg.Expect(t, ride.GetID(), 1)
		pkg.Expect(t, ride.GetFare(), fare)
	})

	var tests = []struct {
		latitude   float64
		longitude  float64
		timestamp  int64
		wantlength int
	}{
		{37.966660, 23.728308, 1405594957, 1},
		{37.966195, 23.728613, 1405595034, 2},
		{37.965377, 23.727717, 1405595068, 3},
		{38.966189, 25.728613, 1405598557, 4},
		{37.966660, 23.728308, 1405594957, 5},
		{37.966195, 23.728613, 1405595034, 6},
		{37.965377, 23.727717, 1405595068, 7},
		{38.966189, 25.728613, 1405598557, 8},
	}
	ride := NewRide()

	for _, tt := range tests {
		t.Run("TestRideType", func(t *testing.T) {
			ride.AppendCoordinate(Coordinate{
				Latitude:  tt.latitude,
				Longitude: tt.longitude,
				Timestamp: time.Unix(tt.timestamp, 0),
			})

			pkg.Expect(t, len(ride.GetCoordinates()), tt.wantlength)
		})
	}
}

// TestNormalizeCoordinatesMethod test cases
func TestNormalizeCoordinatesMethod(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	t.Run("TestNormalizeCoordinates", func(t *testing.T) {
		ride := NewRide()
		ride.SetID(1)

		var coordinates = []struct {
			latitude  float64
			longitude float64
			timestamp int64
		}{
			{37.966660, 23.728308, 1405594957},
		}

		for _, coordinate := range coordinates {
			ride.AppendCoordinate(Coordinate{
				Latitude:  coordinate.latitude,
				Longitude: coordinate.longitude,
				Timestamp: time.Unix(coordinate.timestamp, 0),
			})
		}
		// We should have one coordinate
		pkg.Expect(t, len(ride.GetCoordinates()), 1)
		// Normalize will remove the single coordinate
		pkg.Expect(t, ride.NormalizeCoordinates(), 1)
		// Now we should have a zero coordinates
		pkg.Expect(t, len(ride.GetCoordinates()), 0)
	})

	t.Run("TestNormalizeCoordinates", func(t *testing.T) {
		ride := NewRide()
		ride.SetID(1)

		var coordinates = []struct {
			latitude  float64
			longitude float64
			timestamp int64
		}{
			{52.380746, 4.651924, 1608034878},
			{52.380337, 4.653338, 1608034923},

			{52.371108, 4.647479, 1608034938}, // Invalid value (the speed more than 100 km/h)
			{52.372317, 4.652988, 1608034953}, // Invalid value (the speed more than 100 km/h)

			{52.379980, 4.654710, 1608034968},
			{52.379486, 4.656284, 1608035013},
			{52.379046, 4.659400, 1608035058},
		}

		for _, coordinate := range coordinates {
			ride.AppendCoordinate(Coordinate{
				Latitude:  coordinate.latitude,
				Longitude: coordinate.longitude,
				Timestamp: time.Unix(coordinate.timestamp, 0),
			})
		}
		// We should have 7 coordinates
		pkg.Expect(t, len(ride.GetCoordinates()), 7)
		// Normalize will remove the two invalid ones
		pkg.Expect(t, ride.NormalizeCoordinates(), 2)
		// Now we should have 5 coordinates
		pkg.Expect(t, len(ride.GetCoordinates()), 5)

		// The ones we removed should be missing
		for _, coordinate := range ride.GetCoordinates() {
			pkg.Expect(t, coordinate.Latitude != 52.371108, true)
			pkg.Expect(t, coordinate.Longitude != 4.647479, true)

			pkg.Expect(t, coordinate.Latitude != 52.372317, true)
			pkg.Expect(t, coordinate.Longitude != 4.652988, true)
		}
	})

	t.Run("TestNormalizeCoordinates", func(t *testing.T) {
		ride := NewRide()
		ride.SetID(1)

		var coordinates = []struct {
			latitude  float64
			longitude float64
			timestamp int64
		}{
			{52.380746, 4.651924, 1608034878},
			{52.380337, 4.653338, 1608034923},

			{52.371108, 4.647479, 1608034938}, // Invalid value (the speed more than 100 km/h)
			{52.372317, 4.652988, 1608034953}, // Invalid value (the speed more than 100 km/h)

			{52.379980, 4.654710, 1608034968},
			{52.379486, 4.656284, 1608035013},

			{52.332350, 4.873741, 1608035057}, // Invalid value (the speed more than 100 km/h)

			{52.379046, 4.659400, 1608035058},

			{52.333583, 4.887438, 1608035118}, // Invalid value (the speed more than 100 km/h)
		}

		for _, coordinate := range coordinates {
			ride.AppendCoordinate(Coordinate{
				Latitude:  coordinate.latitude,
				Longitude: coordinate.longitude,
				Timestamp: time.Unix(coordinate.timestamp, 0),
			})
		}
		// We should have 9 coordinates
		pkg.Expect(t, len(ride.GetCoordinates()), 9)
		// Normalize will remove the four invalid ones
		pkg.Expect(t, ride.NormalizeCoordinates(), 4)
		// Now we should have 5 coordinates
		pkg.Expect(t, len(ride.GetCoordinates()), 5)

		// The ones we removed should be missing
		for _, coordinate := range ride.GetCoordinates() {
			pkg.Expect(t, coordinate.Latitude != 52.371108, true)
			pkg.Expect(t, coordinate.Longitude != 4.647479, true)

			pkg.Expect(t, coordinate.Latitude != 52.372317, true)
			pkg.Expect(t, coordinate.Longitude != 4.652988, true)

			pkg.Expect(t, coordinate.Latitude != 52.332350, true)
			pkg.Expect(t, coordinate.Longitude != 4.873741, true)

			pkg.Expect(t, coordinate.Latitude != 52.333583, true)
			pkg.Expect(t, coordinate.Longitude != 4.887438, true)
		}
	})
}
