package main

import (
	"encoding/json"
	"final/cmd"
	"final/cmd/middleware"
	"final/cmd/models"
	"final/weather"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			log.Println("Echo middleware was called")
			return next(ctx)
		}
	})

	router.GET("api/lists/:id/tasks", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "GET")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

		param := ctx.Param("id")
		log.Println(param)
		payload := middleware.GetAllTasks(param)

		return json.NewEncoder(ctx.Response()).Encode(payload)
	})

	router.POST("api/lists/:id/tasks", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var task models.Task
		id := ctx.Param("id")
		i, _ := strconv.ParseInt(id, 10, 64)
		task.ListId = i
		_ = json.NewDecoder(ctx.Request().Body).Decode(&task)
		middleware.InsertTask(task)

		return json.NewEncoder(ctx.Response()).Encode(task)
	})

	router.PATCH("api/tasks/:id", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Content-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "PATCH")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

		id := ctx.Param("id")
		i, _ := strconv.ParseInt(id, 10, 64)
		middleware.UpdateTask(i)

		return json.NewEncoder(ctx.Response()).Encode(i)
	})

	router.DELETE("api/tasks/:id", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "DELETE")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

		param := ctx.Param("id")
		err := middleware.DeleteOneTask(param)
		return err
	})

	router.GET("api/lists", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "DELETE")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
		payload := middleware.GetAllLists()
		return json.NewEncoder(ctx.Response()).Encode(payload)
	})

	router.POST("api/lists", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var list models.List
		_ = json.NewDecoder(ctx.Request().Body).Decode(&list)
		middleware.InsertList(list)

		return json.NewEncoder(ctx.Response()).Encode(list)
	})

	router.DELETE("api/lists/:id", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "DELETE")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

		param := ctx.Param("id")
		err := middleware.DeleteOneList(param)
		return err
	})

	router.GET("api/weather", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response().Header().Set("Access-Control-Allow-Methods", "GET")
		ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
		apiConfig, err := weather.LoadApiConfig(".apiConfig")
		if err != nil {
			log.Println("Error loading api config key!")
		}

		latitude := ctx.Request().Header.Get("lat")
		longitude := ctx.Request().Header.Get("lon")

		var resp *http.Response
		if latitude != "" && longitude != "" {
			resp, _ = http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + latitude + "&lon=" + longitude + "&appid=" + apiConfig.OpenWeatherMapApiKey)
		} else {
			resp, _ = http.Get("https://api.openweathermap.org/data/2.5/weather?lat=41.997345&lon=21.427996&appid=" + apiConfig.OpenWeatherMapApiKey)
		}

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var weatherData weather.WeatherData
		json.Unmarshal(responseData, &weatherData)

		var weatherInfo weather.WeatherInfo
		weatherInfo.City = weatherData.City
		weatherInfo.Description = weatherData.Weather[0].Description
		weatherInfo.FormatedTemp = fmt.Sprintf("%f", weatherData.FormatedTemp.Temp-273.15)

		return ctx.JSON(200, weatherInfo)
	})

	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
