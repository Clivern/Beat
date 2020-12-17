// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"testing"
	"time"

	"github.com/franela/goblin"
)

// TestCoordinateType test cases
func TestCoordinateType(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("CoordinateStruct", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				latitude      float64
				longitude     float64
				timestamp     int64
				wantLatitude  float64
				wantLongitude float64
				wantTimestamp int64
			}{
				{37.966660, 23.728308, 1405594957, 37.966660, 23.728308, 1405594957},
				{37.966195, 23.728613, 1405595034, 37.966195, 23.728613, 1405595034},
				{37.965377, 23.727717, 1405595068, 37.965377, 23.727717, 1405595068},
				{38.966189, 25.728613, 1405598557, 38.966189, 25.728613, 1405598557},
			}

			for _, tt := range tests {
				coordinate := Coordinate{
					Latitude:  tt.latitude,
					Longitude: tt.longitude,
					Timestamp: time.Unix(tt.timestamp, 0),
				}

				g.Assert(coordinate.Latitude).Equal(tt.wantLatitude)
				g.Assert(coordinate.Longitude).Equal(tt.wantLongitude)
				g.Assert(coordinate.Timestamp).Equal(time.Unix(tt.wantTimestamp, 0))
			}
		})
	})
}

// TestGetElapsedTimeMethod test cases
func TestGetElapsedTimeMethod(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GetElapsedTime", func() {
		g.It("Elapsed time should equal the value provided in the test cases", func() {
			var tests = []struct {
				oldLatitude  float64
				oldLongitude float64
				oldTimestamp int64
				newLatitude  float64
				newLongitude float64
				newTimestamp int64

				wantTimeElapsedInHour   float64
				wantTimeElapsedErrorNil bool
			}{
				{37.966660, 23.728308, 1405594957, 37.966660, 23.728308, 1405594957, 0, true},
				{37.966660, 23.728308, 1405594957, 37.966195, 23.728613, 1405595034, 0.021389, true},
				{37.966660, 23.728308, 1405594957, 37.965377, 23.727717, 1405595068, 0.030833, true},
				{37.966660, 23.728308, 1405594957, 38.966189, 25.728613, 1405598557, 1, true},
			}

			for _, tt := range tests {
				coordinate := Coordinate{
					Latitude:  tt.oldLatitude,
					Longitude: tt.oldLongitude,
					Timestamp: time.Unix(tt.oldTimestamp, 0),
				}

				newCoordinate := Coordinate{
					Latitude:  tt.newLatitude,
					Longitude: tt.newLongitude,
					Timestamp: time.Unix(tt.newTimestamp, 0),
				}

				timeInHour, err := coordinate.GetElapsedTime(newCoordinate)

				g.Assert(timeInHour).Equal(tt.wantTimeElapsedInHour)
				g.Assert(err == nil).Equal(tt.wantTimeElapsedErrorNil)
			}
		})
	})
}

// BenchmarkGetElapsedTime benchmark
func BenchmarkGetElapsedTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coordinate := Coordinate{
			Latitude:  37.966660,
			Longitude: 23.728308,
			Timestamp: time.Unix(1405594957, 0),
		}

		newCoordinate := Coordinate{
			Latitude:  38.966189,
			Longitude: 25.728613,
			Timestamp: time.Unix(1405598557, 0),
		}

		coordinate.GetElapsedTime(newCoordinate)
	}
}

// TestGetDistanceMethod test cases
func TestGetDistanceMethod(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GetDistance", func() {
		g.It("Distance should equal the value provided in the test cases", func() {
			var tests = []struct {
				oldLatitude        float64
				oldLongitude       float64
				oldTimestamp       int64
				newLatitude        float64
				newLongitude       float64
				newTimestamp       int64
				wantDistanceInMile float64
				wantDistanceInKm   float64
			}{
				{37.966660, 23.728308, 1405594957, 37.966660, 23.728308, 1405594957, 0, 0},
				{37.966660, 23.728308, 1405594957, 37.966195, 23.728613, 1405595034, 0.03616282372892886, 0.058209537639465826},
				{37.966660, 23.728308, 1405594957, 37.965377, 23.727717, 1405595068, 0.0942932369664417, 0.15177923514734717},
				{37.966660, 23.728308, 1405594957, 38.966189, 25.728613, 1405598557, 128.34254713922107, 206.58675286103525},
			}

			for _, tt := range tests {
				coordinate := Coordinate{
					Latitude:  tt.oldLatitude,
					Longitude: tt.oldLongitude,
					Timestamp: time.Unix(tt.oldTimestamp, 0),
				}

				newCoordinate := Coordinate{
					Latitude:  tt.newLatitude,
					Longitude: tt.newLongitude,
					Timestamp: time.Unix(tt.newTimestamp, 0),
				}

				distInMile, distInKm := coordinate.GetDistance(newCoordinate)

				g.Assert(distInMile).Equal(tt.wantDistanceInMile)
				g.Assert(distInKm).Equal(tt.wantDistanceInKm)
			}
		})
	})
}

// BenchmarkGetDistance benchmark
func BenchmarkGetDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coordinate := Coordinate{
			Latitude:  37.966660,
			Longitude: 23.728308,
			Timestamp: time.Unix(1405594957, 0),
		}

		newCoordinate := Coordinate{
			Latitude:  38.966189,
			Longitude: 25.728613,
			Timestamp: time.Unix(1405598557, 0),
		}

		coordinate.GetDistance(newCoordinate)
	}
}

// TestGetSpeedMethod test cases
func TestGetSpeedMethod(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GetSpeed", func() {
		g.It("Speed should equal the value provided in the test cases", func() {
			var tests = []struct {
				oldLatitude       float64
				oldLongitude      float64
				oldTimestamp      int64
				newLatitude       float64
				newLongitude      float64
				newTimestamp      int64
				wantSpeedInKm     float64
				wantSpeedErrorNil bool
			}{
				{37.966660, 23.728308, 1405594957, 37.966660, 23.728308, 1405594957, 0, true},
				{37.966660, 23.728308, 1405594957, 37.966195, 23.728613, 1405595034, 2.72, true},
				{37.966660, 23.728308, 1405594957, 37.965377, 23.727717, 1405595068, 4.92, true},
				{37.966660, 23.728308, 1405594957, 38.966189, 25.728613, 1405598557, 206.59, true},
			}
			for _, tt := range tests {
				coordinate := Coordinate{
					Latitude:  tt.oldLatitude,
					Longitude: tt.oldLongitude,
					Timestamp: time.Unix(tt.oldTimestamp, 0),
				}

				newCoordinate := Coordinate{
					Latitude:  tt.newLatitude,
					Longitude: tt.newLongitude,
					Timestamp: time.Unix(tt.newTimestamp, 0),
				}

				speed, err := coordinate.GetSpeed(newCoordinate)

				g.Assert(speed).Equal(tt.wantSpeedInKm)
				g.Assert(err == nil).Equal(tt.wantSpeedErrorNil)
			}
		})
	})
}

// BenchmarkGetSpeed benchmark
func BenchmarkGetSpeed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		coordinate := Coordinate{
			Latitude:  37.966660,
			Longitude: 23.728308,
			Timestamp: time.Unix(1405594957, 0),
		}

		newCoordinate := Coordinate{
			Latitude:  38.966189,
			Longitude: 25.728613,
			Timestamp: time.Unix(1405598557, 0),
		}

		coordinate.GetSpeed(newCoordinate)
	}
}
