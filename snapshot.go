package snapurl

import (
	"context"
	"math"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

const (
	DefaultWidth  int64   = 1440
	DefaultHeight int64   = 900
	DefaultScale  float64 = 1.0
)

// Params Snapshot params
type Params struct {
	WaitPeriod time.Duration
	Width      int64
	Height     int64
	Scale      float64
}

// Snap Get snapshot from website, based on configuration params
func Snap(ctx context.Context, url string, p Params) (buf []byte, err error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	if p.Width == 0 {
		p.Width = DefaultWidth
	}

	if p.Height == 0 {
		p.Width = DefaultHeight
	}

	if p.Scale == 0 {
		p.Scale = DefaultScale
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
