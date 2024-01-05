package scraper

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"context"

	"github.com/JaiiR320/carlistingsaver/types"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

const (
	titleClassesRaw       = `x193iq5w xeuugli x13faqbe x1vvkbs x1xmvt09 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x x14z4hjw x3x7a5m xngnso2 x1qb5hxa x1xlr1w8 xzsf02u`
	mileageClassesRaw     = `x193iq5w xeuugli x13faqbe x1vvkbs x1xmvt09 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x xudqn12 x3x7a5m x6prxxf xvq8zen xo1l8bm xzsf02u`
	priceClassesRaw       = `x193iq5w xeuugli x13faqbe x1vvkbs x1xmvt09 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x xudqn12 x676frb x1lkfr7t x1lbecb7 x1s688f xzsf02u`
	descriptionClassesRaw = `x193iq5w xeuugli x13faqbe x1vvkbs x1xmvt09 x1lliihq x1s928wv xhkezso x1gmr53x x1cpjm7i x1fgarty x1943h6x xudqn12 x3x7a5m x6prxxf xvq8zen xo1l8bm xzsf02u`
	descDivRaw            = `xod5an3`
	descButton            = `x1i10hfl xjbqb8w x6umtig x1b1mbwd xaqea5y xav7gou x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1o1ewxj x3x9cwd x1e5q0jg x13rtm0m x1n2onr6 x87ps6o x1lku1pv x1a2a7pz`
	closeBtn              = `x1i10hfl x6umtig x1b1mbwd xaqea5y xav7gou x1ypdohk xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r x16tdsg8 x1hl2dhg xggy1nq x87ps6o x1lku1pv x1a2a7pz x6s0dn4 x14yjl9h xudhj91 x18nykt9 xww2gxu x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x78zum5 xl56j7k xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 xc9qbxq x14qfxbe x1qhmfi1`
)

// Get a listing from a given url.
// Returns a listing object with the data from the url
func GetListing(url string) types.Listing {
	saveHTML(url)

	mileageString, err := getInnerHTML(mileageClassesRaw, 1)
	if err != nil {
		log.Fatal(err)
	}
	priceString, err := getInnerHTML(priceClassesRaw, 0)
	if err != nil {
		log.Fatal(err)
	}
	titleString, err := getInnerHTML(titleClassesRaw, 1)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Remove("output.html"); err != nil {
		log.Fatal(err)
	}
	log.Println("output.html has been deleted")

	priceString = strings.Replace(priceString, "$", "", 1)
	priceString = strings.Replace(priceString, ",", "", 1)
	price, err := strconv.Atoi(priceString)
	if err != nil {
		log.Fatal(err)
	}

	// Use a regular expression to match all digits and commas
	re := regexp.MustCompile("[0-9,]+")
	matches := re.FindAllString(mileageString, -1)

	if len(matches) == 0 {
		log.Fatal("No numbers found in the string")
	}

	// Remove the comma
	mileageString = strings.Replace(matches[0], ",", "", -1)

	// Convert the string to an integer
	mileage, err := strconv.Atoi(mileageString)
	if err != nil {
		log.Fatal(err)
	}

	return *types.NewListing(url, price, titleString, mileage)
}

// gets the inner HTML of an element with the given classes and index
func getInnerHTML(rawClassString string, index int) (string, error) {
	classes := fixClasses(rawClassString)
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
	s := doc.Find(classes).Eq(index)
	html, err := s.Html()
	if err != nil {
		return "", err
	}
	return html, nil
}

// saves the page's html to a file
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

	log.Println("HTML has been saved to output.html")
}

// converts raw list of classes to a format that goquery can use
func fixClasses(classes string) string {
	return "." + strings.ReplaceAll(classes, " ", ".")
}
