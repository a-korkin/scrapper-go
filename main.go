package main

import (
	"fmt"
	"log"
	// "os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
	"time"
)

func main() {
	port := 4444
	selenium.SetDebug(true)
	opts := []selenium.ServiceOption{
		// selenium.Output(os.Stderr),
	}
	service, err := selenium.NewGeckoDriverService(
		"./drivers/geckodriver", port, opts...)
	if err != nil {
		log.Fatalf("failed to create selenium driver: %s", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddFirefox(firefox.Capabilities{
		Binary: "/usr/bin/firefox",
		Args: []string{
			"-private-window",
			// "--headless",
		}})

	urlPrefix := fmt.Sprintf("http://localhost:%d", port)
	driver, err := selenium.NewRemote(caps, urlPrefix)
	// driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatalf("failed to create new remote: %s", err)
	}
	defer driver.Quit()
	driver.MaximizeWindow("")

	err = driver.Get("https://megamarket.ru/")
	if err != nil {
		log.Fatalf("failed to get page: %s", err)
	}

	driver.WaitWithTimeoutAndInterval(getSearchInput, 1000, 1000)
	driver.WaitWithTimeoutAndInterval(submitButton, 1000, 1000)

	time.Sleep(10 * time.Second)
}

func getSearchInput(driver selenium.WebDriver) (bool, error) {
	el, _ := driver.FindElement(selenium.ByCSSSelector, "input.search-field-input")
	shown, err := el.IsDisplayed()
	if shown {
		el.SendKeys("macbook air")
	}
	return shown, err
}

func submitButton(driver selenium.WebDriver) (bool, error) {
	el, _ := driver.FindElement(selenium.ByCSSSelector, "button.header-search-form__search-button")
	shown, err := el.IsDisplayed()
	if shown {
		el.Click()
	}
	return shown, err
}
