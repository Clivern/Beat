// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"bitbucket.org/clivern/beat/core/model"
	"bitbucket.org/clivern/beat/core/util"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GenerateData sends a ride data as string to a channel
// The ride data is a certain number of lines containing coordinates (segments)
func GenerateData(filePath string) (<-chan string, error) {

	channel := make(chan string)

	if !util.FileExists(filePath) {
		return channel, fmt.Errorf("File %s not found", filePath)
	}

	go func() {
		var rideInfo string
		var previousRideID string
		var rideBatch string

		file, _ := os.Open(filePath)

		defer file.Close()

		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')

			// If end of lines reached
			if err == io.EOF && rideInfo == "" {
				break
			} else if err == io.EOF {
				// Send last ride info
				channel <- rideInfo
				break
			}

			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}

			lineItems := strings.Split(line, ",")

			currentRideID := lineItems[0]

			if previousRideID == "" {
				previousRideID = lineItems[0]
			}

			if currentRideID == previousRideID {
				rideInfo = fmt.Sprintf("%s\n%s", rideInfo, line)
				continue
			}

			previousRideID = ""
			rideBatch = strings.TrimSpace(rideInfo)
			rideInfo = strings.TrimSpace(line)

			channel <- rideBatch
		}

		close(channel)
	}()

	return channel, nil
}

// ProcessData gets a ride data as string from input channel and send the ride id and the
// fare estimate to output channel
func ProcessData(inputChannel <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		wg := &sync.WaitGroup{}

		// Limit the number of goroutines
		for t := 0; t < viper.GetInt("app.max_goroutines"); t++ {
			wg.Add(1)
			go ProcessRide(inputChannel, outChannel, wg)
		}

		wg.Wait()

		close(outChannel)
	}()

	return outChannel
}

// ProcessRide calculates the ride fare
func ProcessRide(inputChannel <-chan string, outChannel chan<- string, wg *sync.WaitGroup) {
	for lines := range inputChannel {
		ride := model.NewRide()
		loader := CSVLoader{}

		// Load CSV data into the ride object
		loader.Load(ride, lines)

		// Remove invalid coordinates
		ride.NormalizeCoordinates()

		// Calculate The fare
		fare, err := CalculateRideFare(ride)

		if err != nil {
			log.Debug(fmt.Sprintf(
				"Error while calculating ride %d fare: %s",
				ride.GetID(),
				err.Error(),
			))
		}

		ride.SetFare(fare)

		outChannel <- fmt.Sprintf("%d,%.2f", ride.ID, fare)
	}

	wg.Done()
}

// StoreData store ride id and fare into a file. it gets the values from input channel
func StoreData(filePath string, channel <-chan string) error {
	if util.FileExists(filePath) {
		if err := util.DeleteFile(filePath); err != nil {
			return fmt.Errorf("Error! Unable to delete file %s", filePath)
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf(
			"Error! Unable to write to file %s: %s",
			filePath,
			err.Error(),
		)
	}

	defer file.Close()

	for line := range channel {
		if _, err := file.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
			return fmt.Errorf(
				"Error! Unable to write to file %s: %s",
				filePath,
				err.Error(),
			)
		}
	}

	return nil
}
