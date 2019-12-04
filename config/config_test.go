// Copyright (c) 2016 The daqdata developers. All rights reserved.
// Project site: https://github.com/gotmc/daqdata
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package config

import (
	"io/ioutil"
	"os"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

const (
	configFilename = "./daqdata_config.json"
)

func setupGoodConfig() {
	config := `{
		"num_channels": 8,
		"channels": 1,
		"slopes": [0.0, 0.01, 0.20, 3.00, 40.0, 500.0, 6000.0, 70000.0],
		"offsets": [1.0, 0.01, 0.20, 3.00, 40.0, 500.0, 6000.0, 70000.0],
		"device": {
			"sn": "01ACD31D",
			"mfr": "MCC"
		}

	}`
	ioutil.WriteFile(configFilename, []byte(config), 0644)
}

func TestReadConfiguration(t *testing.T) {
	setupGoodConfig()
	c.Convey("Given I want to parse a DAQDATA configuration file", t, func() {
		config, err := Parse(configFilename)
		c.Convey("When the config file is read", func() {
			c.Convey("There shouldn't be an error", func() {
				c.So(err, c.ShouldBeNil)
			})
			c.Convey("The number of channels should be 8", func() {
				c.So(config.NumChannels, c.ShouldEqual, 8)
			})
			c.Convey("The manufacturer should be MCC", func() {
				c.So(config.Device.Manufacturer, c.ShouldEqual, "MCC")
			})
			c.Convey("The array of slopes should be present and correct", func() {
				c.So(config.Slopes, c.ShouldResemble,
					[]float32{0.0, 0.01, 0.20, 3.00, 40.00, 500.0, 6000.0, 70000.0})
			})
			c.Convey("The array of offsets should be present and correct", func() {
				c.So(config.Offsets, c.ShouldResemble,
					[]float32{1.0, 0.01, 0.20, 3.00, 40.00, 500.0, 6000.0, 70000.0})
			})
		})
	})
	os.Remove(configFilename)
}
