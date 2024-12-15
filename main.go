package main

import (
	"fmt"
	"log"
	// "os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
	// "github.com/tebeka/selenium/chrome"
)

func main() {
	port := 4444
	// service, err := selenium.NewChromeDriverService("./drivers/chromedriver", port)
	service, err := selenium.NewGeckoDriverService("./drivers/geckodriver", port)
	if err != nil {
		log.Fatalf("failed to create selenium driver: %s", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddFirefox(firefox.Capabilities{
		Binary: "/usr/bin/firefox",
		Args: []string{
			"--window-size=640,480",
			// "--headless",
		}})
	// caps.AddChrome(chrome.Capabilities{
	// 	Path: "/snap/bin/chromium",
	// 	Args: []string{
	// 		// "--headless",
	// 	}})

	selenium.SetDebug(true)

	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d", port))
	// driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatalf("failed to create new remote: %s", err)
	}
	defer driver.Quit()

	err = driver.Get("https://www.google.com")
	if err != nil {
		log.Fatalf("failed to get page: %s", err)
	}

	html, err := driver.PageSource()
	if err != nil {
		log.Fatalf("failed to get page source: %s", err)
	}
	fmt.Println(html)
}
