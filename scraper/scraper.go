package scraper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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

// Scrap jobs by a term
func Scrap(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	totalPages := getPages(baseURL)
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(baseURL, i)
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(url string, page int) []extractedJob {
	var jobs []extractedJob
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkRes(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func getPages(url string) int {
	res, err := http.Get(url)
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

func extractJob(card *goquery.Selection) extractedJob {
	link, _ := card.Attr("href")
	title := CleanString(card.Find("h2 > span").Text())
	company := card.Find(".companyName").Text()
	location := card.Find(".companyLocation").Text()
	salary := card.Find(".salary-snippet > span").Text()
	summary := CleanString(card.Find(".job-snippet").Text())
	return extractedJob{
		link:     link,
		title:    title,
		company:  company,
		location: location,
		salary:   salary,
		summary:  summary}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	w.Comma = '|'
	defer w.Flush()

	headers := []string{"Link", "Title", "Company", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com" + job.link, job.title, job.company, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
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

// CleanString cleans a string
func CleanString(str string) string {
	return strings.ReplaceAll(strings.Join(strings.Fields(strings.TrimSpace(str)), " "), "|", ",")
}
