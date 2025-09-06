package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/playwright-community/playwright-go"
)

// one time call:  install chrome
func init() {
	// Install browsers on first run; comment out after installation to speed up startup
	if err := playwright.Install(); err != nil {
		log.Fatal(err)
	}
}

// finanzfluss return a string like Strg+A and Strg+C from the website
func finanzfluss(isin string) string {

	url := "https://www.finanzfluss.de/informer/etf/" + isin

	// Start Playwright
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	// Launch Chromium headless
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	// New context and page
	ctx, err := browser.NewContext()
	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}
	page, err := ctx.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Navigate and wait
	_, err = page.Goto(url)
	if err != nil {
		log.Fatalf("navigation failed: %v", err)
	}

	// Wait for fonts to be ready so innerText reflects final layout
	if _, err := page.Evaluate(`() => (document.fonts && document.fonts.ready) ? document.fonts.ready : true`); err != nil {
		// Non-fatal if fonts API is unavailable
	}

	// Extract visible text (no HTML)
	visible, err := page.Evaluate(`() => document.body ? document.body.innerText : ""`)
	if err != nil {
		log.Fatalf("extract innerText failed: %v", err)
	}
	text := strings.TrimSpace(visible.(string))

	// Print plain text to stdout
	return text
}

// countryWeighting use the website text to extract the country weights
func countryWeighting(text string) map[string]string {

	// Use a scanner to iterate line by line
	sc := bufio.NewScanner(strings.NewReader(text))

	on := false
	values := make([]string, 0, 50)

	// scan all lines
	// find text between "Länder" and "Regionen"
	for sc.Scan() {
		line := sc.Text()

		// skip invalid lines
		if len(line) <= 2 {
			continue
		}
		// start
		if line == "Länder" {
			on = true
			continue
		}
		// stop
		if line == "Regionen" {
			on = false
			continue
		}
		// collect values
		if on {
			//println(line) // DEBUG !!!
			values = append(values, line)
		}
	}

	// process data
	//    [0]: Vereinigte Staaten
	//    [1]: 17,37%
	//    [2]: Australien
	//    [3]: 13,69%
	//    [4]: Hongkong
	//    [5]: ...
	ret := make(map[string]string)
	for i := 1; i < len(values); i += 2 {
		key := values[i-1]
		key = strings.TrimSpace(key)

		value := values[i]
		value = strings.TrimSpace(value)
		value = strings.ReplaceAll(value, "%", "")
		value = strings.ReplaceAll(value, ",", ".")
		_, err := strconv.ParseFloat(value, 64) // test float
		if err != nil {
			panic(err)
		}

		// FIX COUNTRIES
		if key == "Türkei" {
			key = "Turkei"
		}
		if key == "Saudi-Arabien" {
			key = "Saudi Arabien"
		}
		if key == "Hongkong" {
			key = "Hong Kong"
		}
		if key == "Vereinigtes Königreich" {
			key = "Großbritannien"
		}

		//fmt.Printf("%s: %f\n", key, val)  // DEBUG !!!!
		ret[key] = value
	}

	return ret
}
