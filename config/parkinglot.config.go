package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	FirstNHour         uint    `envconfig:"PARKING_LOT_FIRST_N_HOUR" default:"2"`
	CostFirstNHour     float32 `envconfig:"PARKING_LOT_COST_FIRST_N_HOUR" default:"10"`
	CostAdditionalHour float32 `envconfig:"PARKING_LOT_COST_ADDITIONAL_HOUR" default:"5"`
}

// singleton of data
var data *Config

func Get() *Config {
	if data == nil {
		data = &Config{}
		envconfig.MustProcess("", data)
	}

	return data
}
