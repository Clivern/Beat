// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"testing"
	"time"

	"bitbucket.org/clivern/beat/pkg"
)

// TestCoordinateType test cases
func TestCoordinateType(t *testing.T) {
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
		t.Run("TestCoordinateStruct", func(t *testing.T) {
			coordinate := Coordinate{
				Latitude:  tt.latitude,
				Longitude: tt.longitude,
				Timestamp: time.Unix(tt.timestamp, 0),
			}

			pkg.Expect(t, coordinate.Latitude, tt.wantLatitude)
			pkg.Expect(t, coordinate.Longitude, tt.wantLongitude)
			pkg.Expect(t, coordinate.Timestamp, time.Unix(tt.wantTimestamp, 0))

		})
	}
}

// TestGetElapsedTimeMethod test cases
func TestGetElapsedTimeMethod(t *testing.T) {
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
		t.Run("TestGetElapsedTimeMethod", func(t *testing.T) {
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

			pkg.Expect(t, timeInHour, tt.wantTimeElapsedInHour)
			pkg.Expect(t, err == nil, tt.wantTimeElapsedErrorNil)
		})
	}
}

// TestGetDistanceMethod test cases
func TestGetDistanceMethod(t *testing.T) {
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
		t.Run("TestGetDistanceMethod", func(t *testing.T) {
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

			pkg.Expect(t, distInMile, tt.wantDistanceInMile)
			pkg.Expect(t, distInKm, tt.wantDistanceInKm)
		})
	}
}

// TestGetSpeedMethod test cases
func TestGetSpeedMethod(t *testing.T) {
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
		t.Run("TestGetSpeedMethod", func(t *testing.T) {
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

			pkg.Expect(t, speed, tt.wantSpeedInKm)
			pkg.Expect(t, err == nil, tt.wantSpeedErrorNil)
		})
	}
}
