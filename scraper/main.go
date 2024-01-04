package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"context"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

const (
	titleClassesRaw   = `x193iq5w xeuugli x13faqbe x1vvkbs x1xmvt09 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x14z4hjw x3x7a5m xngnso2 x1qb5hxa x1xlr1w8 xzsf02u`
	mileageClassesRaw = `x193iq5w xeuugli x13faqbe x1vvkbs x1xmvt09 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x xudqn12 x3x7a5m x6prxxf xvq8zen xo1l8bm xzsf02u`
	priceClassesRaw   = `x193iq5w xeuugli x13faqbe x1vvkbs x1xmvt09 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x xudqn12 x676frb x1lkfr7t x1lbecb7 x1s688f xzsf02u`
)

// A struct to represent a Listing, will add more fields later
type Listing struct {
	Url     string
	Price   string
	Title   string
	Mileage string
}

func main() {
	url := "https://www.facebook.com/marketplace/item/306059595765532"
	saveHTML(url)

	mileageElement := htmlElement{
		classes:   fixClasses(mileageClassesRaw),
		occurance: 1,
	}

	priceElement := htmlElement{
		classes:   fixClasses(priceClassesRaw),
		occurance: 0,
	}

	titleElement := htmlElement{
		classes:   fixClasses(titleClassesRaw),
		occurance: 1,
	}

	mileageString, err := getInnerHTML(mileageElement)
	if err != nil {
		panic(err)
	}
	priceString, err := getInnerHTML(priceElement)
	if err != nil {
		panic(err)
	}
	titleString, err := getInnerHTML(titleElement)
	if err != nil {
		panic(err)
	}

	fmt.Println(titleString)
	fmt.Println(priceString)
	fmt.Println(mileageString)
}

func getInnerHTML(e htmlElement) (string, error) {
	// Open the file
	f, err := os.Open("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse the file
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	// Find the element with the two classes
	s := doc.Find(e.classes).Eq(e.occurance)
	html, err := s.Html()
	if err != nil {
		return "", err
	}
	return html, nil
}

func saveHTML(url string) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// This will hold the HTML from the page
	var html string

	// Run tasks
	err := chromedp.Run(ctx,
		// Navigate to the page you want to screenshot
		chromedp.Navigate(url),
		// Get the HTML of the root node
		chromedp.OuterHTML("html", &html),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Save the HTML to a file
	err = os.WriteFile("output.html", []byte(html), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTML has been saved to output.html")
}

func fixClasses(classes string) string {
	return "span." + strings.ReplaceAll(classes, " ", ".")
}

type htmlElement struct {
	classes   string
	occurance int
}
