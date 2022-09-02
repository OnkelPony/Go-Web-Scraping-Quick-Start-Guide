package main

import (
	"github.com/tebeka/selenium"
)

func main() {

	// The paths to these binaries will be different on your machine!

	const (
		seleniumPath    = "/home/gio/go/pkg/mod/github.com/tebeka/selenium@v0.9.9/vendor/selenium-server.jar"
		geckoDriverPath = "/home/gio/go/pkg/mod/github.com/tebeka/selenium@v0.9.9/vendor/geckodriver"
	)

	service, err := selenium.NewSeleniumService(
		seleniumPath,
		8080,
		selenium.GeckoDriver(geckoDriverPath))

	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, "http://localhost:8080/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	err = wd.Get("https://www.packtpub.com/product/mastering-go/9781801079310")
	if err != nil {
		panic(err)
	}

	var elems []selenium.WebElement
	wd.Wait(func(wd2 selenium.WebDriver) (bool, error) {
		elems, err = wd.FindElements(selenium.ByCSSSelector, "div.product-reviews-review div.review-body")
		if err != nil {
			return false, err
		} else {
			return len(elems) > 0, nil
		}
	})

	for _, review := range elems {
		body, err := review.Text()
		if err != nil {
			panic(err)
		}
		println(body)
	}
}
