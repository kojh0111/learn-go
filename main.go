package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var BaseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)
}

func getPages() int {
	res, err := http.Get(BaseURL)
	checkErr(err)
	checkRes(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	pages := 0

	doc.Find("#searchCountPages").Each(func(i int, s *goquery.Selection) {
		countString := strings.Split(s.Text(), " 결과 ")
		splitCount := strings.Replace(strings.Split(countString[len(countString)-1], "건")[0], ",", "", -1)
		countJob, _ := strconv.Atoi(splitCount)
		pages = countJob / 50
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkRes(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
