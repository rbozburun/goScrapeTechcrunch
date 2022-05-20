package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func writeFile(data, filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	check(err)

	file.WriteString(data)
}

func main() {
	url := "https://techcrunch.com/"

	response, error := http.Get(url)
	defer response.Body.Close()

	check(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	check(err)

	//Create CSV file
	file, err := os.Create("posts.csv")
	check(err)

	writer := csv.NewWriter(file)

	doc.Find("div.river").Find("div.post-block").Each(func(index int, item *goquery.Selection) {
		h2 := item.Find("h2")
		title := strings.TrimSpace(h2.Text()) // Get rid of the spaces
		url, _ := h2.Find("a").Attr("href")   // Attr returns item, bool (attr value exists or not)
		excerpt := strings.TrimSpace(item.Find("div.post-block__content").Text())

		posts := []string{title, url, excerpt}

		//Write data to csv file
		writer.Write(posts)
	})
	writer.Flush()

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
