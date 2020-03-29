package util

import (
	"path/filepath"
	"strings"
)

// Verify if this file is a image with extension .jpeg,
// it's not case sensitive control
func IsJpegFile(path string) bool {
	extension := filepath.Ext(path)
	return strings.EqualFold(extension, ".jpg")
}
