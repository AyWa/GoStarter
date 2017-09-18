package config

import "flag"

// Flag is a struct containing all the information of the flag
type Flag struct {
	IsDev *bool
}

// Setting is a struct containing the real setting from the flag
type Setting struct {
	Hostname string
}

var flagConfig Flag
var settingConfig Setting

func init() {
	flagConfig = Flag{
		IsDev: flag.Bool("isDev", false, "define if the package should run in dev mode"),
	}
	flag.Parse()
	initializeSetting()
}

func initializeSetting() {
	if *flagConfig.IsDev == true {
		settingConfig.Hostname = "localhost:8080"
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
