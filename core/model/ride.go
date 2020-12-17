// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Ride struct type
type Ride struct {
	ID          int          `json:"id"`
	Coordinates []Coordinate `json:"coordinates"`
	Fare        float64      `json:"fare"`
}

// NewRide creates a new instance of Ride
func NewRide() *Ride {
	return &Ride{
		ID:          0,
		Coordinates: make([]Coordinate, 0),
		Fare:        0,
	}
}

// AppendCoordinate add a new coordinate to the ride
func (r *Ride) AppendCoordinate(coordinate Coordinate) {
	r.Coordinates = append(r.Coordinates, coordinate)
}

// SetID sets ride id
func (r *Ride) SetID(id int) {
	r.ID = id
}

// SetFare set ride fare
func (r *Ride) SetFare(fare float64) {
	r.Fare = fare
}

// GetFare gets ride fare
func (r *Ride) GetFare() float64 {
	return r.Fare
}

// GetID gets ride ID
func (r *Ride) GetID() int {
	return r.ID
}

// GetCoordinates gets ride coordinates
func (r *Ride) GetCoordinates() []Coordinate {
	return r.Coordinates
}

// NormalizeCoordinates removes invalid coordinate and return the count.
// a coordinate is considered invalid if the speed used to reach that
// coordinate from the previous one is more than 100 Km/h
func (r *Ride) NormalizeCoordinates() int {
	normalizedCoordinates := make([]Coordinate, 0)

	log.Debug(fmt.Sprintf(
		"Normalize coordinates for ride with ID %d",
		r.ID,
	))

	for index, coordinate := range r.Coordinates {
		if index == len(r.Coordinates)-1 {
			break
		}

		if len(normalizedCoordinates) == 0 {
			normalizedCoordinates = append(normalizedCoordinates, coordinate)
		}

		// Get the speed from the last normalized coordinate
		speed, err := normalizedCoordinates[len(normalizedCoordinates)-1].GetSpeed(r.Coordinates[index+1])

		log.Debug(fmt.Sprintf(
			"Speed for coodinate (%f, %f, %s) and coodinate (%f, %f, %s) is %f km/hour",
			normalizedCoordinates[len(normalizedCoordinates)-1].Latitude,
			normalizedCoordinates[len(normalizedCoordinates)-1].Longitude,
			normalizedCoordinates[len(normalizedCoordinates)-1].Timestamp,
			r.Coordinates[index+1].Latitude,
			r.Coordinates[index+1].Longitude,
			r.Coordinates[index+1].Timestamp,
			speed,
		))

		if err == nil && speed <= viper.GetFloat64("segment.max_speed_threshold") {
			normalizedCoordinates = append(normalizedCoordinates, r.Coordinates[index+1])
		} else {
			log.Debug(fmt.Sprintf(
				"Remove invalid coodinate (%f, %f, %s) because speed is %f km/hour more than %f km/hour",
				r.Coordinates[index+1].Latitude,
				r.Coordinates[index+1].Longitude,
				r.Coordinates[index+1].Timestamp,
				speed,
				viper.GetFloat64("segment.max_speed_threshold"),
			))
		}
	}

	invalidCoordinatesCount := len(r.Coordinates) - len(normalizedCoordinates)

	log.Debug(fmt.Sprintf(
		"Total invalid coordinates for ride with ID %d is %d",
		r.ID,
		invalidCoordinatesCount,
	))

	r.Coordinates = normalizedCoordinates

	return invalidCoordinatesCount
}
