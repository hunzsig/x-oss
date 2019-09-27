package oss

import (
	"bytes"
	"github.com/kataras/iris/core/errors"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

/**
 * 图片编码
 */
func ImageEncode(inputName string, file *os.File, rgba *image.RGBA) error {
	var err error
	if strings.HasSuffix(inputName, "jpg") || strings.HasSuffix(inputName, "jpeg") {
		err = jpeg.Encode(file, rgba, nil)
	} else if strings.HasSuffix(inputName, "png") {
		err = png.Encode(file, rgba)
	} else if strings.HasSuffix(inputName, "gif") {
		err = gif.Encode(file, rgba, nil)
	} else {
		err = errors.New("Not support this format")
	}
	return err
}

/**
 * 图片色彩反转
 */
func ImageColorReverse(source string) *image.RGBA {
	ff, _ := ioutil.ReadFile(source)
	bbb := bytes.NewBuffer(ff)
	m, _, _ := image.Decode(bbb)
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			rUint8 := uint8(r >> 8)
			gUint8 := uint8(g >> 8)
			bUint8 := uint8(b >> 8)
			aUint8 := uint8(a >> 8)
			rUint8 = 255 - rUint8
			gUint8 = 255 - gUint8
			bUint8 = 255 - bUint8
			newRgba.SetRGBA(i, j, color.RGBA{R: rUint8, G: gUint8, B: bUint8, A: aUint8})
		}
	}
	return newRgba
}

/**
 * 图片灰度化处理
 */
func ImageColorGrayscale(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			_, g, _, a := colorRgb.RGBA()
			gUint8 := uint8(g >> 8)
			aUint8 := uint8(a >> 8)
			newRgba.SetRGBA(i, j, color.RGBA{R: gUint8, G: gUint8, B: gUint8, A: aUint8})
		}
	}
	return newRgba
}
