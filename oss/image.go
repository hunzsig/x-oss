package oss

import (
	"bytes"
	"fmt"
	"github.com/hunzsig/graphics"
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
 * 获取图片source
 */
func ImageFetch(source string) image.Image {
	ff, _ := ioutil.ReadFile(source)
	bbb := bytes.NewBuffer(ff)
	m, _, _ := image.Decode(bbb)
	return m
}

/**
 * 图片色彩反转
 */
func ImageColorReverse(m image.Image) *image.RGBA {
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

/**
 * 图片缩放
 */
func ImageResize(m image.Image, newdx int) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(image.Rect(0, 0, newdx, newdx*dy/dx))
	graphics.Scale(newRgba, m)
	return newRgba
}

/**
 * 图片转为字符画（简易版）
 */
func ImageAscll(m image.Image, target string) {
	if m.Bounds().Dx() > 300 {
		m = ImageResize(m, 300)
	}
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	arr := []string{"M", "N", "H", "Q", "$", "O", "C", "?", "7", ">", "!", ":", "–", ";", "."}

	fileName := target
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := m.At(j, i)
			_, g, _, _ := colorRgb.RGBA()
			avg := uint8(g >> 8)
			num := avg / 18
			dstFile.WriteString(arr[num])
			if j == dx-1 {
				dstFile.WriteString("\n")
			}
		}
	}
}
