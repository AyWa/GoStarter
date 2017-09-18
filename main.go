package main

import (
	"fmt"

	"github.com/aywa/goNotify/config"
)

func main() {
	myFlags := config.GetFlag()
	mySettings := config.GetSetting()
	fmt.Println(*myFlags.IsDev)
	fmt.Println(mySettings.Hostname)
}
