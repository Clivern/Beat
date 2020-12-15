// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"

	"bitbucket.org/clivern/beat/core/model"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// CalculateRideFare calculates the whole ride fare
func CalculateRideFare(ride *model.Ride) (float64, error) {
	// Init total from the standard fee
	total := viper.GetFloat64("fare.standard_fee")

	coordinates := ride.GetCoordinates()

	for index, coordinate := range coordinates {
		// If it is the last element, break
		if index == len(coordinates)-1 {
			break
		}

		// Calculate the segment fare
		subTotal, err := calculateSegmentFare(coordinate, coordinates[index+1])

		if err != nil {
			return total, err
		}

		log.Debug(fmt.Sprintf(
			"Segment fare for coodinate (%f, %f, %s) and coodinate (%f, %f, %s) is %f",
			coordinate.Latitude,
			coordinate.Longitude,
			coordinate.Timestamp,
			coordinates[index+1].Latitude,
			coordinates[index+1].Longitude,
			coordinates[index+1].Timestamp,
			subTotal,
		))

		// Add segment fare to the total price
		total += subTotal
	}

	// If fare is less than the minimum, override with the
	// minimum value
	if total < viper.GetFloat64("fare.minimum") {
		total = viper.GetFloat64("fare.minimum")
	}

	log.Debug(fmt.Sprintf(
		"Total fare for ride with ID %d is %f",
		ride.GetID(),
		total,
	))

	return total, nil
}

// calculateSegmentFare calculates the fare for a segment,
// segment is just a two coordinates
func calculateSegmentFare(oldCoordinate model.Coordinate, newCoordinate model.Coordinate) (float64, error) {
	var total float64

	speed, err := oldCoordinate.GetSpeed(newCoordinate)

	if err != nil {
		return total, err
	}

	if speed > viper.GetFloat64("segment.pricing.idle.min_threshold") {
		// The car was moving
		_, distance := oldCoordinate.GetDistance(newCoordinate)

		// Segment start hour
		hour, _, _ := oldCoordinate.Timestamp.Clock()

		// If hour is more or equal 05:00 and less than or equal 23:00
		if hour >= 5 && hour <= 23 {
			total += distance * viper.GetFloat64("segment.pricing.moving.from_05_00_per_km")
		} else {
			total += distance * viper.GetFloat64("segment.pricing.moving.from_00_05_per_km")
		}
	} else {
		// the car was idle
		timeElapsed, err := oldCoordinate.GetElapsedTime(newCoordinate)

		if err != nil {
			return total, err
		}

		total += viper.GetFloat64("segment.pricing.idle.price_per_hour") * timeElapsed
	}

	return total, nil
}
