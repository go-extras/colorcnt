package image

import (
	"container/heap"
	"fmt"
)

type RGBColor [3]uint8

func (c RGBColor) Red() uint8 {
	return c[0]
}

func (c RGBColor) Green() uint8 {
	return c[1]
}

func (c RGBColor) Blue() uint8 {
	return c[2]
}

func (c RGBColor) String() string {
	return fmt.Sprintf("#%02X%02X%02X", c.Red(), c.Green(), c.Blue())
}

type RGBColorMap map[RGBColor]int

func (m RGBColorMap) MaxUsedColors(n uint) (result []RGBColor, cnt uint) {
	h := getHeap(m)
	maxColors := uint(len(m))

	result = make([]RGBColor, n, n)

	cnt = n
	if maxColors < cnt {
		// edge case: image has fewer than n colors
		cnt = maxColors
	}

	for i := uint(0); i < cnt; i++ {
		v := heap.Pop(h)
		color := v.(kv).Key
		result[i] = color
	}

	return result, cnt
}
