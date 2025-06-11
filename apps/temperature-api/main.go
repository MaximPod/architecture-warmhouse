package main

import (
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// сенсор температуры
type Sensor struct {
	Id        string  `json:"id"`        // Идентификатор датчика
	Location  string  `json:"location"`  // Местоположение датчика
	Value     float64 `json:"value"`     // Значение температуры
	Unit      string  `json:"unit"`      // Единица измерения
	Timestamp string  `json:"timestamp"` // Время измерения
	Status    string  `json:"status"`    // Статус датчика
}

func main() {
	// Создаем роутер Gin
	r := gin.Default()

	// Создаем новый генератор случайных чисел
	rng := rand.New(rand.NewSource(rand.Int63()))

	// endpoint temperature by location
	r.GET("/temperature", func(c *gin.Context) {
		// Получаем параметр location из query string
		location := c.Query("location")

		id := getIdByLocation(location)

		sensor := &Sensor{
			Id:        id,
			Location:  location,
			Value:     getRandomTemperature(rng),
			Unit:      "Celsius",
			Timestamp: time.Now().Format(time.RFC3339),
			Status:    "ok",
		}

		// Возвращаем ответ в формате JSON
		c.JSON(http.StatusOK, sensor)
	})

	// endpoint temperature by id
	r.GET("/temperature/:id", func(c *gin.Context) {
		id := c.Param("id")

		location := getLocationById(id)

		sensor := &Sensor{
			Id:        id,
			Location:  location,
			Value:     getRandomTemperature(rng),
			Unit:      "Celsius",
			Timestamp: time.Now().Format(time.RFC3339),
			Status:    "ok",
		}

		// Возвращаем ответ в формате JSON
		c.JSON(http.StatusOK, sensor)
	})

	// Запускаем сервер на порту 8080
	r.Run(":8081")
}

// getIdByLocation возвращает идентификатор датчика по местоположению
func getIdByLocation(location string) (sensorID string) {
	switch location {
	case "Living Room":
		sensorID = "1"
	case "Bedroom":
		sensorID = "2"
	case "Kitchen":
		sensorID = "3"
	default:
		sensorID = "0"
	}

	return sensorID
}

// getLocationById возвращает местоположение датчика по идентификатору
func getLocationById(sensorID string) (location string) {
	switch sensorID {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		location = "Unknown"
	}

	return location
}

// Генерация случайной температуры от -30 до +40 градусов c точностью до двух знаков после запятой
func getRandomTemperature(rand *rand.Rand) float64 {
	
	return roundFloat(rand.Float64()*70-30, 2)
}

// roundFloat функция округления числа до заданной точности
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	
	return math.Round(val*ratio) / ratio
}
