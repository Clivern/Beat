// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"math"
	"time"

	"bitbucket.org/clivern/beat/core/util"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

// Coordinate struct type
type Coordinate struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

// GetDistance gets a distance from another new coordinate
// Calculations based on Haversine Formula https://en.wikipedia.org/wiki/Haversine_formula
func (p *Coordinate) GetDistance(newCoordinate Coordinate) (inMile, inKm float64) {
	lat1, lng1 := p.toRadians()
	lat2, lng2 := newCoordinate.toRadians()

	diffLat := lat2 - lat1
	diffLon := lng2 - lng1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	inMile = c * earthRadiusMi
	inKm = c * earthRaidusKm

	return inMile, inKm
}

// GetElapsedTime gets the elapsed time in hours to move to a new coordinate
func (p *Coordinate) GetElapsedTime(newCoordinate Coordinate) (float64, error) {
	diff := newCoordinate.Timestamp.Sub(p.Timestamp)
	return util.StringToFloat64(fmt.Sprintf("%f", diff.Hours()))
}

// GetSpeed gets the movement speed in Km/hour to another coordinate
func (p *Coordinate) GetSpeed(newCoordinate Coordinate) (float64, error) {
	var result float64

	_, inKm := p.GetDistance(newCoordinate)
	timeElapsed, err := p.GetElapsedTime(newCoordinate)

	if timeElapsed == 0 {
		return result, nil
	}

	if err != nil {
		return result, err
	}

	result = inKm / timeElapsed

	return util.StringToFloat64(fmt.Sprintf("%.2f", result))
}

// toRadians converts latitude and longitude from degrees to radians.
func (p *Coordinate) toRadians() (float64, float64) {
	return p.Latitude * math.Pi / 180, p.Longitude * math.Pi / 180
}
