package main

import (
	"fmt"
	"strings"

	"github.com/kojh0111/learngo/scraper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scraper.CleanString(c.FormValue("term")))
	fmt.Println(term)
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrap", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
