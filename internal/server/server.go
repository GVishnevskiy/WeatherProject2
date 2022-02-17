package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GVishnevskiy/WeatherProject2/internal/api"
	"github.com/GVishnevskiy/WeatherProject2/internal/entities"
	"github.com/GVishnevskiy/WeatherProject2/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"os"
	"time"
)

var redisClient *redis.Client
var ctx = context.Background()

func StartServer(router *gin.Engine) {
	router.LoadHTMLGlob("html/templates/*.html")
	addPageListeners(router)

	url, _ := os.LookupEnv("SITE_URL")
	logger.LogData(url)
	port, _ := os.LookupEnv("SITE_PORT")
	redisClient = SetupRedis(url, port)
	err := router.Run(port)
	if logger.LogErr(err) {
		return
	}
}

func SetupRedis(url string, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

func addPageListeners(router *gin.Engine) {
	router.GET("/weather", handleWeatherRequest)
	router.GET("/", startPage)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
}

func startPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func handleWeatherRequest(c *gin.Context) {
	city := c.Query("city")

	var weather entities.Weather

	weatherJson, err := redisClient.Get(city).Bytes()
	if logger.LogErr(err) {
		weather, err = api.GetWeather(city)
		json, err := json.Marshal(weather)
		redisClient.Set(city, json, time.Minute)
		if logger.LogErr(err) {
			return
		}
	} else {
		err = json.Unmarshal(weatherJson, &weather)
		if logger.LogErr(err) {
			return
		}
	}

	c.HTML(http.StatusOK, "update", gin.H{
		"Temp":      weather.Main.Temp,
		"Feels":     weather.Main.FeelsLike,
		"Humidity":  weather.Main.Humidity,
		"WindSpeed": weather.Wind.Speed,
	})
}
