package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var PORT string = ":8000"

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	//server.SetTrustedProxies(nil)
	server.Use(cors.Default())
	// GET** <your-domain.com>/api/classify-number?number=371
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	server.GET("/api/classify-number", classifyNumber)
	log.Printf("running on http://localhost%s/api/classify-number\n", PORT)
	server.Run(PORT)
}

func classifyNumber(ctx *gin.Context) {
	// get the query param i.e ?number=int
	query := ctx.Query("number")
	// cache control
	ctx.Header("Cache-Control", "public, max-age=31536000, immutable")
	ctx.Header("Expires", "Thu, 31 Dec 9999 23:59:59 GMT")
	// handle blank query param
	if query == "" {
		ctx.AbortWithStatusJSON(400, gin.H{
			"number": "alphabet",
			"error":  true,
			// "error":  "query param cannot be blank",
		})
		return
	}
	// handle data types which are not intergers
	numQuery, err := strconv.Atoi(query)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"number": "alphabet",
			"error":  true,
			// "error":  "query param cannot be converted to an interger",
		})
		return
	}
	// assign the number to check to the Number struct
	numToCheck := &numClass{
		Number: numQuery,
	}
	parseNum(numToCheck)
	// return properties of the number in a JSON format
	ctx.JSON(200, numToCheck)
}
