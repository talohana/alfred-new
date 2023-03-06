package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	shortcutsURL        = "https://whats.new/shortcuts"
	rowSelector         = ".data-table__row"
	descriptionSelector = ".data-table__cell--description"
)

type Icon struct {
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
}

type Item struct {
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Arg      string `json:"arg,omitempty"`
	Icon     *Icon  `json:"icon,omitempty"`
}

type Workflow struct {
	Items []*Item `json:"items"`
}

func main() {
	res, err := http.Get(shortcutsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var workflow Workflow

	doc.Find(rowSelector).Each(func(_ int, row *goquery.Selection) {
		description := row.Find(descriptionSelector).Text()
		href, _ := row.Find("a").Attr("href")
		workflow.Items = append(workflow.Items, &Item{
			Title:    strings.TrimSpace(description),
			Subtitle: strings.TrimPrefix(href, "https://"),
			Arg:      href,
			Icon: &Icon{
				Path: "./icon.svg",
			},
		})
	})

	data, err := json.MarshalIndent(workflow, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", data)
}
