package opengl

import (
	"math"

	pixel "github.com/Retro-Carnage-Team/pixel2/"
)

func intBounds(bounds pixel.Rect) (x, y, w, h int) {
	x0 := int(math.Floor(bounds.Min.X))
	y0 := int(math.Floor(bounds.Min.Y))
	x1 := int(math.Ceil(bounds.Max.X))
	y1 := int(math.Ceil(bounds.Max.Y))
	return x0, y0, x1 - x0, y1 - y0
}
