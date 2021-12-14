package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	link     string
	title    string
	company  string
	location string
	salary   string
	summary  string
}

// BaseURL - job search by pyhon in indeed
var BaseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()
	for i := 0; i <= totalPages; i++ {
		getPage(i)
	}
	// getPage(0)
}

func getPage(page int) {
	pageURL := BaseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkRes(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		link, _ := card.Attr("href")
		title := card.Find("h2 > span").Text()
		company := card.Find(".companyName").Text()
		location := card.Find(".companyLocation").Text()
		salary := card.Find(".salary-snippet > span").Text()
		summary := card.Find(".job-snippet").Text()
		job := extractedJob{
			link:     link,
			title:    title,
			company:  company,
			location: location,
			salary:   salary,
			summary:  summary}
		fmt.Println(job.link)
	})
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
