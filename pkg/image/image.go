package image

import (
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type Image struct {
	reader io.Reader
}

func NewImage(r io.Reader) *Image {
	return &Image{
		reader: r,
	}
}

func (i *Image) loadImage() (image.Image, error) {
	img, _, err := image.Decode(i.reader)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (i *Image) imageToRGBA(img image.Image, rect image.Rectangle) *image.RGBA {
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	return rgba
}

func (i *Image) GetRGBColorMap() (RGBColorMap, error) {
	colors := make(RGBColorMap)

	img, err := i.loadImage()
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	rgba := i.imageToRGBA(img, rect)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			i := rgba.PixOffset(x, y)
			s := rgba.Pix[i : i+4 : i+4]
			color := RGBColor{s[0], s[1], s[2]}
			colors[color]++
		}
	}

	return colors, nil
}
