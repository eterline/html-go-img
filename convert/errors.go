package convert

import (
	"errors"
	"fmt"
)

type HtmlConverterError struct {
	code int
	err  error
}

func (e *HtmlConverterError) Error() string {
	return fmt.Sprintf("html converter error: %s", e.err.Error())
}

var (
	ErrUnsupportedExt = &HtmlConverterError{
		err: errors.New("unsupported output file extension"),
	}

	ErrNilImage = &HtmlConverterError{
		err: errors.New("decoded image is nil after processing"),
	}

	ErrNilPayload = &HtmlConverterError{
		err: errors.New("invalid payload or corrupt image data"),
	}

	ErrEncode = func(err error) *HtmlConverterError {
		return &HtmlConverterError{
			err: fmt.Errorf("failed to encode image: %w", err),
		}
	}
)
