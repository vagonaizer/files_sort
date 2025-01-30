package internal

import (
	"path/filepath"
	"strings"
)

func GetFileExtension(filename string) string { // функция для проверки расширения файла
	ext := filepath.Ext(filename)
	return strings.ToLower(ext)
}
