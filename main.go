package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
	"tugas3/model"

	"encoding/json"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {

		for {

			data, err := os.Open("data.json")
			if err != nil {
				panic(err)
			}

			byteValue, _ := io.ReadAll(data)
			var result model.Data
			err = json.Unmarshal(byteValue, &result)
			if err != nil {
				panic(err)
			}

			rand.Seed(time.Now().UnixNano())
			waterNew := rand.Intn(100)
			windNew := rand.Intn(100)

			result.Status.Water = strconv.Itoa(waterNew)
			result.Status.Wind = strconv.Itoa(windNew)

			file, err := os.Create("data.json")
			if err != nil {
				panic(err)
			}

			jsonData, err := json.Marshal(result)
			if err != nil {
				panic(err)
			}

			_, err = file.Write(jsonData)
			if err != nil {
				panic(err)
			}

			data.Close()

			time.Sleep(15 * time.Second)

		}

	}()

	router := gin.Default()
	router.GET("/data", func(c *gin.Context) {
		data, err := os.Open("data.json")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer data.Close()

		byteValue, _ := io.ReadAll(data)
		var result model.Data
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			log.Fatal(err.Error())
		}

		var resultData model.Result

		water, _ := strconv.Atoi(result.Status.Water)
		wind, _ := strconv.Atoi(result.Status.Wind)

		switch {
		case water < 5:
			resultData.Status = "Aman"
		case water >= 6 && water <= 8:
			resultData.Status = "Siaga"
		case water > 8:
			resultData.Status = "Bahaya"
		case wind < 6:
			resultData.Status = "Aman"
		case wind >= 7 && wind <= 15:
			resultData.Status = "Siaga"
		case wind > 15:
			resultData.Status = "Bahaya"
		}

		resultData.Weather.Status.Water = result.Status.Water + " m"
		resultData.Weather.Status.Wind = result.Status.Wind + " m/s"

		c.JSON(200, resultData)
	})

	router.Run(":8000")
}
