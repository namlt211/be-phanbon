package main

import (
	"green/helpers/mytime"
	"green/models"
	"green/routers"
)

func main() {

	_ = mytime.SetTimezone("UTC")

	models.ConnectDB()
	routers.NewRouter()
}