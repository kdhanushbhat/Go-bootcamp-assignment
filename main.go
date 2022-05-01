package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	AUthor string
}

var quotes []quote = []quote{
	{"do well all you do", "dafoodils"},
	{"leed well", "svit"},
	{"You cannot believe in god unitl u believe in yourself", "swami vivekananda"},
	{"Facts are many truth is one", "R Tagore"},
}

var html = fmt.Sprintf(`<h1>K DHANUSH BHAT</h1>
		<h2>Quotes</h2>
		<a href="http://localhost:20200/quotes/random">Get Random Quote</a>
		<p><span style="font-weight:600">Title : </span>` + quotes[0].Title + `<br><span style="font-weight:600">Value : </span>` + quotes[0].AUthor + `</p>
		<p><span style="font-weight:600">Title : </span>` + quotes[1].Title + `<br><span style="font-weight:600">Value : </span>` + quotes[1].AUthor + `</p>
		<p><span style="font-weight:600">Title : </span>` + quotes[2].Title + `<br><span style="font-weight:600">Value : </span>` + quotes[2].AUthor + `</p>
		<p><span style="font-weight:600">Title : </span>` + quotes[3].Title + `<br><span style="font-weight:600">Value : </span>` + quotes[3].AUthor + `</p>`)

func main() {
	api := echo.New()

	api.GET("/quotes", getQuotes)
	api.GET("/quotes/random", getRandomQuote)
	api.POST("/quotes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, quotes)
	})
	api.PUT("/quotes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, quotes)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}
	api.Start(":20200")
}

func getQuotes(c echo.Context) error {
	return c.HTML(http.StatusOK, html)
}

func getRandomQuote(c echo.Context) error {
	index := rand.Intn(len(quotes))
	randhtml := fmt.Sprintf(`<h1>K DHANUSH BHAT</h1>
	<h2>Random Quote</h2>
	<a href="http://localhost:20200/quotes">Back</a>
	<p><span style="font-weight:600">Title : </span>` + quotes[index].Title + `<br><br><span style="font-weight:600">Value : </span>` + quotes[index].AUthor + `</p>`)
	return c.HTML(http.StatusOK, randhtml)

}
