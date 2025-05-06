package main

import (
	"container/list"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
)

func getGaugeProgress(html string) int {

	if strings.Contains(html, "gauge-0") {
		return 0
	}

	if strings.Contains(html, "gauge-1") {
		return 1
	}

	if strings.Contains(html, "gauge-2") {
		return 2
	}

	if strings.Contains(html, "gauge-3") {
		return 3
	}

	if strings.Contains(html, "gauge-4") {
		return 4
	}

	if strings.Contains(html, "gauge-5") {
		return 5
	}

	if strings.Contains(html, "gauge-6") {
		return 6
	}

	if strings.Contains(html, "gauge-7") {
		return 7
	}

	return 8

}

func getReportedTimeStamp(html string) string {
	s := strings.Split(html, "ldst_strftime(")
	if len(s) != 2 {
		log.Fatal("Somehow, the time is not valid...")
	}

	s = strings.Split(s[1], ",")

	return s[0]
}

func processWebsiteData(doc *goquery.Document) *list.List {
	DataCenters := []string{"Aether", "Crystal", "Dynamis", "Primal", "Chaos", "Light", "Materia", "Elemental", "Gaia", "Mana", "Meteor"}
	Reports := list.New()

	ts, _ := doc.Find(".cosmic__report__update").Html()
	reportedTimeStamp := getReportedTimeStamp(ts)

	for _, dataCenter := range DataCenters {

		doc.Find("#" + dataCenter + " .cosmic__report__card").Each(func(i int, s *goquery.Selection) {
			info := strings.Fields(s.Text())
			gaugeProgressStr, _ := s.Find(".cosmic__report__status__progress").Html()
			gradeProgressStr, _ := s.Find(".cosmic__report__grade__level p").Html()
			gradeProgress := strings.TrimSpace(gradeProgressStr)

			server := Report{uuid.New().String(), dataCenter, info[0], gradeProgress, getGaugeProgress(gaugeProgressStr), reportedTimeStamp, strconv.FormatInt(time.Now().Unix(), 10)}

			fmt.Println(server)
			Reports.PushBack(server)
		})
	}

	return Reports
}

func getWebsiteData() *goquery.Document {
	res, err := http.Get("https://na.finalfantasyxiv.com/lodestone/cosmic_exploration/report/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("Not a 200...server down?")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func main() {
	doc := getWebsiteData()
	data := processWebsiteData(doc)

	fmt.Println(data)
}
