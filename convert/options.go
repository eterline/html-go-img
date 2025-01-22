package convert

import "fmt"

type HtmlConverterOptionFunc func(o *HtmlConverterOptions)

type HtmlConverterOptions struct {
	Height string
	Width  string
}

func SetWidth(n uint) HtmlConverterOptionFunc {
	return func(o *HtmlConverterOptions) {
		if n != 0 {
			o.Width = fmt.Sprintf("w %v", n)
		}
	}
}

func SetHeight(n uint) HtmlConverterOptionFunc {
	return func(o *HtmlConverterOptions) {
		if n != 0 {
			o.Height = fmt.Sprintf("h %v", n)
		}
	}
}

func SetSquared(side uint) HtmlConverterOptionFunc {
	return func(o *HtmlConverterOptions) {
		if side != 0 {
			o.Height = fmt.Sprintf("h %v", side)
			o.Width = fmt.Sprintf("w %v", side)
		}
	}
}

func SetProportional(bottom uint, widthScalar, heightScalar uint8) HtmlConverterOptionFunc {
	return func(o *HtmlConverterOptions) {
		if bottom != 0 && widthScalar != 0 && heightScalar != 0 {

			h := uint(bottom/uint(widthScalar)) * uint(heightScalar)

			o.Height = fmt.Sprintf("h %v", h)
			o.Width = fmt.Sprintf("w %v", bottom)
		}
	}
}
