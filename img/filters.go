package img

import (
	"image"
	"image/draw"

	"github.com/disintegration/gift"
)

type TwoMedFilter struct {
	size int
}

func (p *TwoMedFilter) Bounds(srcBounds image.Rectangle) (dstBounds image.Rectangle) {
	dstBounds = image.Rect(0, 0, srcBounds.Dx(), srcBounds.Dy())
	return
}

func (p *TwoMedFilter) Draw(dst draw.Image, src image.Image, options *gift.Options) {
	// twomed
	g := gift.New(
		gift.GaussianBlur(.1),
		gift.CropToSize(1000, 700, gift.BottomAnchor),
		gift.Contrast(-20),
		gift.Brightness(-10),
		gift.Saturation(50),
	)
	g.Draw(dst, src)
}

