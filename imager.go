/*
  Attempting to create identicon based on Chris Branson's Ruby implementation   http://goo.gl/yVdEJT
*/
package avatar

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
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

func (identicon *Identicon) drawIdenticon(hash int, m *image.RGBA) *image.RGBA {

	fgColor := color.NRGBA{uint8(hash & 0xff), uint8((hash >> 8) & 0xff), uint8((hash >> 16) & 0xff), 0xff}

	hash = hash >> 24

	var sqx, sqy, x, y int

	half := (identicon.gridSize * (identicon.gridSize + 1)) / 2

	for i := 0; i < half; i++ {
		if hash&1 == 1 {
			x = identicon.borderSize + (sqx * identicon.squareSize)
			y = identicon.borderSize + (sqy * identicon.squareSize)

			// draw left size
			draw.Draw(m, image.Rect(x, y, x+identicon.squareSize, y+identicon.squareSize), &image.Uniform{fgColor}, image.ZP, draw.Src)

			// mirror right hand side
			x = identicon.borderSize + ((identicon.gridSize - 1 - sqx) * identicon.squareSize)
			draw.Draw(m, image.Rect(x, y, x+identicon.squareSize, y+identicon.squareSize), &image.Uniform{fgColor}, image.ZP, draw.Src)
		}

		hash = hash >> 1
		sqy += 1

		if sqy == identicon.gridSize {
			sqy = 0
			sqx += 1
		}

	}
	return m
}

// Create an identicon from a hash
func (identicon *Identicon) Create() bool {
	if identicon.filename == "" {
		identicon.filename = fmt.Sprintf("%s%s", identicon.avatar.Hash(), ".png")
	}

	f, err := os.OpenFile(identicon.filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Turn hash into number
	var hash int

	hashLength := len(identicon.avatar.Hash())
	for i, ord := range identicon.avatar.Hash() {
		hash = hash + (int(ord) << uint(hashLength-i))
	}

	// Get new image
	identiconSize := identicon.squareSize * identicon.gridSize
	totalSize := (2 * identicon.borderSize) + identiconSize

	m := image.NewRGBA(image.Rect(0, 0, totalSize, totalSize))

	bgColor := color.NRGBA{uint8(hash >> 8), uint8((hash >> 16) & 0xff), uint8((hash >> 24) & 0xff), 0xff}

	// Draw background color
	draw.Draw(m, m.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	// Draw identicon
	identicon.drawIdenticon(hash, m)

	if err = png.Encode(f, m); err != nil {
		fmt.Println(err)
	}
	return true
}
