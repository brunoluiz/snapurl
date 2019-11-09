package snapurl

import (
	"context"
	"math"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/creasty/defaults"
)

// Params Snapshot params
type Params struct {
	WaitPeriod time.Duration `default:"5s"`
	Width      int64         `default:"1440"`
	Height     int64         `default:"900"`
	Scale      float64       `default:"1.0"`
	Format     string        `default:"jpeg"`
	Quality    int64         `default:"90"`
}

// Snap Get snapshot from website, based on configuration params
func Snap(ctx context.Context, url string, p Params) (buf []byte, err error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	if err := defaults.Set(&p); err != nil {
		return nil, err
	}

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(p.WaitPeriod),
		chromedp.EmulateViewport(p.Width, p.Height, chromedp.EmulateScale(p.Scale)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.
				SetDeviceMetricsOverride(width, height, p.Scale, false).
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
				WithFormat(page.CaptureScreenshotFormat(p.Format)).
				WithQuality(p.Quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  p.Scale,
				}).
				Do(ctx)

			return err
		}),
	})

	return buf, err
}
