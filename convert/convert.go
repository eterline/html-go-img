package convert

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/eterline/html-go-img/execute"
)

type ImageExt int

const (
	_ ImageExt = iota
	PNG
	JPG
)

const (
	DefaultFile    = "converted-html"
	DefaultFileExt = "png"
)

type HtmlConverter struct {
	exec    *execute.Executer // executable binary parameters
	payload []byte            // input HTML
	output  []byte            // output image/pdf data
}

// Create new image converter
func NewConverterImg() *HtmlConverter {

	args := []execute.BinArg{"q"}

	return &HtmlConverter{
		exec: execute.NewExecuter(
			execute.HTMLtoIMGPath(),
			args,
		),
	}
}

// Create new pdf converter. TODO: In process
func NewConverterPdf() *HtmlConverter {
	return &HtmlConverter{
		exec: execute.NewExecuter(
			execute.HTMLtoPDFPath(),
			[]execute.BinArg{"q"},
		),
	}
}

// HTML content from string
func (c *HtmlConverter) StringPayload(payload string) {
	c.payload = []byte(payload)
}

// HTML content byte stream
func (c *HtmlConverter) BytesPayload(payload []byte) {
	c.payload = payload
}

// HTML content from io.Reader
func (c *HtmlConverter) ReadPayload(r io.Reader) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	c.payload = bytes
	return nil
}

// HTML content from string
func (c *HtmlConverter) WriteTo(w io.Writer) error {
	_, err := w.Write(c.output)
	return err
}

// Creates new image file with name and path
func (c *HtmlConverter) SaveFile(name string, ext ImageExt) error {

	switch ext {
	case PNG:
		name += ".png"
	case JPG:
		name += ".jpg"
	default:
		return ErrUnsupportedExt
	}

	data, err := c.ToFormat(ext)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

// Return base64 string
func (c *HtmlConverter) ToBase64() string {
	return base64.StdEncoding.EncodeToString(c.output)
}

// Make convert command
func (c *HtmlConverter) Convert() error {
	out, err := c.exec.ProcessConverter(c.payload)

	c.output = out
	return err
}

// Output format bytes convert
func (c *HtmlConverter) ToFormat(ext ImageExt) ([]byte, error) {

	buf := new(bytes.Buffer)

	decoded, err := jpeg.Decode(bytes.NewReader(c.output))
	for err != nil {
		c.output = c.output[1:]
		if len(c.output) == 0 {
			return nil, ErrNilPayload
		}
		decoded, err = jpeg.Decode(bytes.NewReader(c.output))
	}

	if decoded == nil {
		return nil, ErrNilImage
	}

	switch ext {

	case PNG:
		err = png.Encode(buf, decoded)
		if err != nil {
			return nil, ErrEncode(err)
		}
		return buf.Bytes(), nil

	case JPG:
		err = jpeg.Encode(buf, decoded, nil)
		if err != nil {
			return nil, ErrEncode(err)
		}
		return buf.Bytes(), nil

	default:
		return nil, ErrUnsupportedExt

	}
}
