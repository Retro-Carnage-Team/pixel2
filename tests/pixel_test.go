package pixel_test

import (
	"bytes"
	"image"
	"os"
	"testing"

	_ "image/png"

	pixel "github.com/Retro-Carnage-Team/pixel2"
	"github.com/Retro-Carnage-Team/pixel2/backends/opengl"
)

// onePixelImage is the byte representation of a 1x1 solid white png file
var onePixelImage = []byte{
	137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 1, 0, 0, 0, 1, 8, 2,
	0, 0, 0, 144, 119, 83, 222, 0, 0, 1, 130, 105, 67, 67, 80, 73, 67, 67, 32, 112, 114, 111, 102, 105, 108, 101, 0,
	0, 40, 145, 125, 145, 59, 72, 3, 65, 20, 69, 143, 73, 68, 17, 37, 133, 41, 68, 44, 182, 80, 43, 5, 81, 17, 75,
	141, 66, 16, 34, 132, 168, 96, 212, 194, 221, 141, 137, 66, 118, 13, 187, 9, 54, 150, 130, 109, 192, 194, 79,
	227, 175, 176, 177, 214, 214, 194, 86, 16, 4, 63, 32, 54, 182, 86, 138, 54, 18, 214, 55, 73, 32, 65, 140, 3,
	195, 28, 238, 188, 123, 121, 243, 6, 124, 71, 25, 211, 114, 3, 3, 96, 217, 57, 39, 30, 9, 107, 243, 137, 5, 173,
	233, 149, 0, 65, 90, 104, 0, 221, 116, 179, 227, 177, 88, 148, 186, 235, 235, 94, 213, 193, 93, 191, 202, 170,
	95, 247, 231, 106, 75, 174, 184, 38, 52, 104, 194, 99, 102, 214, 201, 9, 47, 11, 143, 108, 228, 178, 138, 247,
	132, 67, 230, 170, 158, 20, 62, 23, 238, 115, 164, 65, 225, 71, 165, 27, 101, 126, 83, 156, 46, 177, 79, 101,
	134, 156, 217, 248, 132, 112, 72, 88, 75, 215, 176, 81, 195, 230, 170, 99, 9, 15, 11, 119, 39, 45, 91, 242, 125,
	243, 101, 78, 42, 222, 84, 108, 101, 242, 102, 165, 79, 245, 194, 214, 21, 123, 110, 70, 233, 178, 187, 136, 48,
	197, 52, 49, 52, 12, 242, 172, 145, 33, 71, 191, 156, 182, 40, 46, 113, 185, 15, 215, 241, 119, 150, 252, 49,
	113, 25, 226, 90, 195, 20, 199, 36, 235, 88, 232, 37, 63, 234, 15, 126, 207, 214, 77, 13, 13, 150, 147, 90, 195,
	208, 248, 226, 121, 31, 61, 208, 180, 3, 197, 130, 231, 125, 31, 123, 94, 241, 4, 252, 207, 112, 101, 87, 253,
	235, 71, 48, 250, 41, 122, 161, 170, 117, 31, 66, 112, 11, 46, 174, 171, 154, 177, 11, 151, 219, 208, 241, 148,
	213, 29, 189, 36, 249, 101, 251, 82, 41, 120, 63, 147, 111, 74, 64, 251, 45, 180, 44, 150, 231, 86, 185, 231,
	244, 1, 102, 101, 86, 209, 27, 216, 63, 128, 222, 180, 100, 47, 213, 121, 119, 115, 237, 220, 254, 173, 169,
	204, 239, 7, 178, 211, 114, 90, 10, 150, 157, 65, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 46, 35, 0, 0, 46, 35, 1,
	120, 165, 63, 118, 0, 0, 0, 7, 116, 73, 77, 69, 7, 227, 4, 15, 10, 5, 36, 189, 4, 224, 88, 0, 0, 0, 25, 116, 69,
	88, 116, 67, 111, 109, 109, 101, 110, 116, 0, 67, 114, 101, 97, 116, 101, 100, 32, 119, 105, 116, 104, 32, 71,
	73, 77, 80, 87, 129, 14, 23, 0, 0, 0, 12, 73, 68, 65, 84, 8, 215, 99, 120, 241, 226, 61, 0, 5, 123, 2, 192, 194,
	77, 211, 95, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130,
}

func TestMain(m *testing.M) {
	opengl.Run(func() {
		os.Exit(m.Run())
	})
}

func TestSprite_Draw(t *testing.T) {
	img, _, err := image.Decode(bytes.NewReader(onePixelImage))
	if err != nil {
		t.Fatalf("Could not decode image: %v", err)
	}
	pic := pixel.PictureDataFromImage(img)

	sprite := pixel.NewSprite(pic, pic.Bounds())

	cfg := opengl.WindowConfig{
		Title:     "testing",
		Bounds:    pixel.R(0, 0, 150, 150),
		Invisible: true,
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		t.Fatalf("Could not create window: %v", err)
	}

	sprite.Draw(win, pixel.IM)
}
