package main

import (
	"green/helpers/mytime"
	"green/models"
	"green/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)
var (
	engine   *gin.Engine
	cfg      *viper.Viper
	termChan chan os.Signal
)


func main() {

	_ = mytime.SetTimezone("UTC")
	engine = gin.New()
	models.ConnectDB()
	routers.NewRouter(engine)
}