package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"time"

	mr "math/rand"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
	"github.com/romainmenke/stylesheet-randomizer/randomstylesheet"
)

func main() {
	mr.Seed(time.Now().UnixNano())

	for i := 0; i < 200; i++ {
		capture(i, int64(mr.Int()))
	}
}

func capture(iteration int, seed int64) {
	chromedpCtx, chromedpCancel := chromedp.NewContext(context.Background())
	defer chromedpCancel()

	start := time.Now()
	var buf []byte

	err := chromedp.Run(chromedpCtx,
		chromedp.ActionFunc(func(actionCtx context.Context) error {
			actionErr := network.ClearBrowserCache().Do(actionCtx)
			if actionErr != nil {
				return actionErr
			}

			actionErr = network.ClearBrowserCookies().Do(actionCtx)
			if actionErr != nil {
				return actionErr
			}

			actionErr = storage.ClearDataForOrigin("localhost:4567", "all").Do(actionCtx)
			if actionErr != nil {
				return actionErr
			}

			return nil
		}),
		fullScreenshot("http://localhost:4567?seed="+fmt.Sprint(seed), 90, &buf),
	)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("data/"+fmt.Sprint(seed)+".png", buf, 0644)
	if err != nil {
		panic(err)
	}

	jsonData, err := randomstylesheet.JSON(seed, true)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("data/"+fmt.Sprint(seed)+".json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("run : %d - duration : %v", iteration, time.Since(start)))
}

func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
