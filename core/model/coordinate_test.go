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

// TestCoordinateType test cases
func TestCoordinateType(t *testing.T) {
	// TestCoordinateStruct
	t.Run("TestCoordinateStruct", func(t *testing.T) {
		var latitude float64
		var longitude float64

		latitude = 37.966660
		longitude = 23.728308

		coordinate := Coordinate{
			Latitude:  37.966660,
			Longitude: 23.728308,
			Timestamp: time.Unix(1405594959, 0),
		}

		pkg.Expect(t, coordinate.Latitude, latitude)
		pkg.Expect(t, coordinate.Longitude, longitude)

	})
}

// TestCoordinateTypeMethods test cases
func TestCoordinateTypeMethods(t *testing.T) {
	coordinate := Coordinate{
		Latitude:  37.966660,
		Longitude: 23.728308,
		Timestamp: time.Unix(1405594957, 0),
	}

	var tests = []struct {
		latitude           float64
		longitude          float64
		timestamp          int64
		wantDistanceInMile float64
		wantDistanceInKm   float64
		wantTimeInHour     float64
		wantTimeErrorNil   bool
		wantSpeed          float64
		wantSpeedErrorNil  bool
	}{
		{37.966660, 23.728308, 1405594957, 0, 0, 0, true, 0, true},
		{37.966195, 23.728613, 1405595034, 0.03616282372892886, 0.058209537639465826, 0.021389, true, 2.72, true},
		{37.965377, 23.727717, 1405595068, 0.0942932369664417, 0.15177923514734717, 0.030833, true, 4.92, true},
		{38.966189, 25.728613, 1405598557, 128.34254713922107, 206.58675286103525, 1, true, 206.59, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("TestCoordinateTypeMethods(%f,%f)", tt.latitude, tt.longitude)

		t.Run(testname, func(t *testing.T) {

			newCoordinate := Coordinate{
				Latitude:  tt.latitude,
				Longitude: tt.longitude,
				Timestamp: time.Unix(tt.timestamp, 0),
			}

			distInMile, distInKm := coordinate.GetDistance(newCoordinate)

			timeInHour, err1 := coordinate.GetElapsedTime(newCoordinate)

			speed, err2 := coordinate.GetSpeed(newCoordinate)

			pkg.Expect(t, distInMile, tt.wantDistanceInMile)
			pkg.Expect(t, distInKm, tt.wantDistanceInKm)
			pkg.Expect(t, timeInHour, tt.wantTimeInHour)
			pkg.Expect(t, err1 == nil, tt.wantTimeErrorNil)
			pkg.Expect(t, speed, tt.wantSpeed)
			pkg.Expect(t, err2 == nil, tt.wantSpeedErrorNil)
		})
	}
}
