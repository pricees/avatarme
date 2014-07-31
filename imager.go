/*
  Attempting to create identicon based on Chris Branson's Ruby implementation   http://goo.gl/yVdEJT
*/
package avatar

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Identicon struct {
	avatar          Avatar
	filename        string
	borderSize      int
	squareSize      int
	gridSize        int
	backgroundColor int
}

// Create an identicon from a hash
func (i *Identicon) Create() bool {
	if i.filename == "" {
		i.filename = fmt.Sprintf("%s%s", i.avatar.Hash(), ".png")
	}

	f, err := os.OpenFile(i.filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	identiconSize := i.squareSize * i.gridSize
	totalSize := (2 * i.borderSize) + identiconSize

	m := image.NewNRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{totalSize, totalSize}})

	var hash int

	hashLength := len(i.avatar.Hash())
	for i, ord := range i.avatar.Hash() {
		hash = hash + (int(ord) << uint(hashLength-i))
	}

	fgColor := color.NRGBA{
		uint8(hash & 0xff), uint8((hash >> 8) & 0xff), uint8((hash >> 16) & 0xff), 0xff}
	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			m.SetNRGBA(x, y, fgColor)

		}
	}
	if err = png.Encode(f, m); err != nil {
		fmt.Println(err)
	}
	return true
}
