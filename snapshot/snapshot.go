package snapshot

import (
	"context"
	"math"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Params Snapshot params
type Params struct {
	WaitPeriod int
}

// Snap Get snapshot from website, based on configuration params
func Snap(ctx context.Context, url string, params Params) (buf []byte, err error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	duration := 15
	if params.WaitPeriod != 0 {
		duration = params.WaitPeriod
	}

	// This is a hack to wait the whole page to load
	interval := time.Duration(duration) * time.Second

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			time.Sleep(interval)

			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypeLandscapePrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			buf, err = page.CaptureScreenshot().
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).
				Do(ctx)

			return err
		}),
	})

	return buf, err
}
