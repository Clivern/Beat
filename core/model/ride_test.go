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
	// TestRideStruct
	t.Run("TestRideStruct", func(t *testing.T) {
		var fare float64
		var latitude float64
		var longitude float64

		ride := NewRide()

		pkg.Expect(t, ride.ID, 0)
		pkg.Expect(t, ride.Fare, fare)

		fare = 22.33

		ride.SetID(1)
		ride.SetFare(fare)
		pkg.Expect(t, ride.GetID(), 1)
		pkg.Expect(t, ride.GetFare(), fare)

		latitude = 37.966660
		longitude = 23.728308

		ride.AppendCoordinate(Coordinate{
			Latitude:  37.966660,
			Longitude: 23.728308,
			Timestamp: time.Unix(1405594957, 0),
		})

		ride.AppendCoordinate(Coordinate{
			Latitude:  40.966660,
			Longitude: 42.728308,
			Timestamp: time.Unix(1405594959, 0),
		})

		pkg.Expect(t, len(ride.GetCoordinates()), 2)

		coordinates := ride.GetCoordinates()

		pkg.Expect(t, coordinates[0].Latitude, latitude)
		pkg.Expect(t, coordinates[0].Longitude, longitude)

	})
}

// TestNormalizeCoordinatesMethod test cases
func TestNormalizeCoordinatesMethod(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	t.Run("TestNormalizeCoordinates01", func(t *testing.T) {
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

	t.Run("TestNormalizeCoordinates02", func(t *testing.T) {
		ride := NewRide()
		ride.SetID(1)

		var coordinates = []struct {
			latitude  float64
			longitude float64
			timestamp int64
		}{
			{52.380746, 4.651924, 1608034878},
			{52.380337, 4.653338, 1608034923},

			{52.371108, 4.647479, 1608034938}, // Invalid value
			{52.372317, 4.652988, 1608034953}, // Invalid value

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
}
