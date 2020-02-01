package main

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	router := gin.Default()

	router.GET("/single", SumSingleAPI)
	router.GET("/double", SumDoubleAPI)
	router.GET("/multiple", SumMultipleAPI)
	router.Run(":3000")
}

func SumSingleAPI (c *gin.Context) {
	numbers := []int {9, 5, 4, 1, 3, 2, 5, 10}
	c.JSON(200, gin.H{
		"sum": Sum(numbers),
	})
}

func SumDoubleAPI (c *gin.Context) {
	numbers := []int {9, 5, 4, 1, 3, 2, 5, 10}
	c.JSON(200, gin.H{
		"sum": SumDoubleThreading(numbers),
	})
}

func SumMultipleAPI (c *gin.Context) {
	numbers := []int {9, 5, 4, 1, 3, 2, 5, 10}
	c.JSON(200, gin.H{
		"sum": SumMultipleThreading(numbers),
	})
}

func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SumDoubleThreading(numbers []int) int {
	mid := len(numbers) / 2

	ch := make(chan int)
	go func() { ch <- Sum(numbers[:mid]) }()
	go func() { ch <- Sum(numbers[mid:]) }()

	total := <-ch + <-ch
	return total
}

func SumMultipleThreading(numbers []int) int {
	nCPU := runtime.NumCPU()
	nNum := len(numbers)

	ch := make(chan int)
	for i := 0; i < nCPU; i++ {
		from := i * nNum / nCPU
		to := (i + 1) * nNum / nCPU
		go func() { ch <- Sum(numbers[from:to]) }()
	}

	total := 0
	for i := 0; i < nCPU; i++ {
		total += <-ch
	}
	return total
}
