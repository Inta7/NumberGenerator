package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/random", func(c *gin.Context) {
		// Получение параметров запроса
		min := c.Query("min")
		max := c.Query("max")
		count := c.Query("count")

		// Проверка наличия параметров
		if min == "" || max == "" || count == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
			return
		}

		// Конвертирование параметров в числа
		minInt, err := strconv.Atoi(min)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter: min"})
			return
		}
		maxInt, err := strconv.Atoi(max)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter: max"})
			return
		}
		countInt, err := strconv.Atoi(count)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter: count"})
			return
		}

		// Генерация случайных чисел
		numbers := make([]int, countInt)
		for i := 0; i < countInt; i++ {
			numbers[i] = rand.Intn(maxInt-minInt) + minInt
		}

		// Отправка результатов
		c.JSON(http.StatusOK, gin.H{"numbers": numbers})
	})

	r.Run(":8080")
}

//http://localhost:8080/random?min=1&max=100&count=5
