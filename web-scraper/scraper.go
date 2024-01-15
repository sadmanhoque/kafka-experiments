package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	scrapeUrl := "https://applications.crtc.gc.ca/portail-portal/eng/listes-lists/broadcasting-services-list/9?_ga=2.149691062.1831938506.1700653356-630674744.1700653356"

	c := colly.NewCollector(colly.AllowedDomains("applications.crtc.gc.ca"))

	c.OnHTML("h1[id=wb-cont]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit(scrapeUrl)
}
