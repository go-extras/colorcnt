package image

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

const biColorPngImage = "sample-images/bi-color.png"
const jpegImage = "sample-images/jpeg.jpg"
const multipleColorsPngImage = "sample-images/multiple-colors.png"
const SingleColorPngImage = "sample-images/single-color-red.png"

type ImageTestSuite struct {
	suite.Suite
}

func TestImage(t *testing.T) {
	suite.Run(t, new(ImageTestSuite))
}

func (s *ImageTestSuite) TestGetRBBColorMap_BiColorPngImage() {
	inputFile, err := os.Open(biColorPngImage)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	img := NewImage(inputFile)
	colors, err := img.GetRGBColorMap()
	s.NoError(err)
	s.Len(colors, 2, "number of colors don't match: %d != 2", len(colors))
	s.Equal(192000, colors[RGBColor{0x00, 0x2B, 0x7F}])
	s.Equal(192000, colors[RGBColor{0xCE, 0x11, 0x26}])
}

func (s *ImageTestSuite) TestGetRBBColorMap_JpegImage() {
	inputFile, err := os.Open(jpegImage)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	img := NewImage(inputFile)
	colors, err := img.GetRGBColorMap()
	s.NoError(err)
	s.Len(colors, 218000, "number of colors don't match: %d != 218000", len(colors))
}

func (s *ImageTestSuite) TestGetRBBColorMap_SingleColorPngImage() {
	inputFile, err := os.Open(SingleColorPngImage)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	img := NewImage(inputFile)
	colors, err := img.GetRGBColorMap()
	s.NoError(err)
	s.Len(colors, 1, "number of colors don't match: %d != 1", len(colors))
	s.Equal(4000000, colors[RGBColor{0xFF, 0x00, 0x00}])
}

func (s *ImageTestSuite) TestGetRBBColorMap_MultipleColorsPng() {
	inputFile, err := os.Open(multipleColorsPngImage)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	img := NewImage(inputFile)
	colors, err := img.GetRGBColorMap()
	s.NoError(err)
	s.Len(colors, 5586, "number of colors don't match: %d != 5586", len(colors))
}
