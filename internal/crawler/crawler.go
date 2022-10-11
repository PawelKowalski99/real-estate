package crawler

import (
	"context"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	"time"
)

func NewCrawler() (browser *rod.Browser) {

	ctx := context.Background()
	idleCounter := utils.NewIdleCounter(time.Second)

	go func() {
		idleCounter.Add()
		// This example is to launch a browser remotely, not connect to a running browser remotely,
		// to connect to a running browser check the "../connect-browser" example.
		// Rod provides a docker image for beginers, run the below to start a launcher.Manager:
		//
		//     docker run -p 7317:7317 ghcr.io/go-rod/rod
		//
		// For more information, check the doc of launcher.Manager
		var l *launcher.Launcher
		var err error
		for true {
			l, err = launcher.NewManaged("ws://rod:7317")
			if err == nil {
				break
			}
			time.Sleep(time.Second)
			fmt.Println("Czekam na browser", err.Error())
		}

		// You can also set any flag remotely before you launch the remote browser.
		// Available flags: https://peter.sh/experiments/chromium-command-line-switches
		l.Set("disable-gpu").Delete("disable-gpu")

		// Launch with headful mode
		l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")

		browser = rod.New().Client(l.MustClient()).MustConnect()

		// You may want to start a server to watch the screenshots of the remote browser.
		launcher.Open(browser.ServeMonitor(""))

		fmt.Println(
			browser.MustPage("https://mdn.dev/").MustEval("() => document.title"),
		)
		idleCounter.Done()
	}()
	idleCounter.Wait(ctx)


	return browser
}


