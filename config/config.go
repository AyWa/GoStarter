package config

import (
	"flag"
	"strconv"
)

// Flag is a struct containing all the information of the flag
type Flag struct {
	IsDev        *bool
	ShouldInitDb *bool
	Port         *int
}

// Setting is a struct containing the real setting from the flag
type Setting struct {
	Hostname string
	Port     int
}

var flagConfig Flag
var settingConfig Setting

func init() {
	flagConfig = Flag{
		IsDev:        flag.Bool("isDev", false, "define if the package should run in dev mode"),
		Port:         flag.Int("port", 8080, "define the port of the server"),
		ShouldInitDb: flag.Bool("shouldInitDb", false, "define if the package should init the db"),
	}
	flag.Parse()
	initializeSetting()
}

func initializeSetting() {
	settingConfig.Port = *flagConfig.Port
	if *flagConfig.IsDev == true {
		settingConfig.Hostname = "localhost" + strconv.Itoa(settingConfig.Port)
	} else {
		settingConfig.Hostname = "myHostname.com"
	}
}

// GetFlag return the initialize from flag struct
func GetFlag() Flag {
	return flagConfig
}

// GetSetting return the initialize from flag Setting struct
func GetSetting() Setting {
	return settingConfig
}
