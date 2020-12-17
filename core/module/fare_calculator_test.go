// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"math"
	"testing"
	"time"

	"bitbucket.org/clivern/beat/core/model"
	"bitbucket.org/clivern/beat/pkg"

	"github.com/franela/goblin"
	"github.com/spf13/viper"
)

// TestCalculateSegmentFare test cases
func TestCalculateSegmentFare(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("CalculateSegmentFare", func() {
		g.It("It should satisfy all provided test cases", func() {
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
				// if the two coordinates are equal
				{37.966660, 23.728308, 1405594957, 37.966660, 23.728308, 1405594957, 0, true},

				// car was moving @6:42pm (distance is 11.46 km)
				{52.316275, 4.678871, 1608056422, 52.370210, 4.535538, 1608057742, 8.462425888868383, true},

				// car was moving @1:00am (distance is 11.46 km)
				{52.316275, 4.678871, 1607994000, 52.370210, 4.535538, 1607995320, 14.866423858822838, true},

				// car was moving @1:00am (distance is 11.46 km for 1 hour)
				{52.316275, 4.678871, 1607994000, 52.370210, 4.535538, 1607997600, 14.866423858822838, true},

				// car was idle for 1.5 hours (speed is 7.64 km/hour)
				{52.316275, 4.678871, 1607994000, 52.370210, 4.535538, 1607999400, 17.85, true},
			}
			for _, tt := range tests {
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

				g.Assert(fare).Equal(tt.wantFare)
				g.Assert(err == nil).Equal(tt.wantErrorNil)
			}
		})
	})
}

// BenchmarkCalculateSegmentFare benchmark
func BenchmarkCalculateSegmentFare(b *testing.B) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	old := model.Coordinate{
		Latitude:  52.316275,
		Longitude: 4.678871,
		Timestamp: time.Unix(1608056422, 0),
	}

	new := model.Coordinate{
		Latitude:  52.370210,
		Longitude: 4.535538,
		Timestamp: time.Unix(1608057742, 0),
	}

	for n := 0; n < b.N; n++ {
		calculateSegmentFare(old, new)
	}
}

// TestCalculateRideFareTwoPoints test cases
func TestCalculateRideFareTwoPoints(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("CalculateRideFare", func() {
		g.It("It should satisfy all provided test cases", func() {
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
				// if the two coordinates are equal
				{37.966660, 23.728308, 1405594957, 37.966660, 23.728308, 1405594957, viper.GetFloat64("fare.minimum"), true},

				// car was moving @6:42pm (distance is 11.46 km)
				{52.316275, 4.678871, 1608056422, 52.370210, 4.535538, 1608057742, 8.462425888868383 + viper.GetFloat64("fare.standard_fee"), true},

				// car was moving @1:00am (distance is 11.46 km)
				{52.316275, 4.678871, 1607994000, 52.370210, 4.535538, 1607995320, 14.866423858822838 + viper.GetFloat64("fare.standard_fee"), true},

				// car was moving @1:00am (distance is 11.46 km for 1 hour)
				{52.316275, 4.678871, 1607994000, 52.370210, 4.535538, 1607997600, 14.866423858822838 + viper.GetFloat64("fare.standard_fee"), true},

				// car was idle for 1.5 hours (speed is 7.64 km/hour)
				{52.316275, 4.678871, 1607994000, 52.370210, 4.535538, 1607999400, 17.85 + viper.GetFloat64("fare.standard_fee"), true},
			}
			for _, tt := range tests {
				ride := model.NewRide()
				ride.AppendCoordinate(model.Coordinate{
					Latitude:  tt.oldLatitude,
					Longitude: tt.oldLongitude,
					Timestamp: time.Unix(tt.OldTimestamp, 0),
				})

				ride.AppendCoordinate(model.Coordinate{
					Latitude:  tt.newLatitude,
					Longitude: tt.newLongitude,
					Timestamp: time.Unix(tt.newTimestamp, 0),
				})

				fare, err := CalculateRideFare(ride)

				// The fare plus the standard flag amount
				g.Assert(fare).Equal(float64(tt.wantFare))
				g.Assert(err == nil).Equal(tt.wantErrorNil)
			}
		})
	})
}

// TestCalculateRideFareMultiplePoints test cases
func TestCalculateRideFareMultiplePoints(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("CalculateRideFare", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				latitude  float64
				longitude float64
				timestamp int64

				wantFare     float64
				wantErrorNil bool
			}{
				{64.29357012490215, -15.444242456502462, 1608111032, viper.GetFloat64("fare.minimum"), true},

				// Add another coordinate but the car didn't move
				{64.29357012490215, -15.444242456502462, 1608111032, viper.GetFloat64("fare.minimum"), true},

				// Add another coordinate with 19.10km distance and 19.10km/h speed
				{64.186612, -15.751840, 1608114632, 14.08 + viper.GetFloat64("fare.standard_fee"), true},

				// Add another coordinate with 10.67km distance and 10.67km/h speed
				{64.150310, -15.954850, 1608118232, 14.08 + viper.GetFloat64("fare.standard_fee") + 7.87, true},

				// Add another coordinate with 7.30km distance and 7.30km/h speed
				{64.116614, -16.083341, 1608121832, 14.08 + viper.GetFloat64("fare.standard_fee") + 7.87 + 11.90, true},

				// Add another coordinate with 31.37km distance and 31.37km/h speed
				{63.914866563139086, -16.530649284050384, 1608125432, 14.08 + viper.GetFloat64("fare.standard_fee") + 7.87 + 11.90 + 23.14, true},
			}

			ride := model.NewRide()

			// At each iteration, we will add a new coordinate and calculate the price
			for _, tt := range tests {
				ride.AppendCoordinate(model.Coordinate{
					Latitude:  tt.latitude,
					Longitude: tt.longitude,
					Timestamp: time.Unix(tt.timestamp, 0),
				})

				fare, err := CalculateRideFare(ride)

				// The fare plus the standard flag amount
				g.Assert(math.Floor(fare*100) / 100).Equal(tt.wantFare)
				g.Assert(err == nil).Equal(tt.wantErrorNil)
			}
		})
	})
}

// BenchmarkCalculateRideFare benchmark
func BenchmarkCalculateRideFare(b *testing.B) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	ride := model.NewRide()
	ride.SetID(1)

	var coordinates = []struct {
		latitude  float64
		longitude float64
		timestamp int64
	}{
		{64.29357012490215, -15.444242456502462, 1608111032},
		{64.29357012490215, -15.444242456502462, 1608111032},
		{64.186612, -15.751840, 1608114632},
		{64.150310, -15.954850, 1608118232},
		{64.116614, -16.083341, 1608121832},
		{63.914866563139086, -16.530649284050384, 1608125432},
	}

	for _, coordinate := range coordinates {
		ride.AppendCoordinate(model.Coordinate{
			Latitude:  coordinate.latitude,
			Longitude: coordinate.longitude,
			Timestamp: time.Unix(coordinate.timestamp, 0),
		})
	}

	for n := 0; n < b.N; n++ {
		CalculateRideFare(ride)
	}
}
