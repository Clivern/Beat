// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"testing"
	"time"

	"bitbucket.org/clivern/beat/pkg"

	"github.com/franela/goblin"
)

// TestRideType test cases
func TestRideType(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("TestRideType", func() {
		g.It("Ride struct methods should return values assigned to object properties", func() {
			var fare float64
			ride := NewRide()

			g.Assert(ride.ID).Equal(0)
			g.Assert(ride.Fare).Equal(fare)

			fare = 22.33

			ride.SetID(1)
			ride.SetFare(fare)
			g.Assert(ride.GetID()).Equal(1)
			g.Assert(ride.GetFare()).Equal(fare)
		})

		g.It("It should satisfy all provided test cases", func() {
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
				ride.AppendCoordinate(Coordinate{
					Latitude:  tt.latitude,
					Longitude: tt.longitude,
					Timestamp: time.Unix(tt.timestamp, 0),
				})

				g.Assert(len(ride.GetCoordinates())).Equal(tt.wantlength)
			}
		})
	})
}

// TestNormalizeCoordinatesMethod test cases
func TestNormalizeCoordinatesMethod(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("NormalizeCoordinates", func() {
		g.It("Ride object should return only the valid coordinates after normalization and discard invalid ones", func() {
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
			g.Assert(len(ride.GetCoordinates())).Equal(1)
			// Normalize will remove the single coordinate
			g.Assert(ride.NormalizeCoordinates()).Equal(1)
			// Now we should have a zero coordinates
			g.Assert(len(ride.GetCoordinates())).Equal(0)
		})

		g.It("Ride object should return only the valid coordinates after normalization", func() {
			ride := NewRide()
			ride.SetID(1)

			var coordinates = []struct {
				latitude  float64
				longitude float64
				timestamp int64
			}{
				{52.380746, 4.651924, 1608034878},
				{52.380337, 4.653338, 1608034923},

				// Invalid value (the speed more than 100 km/h)
				{52.371108, 4.647479, 1608034938},

				// Invalid value (the speed more than 100 km/h)
				{52.372317, 4.652988, 1608034953},

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
			g.Assert(len(ride.GetCoordinates())).Equal(7)
			// Normalize will remove the two invalid ones
			g.Assert(ride.NormalizeCoordinates()).Equal(2)
			// Now we should have 5 coordinates
			g.Assert(len(ride.GetCoordinates())).Equal(5)

			// The ones we removed should be missing
			for _, coordinate := range ride.GetCoordinates() {
				g.Assert(coordinate.Latitude != 52.371108).Equal(true)
				g.Assert(coordinate.Longitude != 4.647479).Equal(true)

				g.Assert(coordinate.Latitude != 52.372317).Equal(true)
				g.Assert(coordinate.Longitude != 4.652988).Equal(true)
			}
		})

		g.It("Ride object should return only the valid coordinates after normalization", func() {
			ride := NewRide()
			ride.SetID(1)

			var coordinates = []struct {
				latitude  float64
				longitude float64
				timestamp int64
			}{
				{52.380746, 4.651924, 1608034878},
				{52.380337, 4.653338, 1608034923},

				// Invalid value (the speed more than 100 km/h)
				{52.371108, 4.647479, 1608034938},

				// Invalid value (the speed more than 100 km/h)
				{52.372317, 4.652988, 1608034953},

				{52.379980, 4.654710, 1608034968},
				{52.379486, 4.656284, 1608035013},

				// Invalid value (the speed more than 100 km/h)
				{52.332350, 4.873741, 1608035057},

				{52.379046, 4.659400, 1608035058},

				// Invalid value (the speed more than 100 km/h)
				{52.333583, 4.887438, 1608035118},
			}

			for _, coordinate := range coordinates {
				ride.AppendCoordinate(Coordinate{
					Latitude:  coordinate.latitude,
					Longitude: coordinate.longitude,
					Timestamp: time.Unix(coordinate.timestamp, 0),
				})
			}
			// We should have 9 coordinates
			g.Assert(len(ride.GetCoordinates())).Equal(9)
			// Normalize will remove the four invalid ones
			g.Assert(ride.NormalizeCoordinates()).Equal(4)
			// Now we should have 5 coordinates
			g.Assert(len(ride.GetCoordinates())).Equal(5)

			// The ones we removed should be missing
			for _, coordinate := range ride.GetCoordinates() {
				g.Assert(coordinate.Latitude != 52.371108).Equal(true)
				g.Assert(coordinate.Longitude != 4.647479).Equal(true)

				g.Assert(coordinate.Latitude != 52.372317).Equal(true)
				g.Assert(coordinate.Longitude != 4.652988).Equal(true)

				g.Assert(coordinate.Latitude != 52.332350).Equal(true)
				g.Assert(coordinate.Longitude != 4.873741).Equal(true)

				g.Assert(coordinate.Latitude != 52.333583).Equal(true)
				g.Assert(coordinate.Longitude != 4.887438).Equal(true)
			}
		})
	})
}

// BenchmarkNormalizeCoordinates benchmark
func BenchmarkNormalizeCoordinates(b *testing.B) {
	ride := NewRide()
	ride.SetID(1)

	var coordinates = []struct {
		latitude  float64
		longitude float64
		timestamp int64
	}{
		{52.380746, 4.651924, 1608034878},
		{52.380337, 4.653338, 1608034923},

		// Invalid value (the speed more than 100 km/h)
		{52.371108, 4.647479, 1608034938},

		// Invalid value (the speed more than 100 km/h)
		{52.372317, 4.652988, 1608034953},

		{52.379980, 4.654710, 1608034968},
		{52.379486, 4.656284, 1608035013},

		// Invalid value (the speed more than 100 km/h)
		{52.332350, 4.873741, 1608035057},

		{52.379046, 4.659400, 1608035058},

		// Invalid value (the speed more than 100 km/h)
		{52.333583, 4.887438, 1608035118},
	}

	for _, coordinate := range coordinates {
		ride.AppendCoordinate(Coordinate{
			Latitude:  coordinate.latitude,
			Longitude: coordinate.longitude,
			Timestamp: time.Unix(coordinate.timestamp, 0),
		})
	}

	for n := 0; n < b.N; n++ {
		ride.NormalizeCoordinates()
		ride.GetCoordinates()
	}
}
