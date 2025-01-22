package utils

import (
	"path/filepath"
	"strings"
)

func MustNameExt(filename string) (name, ext string) {

	ext = filepath.Ext(filename)
	if ext != "" {
		name = strings.ReplaceAll(filename, ("." + ext), "")
	}

	return
}
