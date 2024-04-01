package utils

import (
	"fmt"
	"path/filepath"
	"time"
)

func GenerateImageFileName(originalFileName, prefix string) string {
	extension := filepath.Ext(originalFileName)
	return fmt.Sprintf("%s-%s%s", prefix, time.Now().Format("20060102150405"), extension)
}
