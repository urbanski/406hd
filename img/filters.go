package img

import (
	"image"
	"image/draw"

	"github.com/disintegration/gift"
)

type twomedFilter struct {
	size int
}

func (p *twomedFilter) Bounds(srcBounds image.Rectangle) (dstBounds image.Rectangle) {
	dstBounds = image.Rect(0, 0, srcBounds.Dx(), srcBounds.Dy())
	return
}

func (p *twomedFilter) Draw(dst draw.Image, src image.Image, options *gift.Options) {
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

func TwoMedicine() gift.Filter {
	return &twomedFilter{}
}


type manyglacierFilter struct {
	size int
}

func (p *manyglacierFilter) Bounds(srcBounds image.Rectangle) (dstBounds image.Rectangle) {
	dstBounds = image.Rect(0, 0, srcBounds.Dx(), srcBounds.Dy())
	return
}

func (p *manyglacierFilter) Draw(dst draw.Image, src image.Image, options *gift.Options) {
	// many glacier
	g := gift.New(
		gift.GaussianBlur(.5),
		gift.CropToSize(1024, 750, gift.BottomAnchor),
		gift.Contrast(-30),
		gift.Brightness(10),
		gift.Gamma(0.5),
	)
	g.Draw(dst, src)
}

func ManyGlacier() gift.Filter {
	return &manyglacierFilter{}
}
