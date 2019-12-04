// Copyright (c) 2016 The daqdata developers. All rights reserved.
// Project site: https://github.com/gotmc/daqdata
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// Config contains the configuration information from the header file.
type Config struct {
	Device         Device    `json:"device"`
	NumChannels    int       `json:"num_channels"`
	Channels       byte      `json:channels"`
	Slopes         []float32 `json:"slopes"`
	Offsets        []float32 `json:"offsets"`
	Notes          string    `json:"notes"`
	Timestamp      time.Time `json:"timestamp"`
	BinaryFiles    []string  `json:"binary_files"`
	SamplesPerFile []int     `json:"samples"`
	Samples        int       `json:"samples"`
	SamplingRate   int       `json:"sampling_rate"`
	Resolution     int       `json:"resolution"`
	Signed         bool      `json:"signed"`
}

// Device contains information about the device that captured the DAQ data.
type Device struct {
	Manufacturer string `json:"mfr"`
	SerialNumber string `json:"sn"`
}

// Parse parses the given configuration JSON file and returns the configuration
// struct.
func Parse(filename string) (Config, error) {
	var config Config
	configurationBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(configurationBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

// Save saves the configuration struct into the JSON configuration file.
func (c *Config) Save(filename string) error {
	jsonConfiguration, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling config into JSON %s", err)
	}
	err = ioutil.WriteFile(filename, jsonConfiguration, 0644)
	if err != nil {
		return fmt.Errorf("Error writing config file into %s: %s", filename, err)
	}
	return nil
}
