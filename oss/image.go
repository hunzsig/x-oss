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
	"math"
	"os"
)

/**
 * 图片编码
 */
func ImageEncode(file *os.File, rgba *image.RGBA, suffix string) error {
	var err error
	if suffix == "jpg" || suffix == "jpeg" {
		err = jpeg.Encode(file, rgba, nil)
	} else if suffix == "png" {
		err = png.Encode(file, rgba)
	} else if suffix == "gif" {
		err = gif.Encode(file, rgba, nil)
	} else {
		err = errors.New("Not support this format")
	}
	return err
}

/**
 * 图片转ascii编码
 */
func ImageAscii(file *os.File, rgba *image.RGBA) error {
	bounds := rgba.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	arr := []string{"M", "N", "H", "Q", "$", "O", "C", "?", "7", ">", "!", ":", "–", ";", "."}
	var err error
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := rgba.At(j, i)
			_, g, _, _ := colorRgb.RGBA()
			avg := uint8(g >> 8)
			num := avg / 18
			_, err = file.WriteString(arr[num])
			if j == dx-1 {
				_, err = file.WriteString("\n")
			}
		}
	}
	return err
}

/**
 * 获取图片 source
 */
func ImageSource(uri string) image.Image {
	ff, _ := ioutil.ReadFile(uri)
	b := bytes.NewBuffer(ff)
	m, _, _ := image.Decode(b)
	return m
}

/**
 * 获取图片 RGBA
 */
func ImageRGBA(uri string) *image.RGBA {
	source := ImageSource(uri)
	bounds := source.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(source.Bounds())
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := source.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			rUint8 := uint8(r >> 8)
			gUint8 := uint8(g >> 8)
			bUint8 := uint8(b >> 8)
			aUint8 := uint8(a >> 8)
			newRgba.SetRGBA(i, j, color.RGBA{R: rUint8, G: gUint8, B: bUint8, A: aUint8})
		}
	}
	return newRgba
}

/**
 * 图片裁剪
 */
func ImageThumb(m *image.RGBA, x1 int, y1 int, x2 int, y2 int) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	fmt.Println(dx, dy)
	if x1 < 0 {
		x1 = 0
	}
	if x2 < 0 {
		x2 = 0
	}
	if y1 < 0 {
		y1 = 0
	}
	if y2 < 0 {
		y2 = 0
	}
	if x1 > dx {
		x1 = dx
	}
	if x2 > dx {
		x2 = dx
	}
	if y1 > dy {
		y1 = dy
	}
	if y2 > dy {
		y2 = dy
	}
	source := m.SubImage(image.Rect(x1, y1, x2, y2))
	bounds = source.Bounds()
	dx = bounds.Dx()
	dy = bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := source.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			rUint8 := uint8(r >> 8)
			gUint8 := uint8(g >> 8)
			bUint8 := uint8(b >> 8)
			aUint8 := uint8(a >> 8)
			newRgba.SetRGBA(i, j, color.RGBA{R: rUint8, G: gUint8, B: bUint8, A: aUint8})
		}
	}
	return newRgba
}

/**
 * 图片缩放(int)
 */
func ImageResizeInt(m *image.RGBA, newDx int, newDy int) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	if newDy <= 0 {
		newDy = int(float64(newDx) * (float64(dy) / float64(dx)))
	}
	newRgba := image.NewRGBA(image.Rect(0, 0, newDx, newDy))
	_ = graphics.Scale(newRgba, m)
	return newRgba
}

/**
 * 图片缩放(float)
 */
func ImageResizeFloat(m *image.RGBA, pcX float64, pcY float64) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newDx := float64(dx) * pcX
	newDy := float64(0)
	if pcY <= 0 {
		newDy = newDx * (float64(dy) / float64(dx))
	} else {
		newDy = float64(dy) * pcY
	}
	newRgba := image.NewRGBA(image.Rect(0, 0, int(newDx), int(newDy)))
	_ = graphics.Scale(newRgba, m)
	return newRgba
}

/**
 * 图片旋转
 */
func ImageRotate(m *image.RGBA, angle int) *image.RGBA {
	angle = angle % 360
	if angle == 0 {
		return m
	}
	//弧度转换
	radian := float64(angle) * math.Pi / 180.0
	cos := math.Cos(float64(radian))
	sin := math.Sin(radian)
	//原图的宽高
	w := float64(m.Bounds().Dx())
	h := float64(m.Bounds().Dy())
	//新图高宽
	W := int(math.Max(math.Abs(float64(w*cos-h*sin)), math.Abs(w*cos+h*sin)))
	H := int(math.Max(math.Abs(w*sin-h*cos), math.Abs(w*sin+h*cos)))
	//
	newRgba := image.NewRGBA(image.Rect(0, 0, W, H))
	_ = graphics.Rotate(newRgba, m, &graphics.RotateOptions{Angle: radian})
	return newRgba
}

/**
 * 图片模糊
 */
func ImageBlur(m *image.RGBA, distance float64) *image.RGBA {
	if distance == 0.0 {
		return m
	}
	newRgba := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dx(), m.Bounds().Dy()))
	_ = graphics.Blur(newRgba, m, &graphics.BlurOptions{StdDev: distance})
	return newRgba
}

/**
 * 图片色彩反转
 */
func ImageColorReverse(m *image.RGBA) *image.RGBA {
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
func ImageColorGrayscale(m *image.RGBA) *image.RGBA {
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
