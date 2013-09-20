package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
    "strings"
	"time"
)

type PasswordSets struct {
	Set []PasswordSet
}

type PasswordSet struct {
	Username string
	Password string
}

func getRand() int {
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := rand.Intn(62)
	return randomNum
}

func generateString(length int, population string) string {
    tmpString := ""
    for i:= 0; i < length; i++ {
       tmpString += string(population[getRand()])
    }
    return tmpString
}

var (
	numSets    = flag.Int("count", 1, "Number of username:password sets that you require.")
	userLength = flag.Int("userlen", 8, "Desired username length")
	passLength = flag.Int("passwordlen", 8, "Desired password length")
)


func main() {
	flag.Parse()
	var doc *goquery.Document
	var e error
	fields := make([]string, 0)
	fmt.Println("Contacting grc.com...")
	if doc, e = goquery.NewDocument("https://www.grc.com/passwords.htm"); e != nil {
		fmt.Println("Error contacting grc.com!")
		panic(e.Error())
	}
	fmt.Println("Done.\nGenerating password(s)")
	doc.Find("table [bgcolor=\"#FF0000\"] > tbody > tr > td > table > tbody > tr > td").Each(func(i int, s *goquery.Selection) {
		var text string

		text = s.Find("font").Text()
		fields = append(fields, text)
	})

	asciiChars := fields[1]
	alphaNumChars := fields[2]

	ps := PasswordSets{}

	for i := 0; i < *numSets; i++ {
		set := PasswordSet{"", ""}
		set.Username = generateString(*userLength, alphaNumChars)
		set.Password = generateString(*passLength, asciiChars)
		ps.Set = append(ps.Set, set)
	}
    fmt.Printf("\n\nUSERNAME %s: PASSWORD\n", strings.Repeat(" ",*userLength-8))
	for _, passwordSet := range ps.Set {
		fmt.Printf("%s : %s\n", passwordSet.Username, passwordSet.Password)
	}

}
