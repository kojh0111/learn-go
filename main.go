package main

import (
	"os"
	"strings"
	"time"

	"github.com/kojh0111/learngo/scraper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scraper.CleanString(c.FormValue("term")))
	t := time.Now()
	scraper.Scrap(term)
	return c.Attachment("jobs.csv", term+t.Format(time.ANSIC)+".csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrap", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
