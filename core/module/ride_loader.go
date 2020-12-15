// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"strings"
	"time"

	"bitbucket.org/clivern/beat/core/model"
	"bitbucket.org/clivern/beat/core/util"
)

// RideLoader interface
type RideLoader interface {
	Load(*model.Ride, string) (*model.Ride, error)
}

// CSVLoader struct type
type CSVLoader struct {
}

// Load load a CSV data of a complete ride into ride object
// CSV data provided in the form of (id_ride, lat, lng, timestamp)
func (c CSVLoader) Load(ride *model.Ride, csv string) (*model.Ride, error) {
	var itemsPerLine []string
	var err error
	var id int
	var lat float64
	var lng float64
	var timestamp time.Time

	lines := strings.Split(csv, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])

		if strings.TrimSpace(lines[i]) == "" {
			continue
		}

		itemsPerLine = strings.Split(lines[i], ",")

		id, err = util.StringToInt(itemsPerLine[0])

		if err != nil {
			return ride, err
		}

		lat, err = util.StringToFloat64(itemsPerLine[1])

		if err != nil {
			return ride, err
		}

		lng, err = util.StringToFloat64(itemsPerLine[2])

		if err != nil {
			return ride, err
		}

		timestamp, err = util.StringToTimestamp(itemsPerLine[3])

		if err != nil {
			return ride, err
		}

		ride.SetID(id)
		ride.AppendCoordinate(model.Coordinate{
			Latitude:  lat,
			Longitude: lng,
			Timestamp: timestamp,
		})
	}

	return ride, nil
}
