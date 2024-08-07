package golang

import (
	"GolangPractice/utils/logger"
	"fmt"
	imageColor "image/color"
	"image/png"
	"os"
	"testing"
)

func TestPNG(t *testing.T) {
	reader, err := os.Open("materials/img.png")
	if err != nil {
	    logger.Errorln(err.Error())
		return
	}
	defer reader.Close()
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))

	img, err := png.Decode(reader)
	if err != nil {
		logger.Errorln(err.Error())
		return
	}
	logger.Infoln(ColorToHex(img.At(1,1)))
}

//ColorToHex convert color.Color into Hex string, ignoring the alpha channel.
func ColorToHex(c imageColor.Color) string {
	r, g, b, _ := c.RGBA()
	return RGBToHex(uint8(r>>8), uint8(g>>8), uint8(b>>8))
}

//RGBToHex converts an RGB triple to a Hex string in the format of 0xffff.
func RGBToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}