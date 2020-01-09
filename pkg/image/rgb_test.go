package image

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RGBTestSuite struct {
	colors []RGBColor
	suite.Suite
}

func TestRGB(t *testing.T) {
	suite.Run(t, new(RGBTestSuite))
}

func (s *RGBTestSuite) TestString() {
	rgb := RGBColor([3]uint8{0xAC, 0xCA, 0xBA})
	s.Equal("#ACCABA", rgb.String())
}

func (s *RGBTestSuite) TestSingleColors() {
	rgb := RGBColor([3]uint8{0xAC, 0xCA, 0xBA})
	s.Equal(uint8(0xAC), rgb.Red())
	s.Equal(uint8(0xCA), rgb.Green())
	s.Equal(uint8(0xBA), rgb.Blue())
}

func (s *RGBTestSuite) SetupSuite() {
	s.colors = append(s.colors, RGBColor{0x00, 0x01, 0x02})
	s.colors = append(s.colors, RGBColor{0x01, 0x02, 0x03})
	s.colors = append(s.colors, RGBColor{0x02, 0x03, 0x04})
	s.colors = append(s.colors, RGBColor{0x03, 0x04, 0x05})
	s.colors = append(s.colors, RGBColor{0x04, 0x05, 0x06})
	s.colors = append(s.colors, RGBColor{0x05, 0x06, 0x07})
}

func (s *RGBTestSuite) TestGetTopNColors() {
	colors := RGBColorMap{}
	colors[s.colors[0]] = 1
	colors[s.colors[1]] = 2
	colors[s.colors[2]] = 3
	colors[s.colors[3]] = 4
	colors[s.colors[4]] = 5
	colors[s.colors[5]] = 6

	result, cnt := colors.MaxUsedColors(3)
	s.Equal([]RGBColor{s.colors[5], s.colors[4], s.colors[3]}, result)
	s.Equal(cnt, uint(3))

	result, cnt = colors.MaxUsedColors(6)
	s.Equal([]RGBColor{s.colors[5], s.colors[4], s.colors[3], s.colors[2], s.colors[1], s.colors[0]}, result)
	s.Equal(cnt, uint(6))

	// edge case
	result, cnt = colors.MaxUsedColors(0)
	s.Equal([]RGBColor{}, result)
	s.Equal(cnt, uint(0))

	// edge case
	result, cnt = colors.MaxUsedColors(7)
	s.Equal([]RGBColor{s.colors[5], s.colors[4], s.colors[3], s.colors[2], s.colors[1], s.colors[0], {}}, result)
	s.Equal(cnt, uint(6))
}
