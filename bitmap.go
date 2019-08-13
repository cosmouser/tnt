package tnt

import (
	"image"
)

// Deserialize deserializes the given minimap into a bitmap
func (m *Bitmap) Deserialize() (image.Image, error) {
	maxPoint := image.Point{
		X: m.Width,
		Y: m.Height,
	}
	imageRect := image.Rectangle{
		Min: image.Point{},
		Max: maxPoint,
	}
	result := image.NewPaletted(imageRect, m.Palette)
	result.Pix = m.Data
	return result.SubImage(imageRect), nil
}
