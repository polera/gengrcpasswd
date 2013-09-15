package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	var doc *goquery.Document
	var e error

	if doc, e = goquery.NewDocument("https://www.grc.com/passwords.htm"); e != nil {
		panic(e.Error())
	}

	doc.Find("table [bgcolor=\"#FF0000\"] > tbody > tr > td > table > tbody > tr > td").Each(func(i int, s *goquery.Selection) {
		var text string

		text = s.Find("font").Text()
		fmt.Printf("%s\n", text)
	})
}
