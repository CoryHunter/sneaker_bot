package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
)

func main() {
	const (
		seleniumPath    = "/Users/coryhunter/go/src/github.com/tebeka/selenium/vendor/selenium-server.jar"
		geckoDriverPath = "/Users/coryhunter/go/src/github.com/tebeka/selenium/vendor/chromedriver"
		port            = 3000
	)

	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(geckoDriverPath),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(false)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	if err := wd.Get("https://www.nike.com"); err != nil {
		panic(err)
	}
	showElement, err := wd.FindElement(selenium.ByXPATH,"/html/body/div[1]/div[3]/header/div/div[1]/div[2]/nav/div[2]/ul/li[1]")
	if err != nil {
		panic(err)
	}
	if err := showElement.Click(); err != nil {
		panic(err)
	}
}