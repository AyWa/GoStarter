package main

import (
	"strconv"

	"github.com/aywa/goNotify/api"
	"github.com/aywa/goNotify/config"
	_ "github.com/aywa/goNotify/db"
)

func main() {
	mySettings := config.GetSetting()
	api.StartAPI(strconv.Itoa(mySettings.Port))
}
