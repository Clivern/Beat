// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/viper"

	"bitbucket.org/clivern/beat/core/model"
	"bitbucket.org/clivern/beat/pkg"
)

// TestCalculateSegmentFare test cases
func TestCalculateSegmentFare(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	var tests = []struct {
		oldLatitude  float64
		oldLongitude float64
		OldTimestamp int64

		newLatitude  float64
		newLongitude float64
		newTimestamp int64

		wantFare     float64
		wantErrorNil bool
	}{
		{37.966660, 23.728308, 1405594957, 37.966660, 23.728308, 1405594957, 0, true},
		{52.316275, 4.678871, 1608056422, 52.370210, 4.535538, 1608057742, 8.462425888868383, true},
	}

	for _, tt := range tests {
		t.Run("TestCalculateSegmentFare", func(t *testing.T) {

			old := model.Coordinate{
				Latitude:  tt.oldLatitude,
				Longitude: tt.oldLongitude,
				Timestamp: time.Unix(tt.OldTimestamp, 0),
			}

			new := model.Coordinate{
				Latitude:  tt.newLatitude,
				Longitude: tt.newLongitude,
				Timestamp: time.Unix(tt.newTimestamp, 0),
			}

			fare, err := calculateSegmentFare(old, new)

			pkg.Expect(t, fare, tt.wantFare)
			pkg.Expect(t, err == nil, tt.wantErrorNil)
		})
	}
}

// TestCalculateRideFare test cases
func TestCalculateRideFare(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	t.Run("CalculateRideFareTwoPoints", func(t *testing.T) {
		ride := model.NewRide()
		ride.AppendCoordinate(model.Coordinate{
			Latitude:  52.316275,
			Longitude: 4.678871,
			Timestamp: time.Unix(1608056422, 0),
		})

		ride.AppendCoordinate(model.Coordinate{
			Latitude:  52.370210,
			Longitude: 4.535538,
			Timestamp: time.Unix(1608057742, 0),
		})

		fare, err := CalculateRideFare(ride)

		// The fare plus the standard flag amount
		pkg.Expect(t, fare, float64(8.462425888868383)+viper.GetFloat64("fare.standard_fee"))
		pkg.Expect(t, err, nil)
	})
}
