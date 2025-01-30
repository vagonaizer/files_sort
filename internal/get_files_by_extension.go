package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesByExtension(sourcePath string) (map[string][]string, error) {
	extMap := make(map[string][]string)

	entries, err := os.ReadDir(sourcePath)
	if err != nil {
		return nil, fmt.Errorf("Не удалось прочитать папку %s: %v", sourcePath, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		ext := strings.ToLower(filepath.Ext(filename))

		extMap[ext] = append(extMap[ext], filename)
	}

	return extMap, nil
}
