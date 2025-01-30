package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FilesSort(cfg Config, filesByExt map[string][]string) error {
	for ext, files := range filesByExt {
		for _, fileName := range files {
			oldPath := filepath.Join(cfg.SourcePath, fileName)

			var targetDir string
			switch ext {
			case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
				targetDir = cfg.ImagesPath
			case ".mp4", ".avi", ".mov", ".mkv":
				targetDir = cfg.VideosPath
			default:
				targetDir = cfg.OtherPath
			}

			err := os.MkdirAll(targetDir, 0755)
			if err != nil {
				return fmt.Errorf("не удалось создать директорию %s: %v", targetDir, err)
			}

			newPath := filepath.Join(targetDir, fileName)

			err = os.Rename(oldPath, newPath)
			if err != nil {
				err = CopyFile(oldPath, newPath)
				if err != nil {
					return fmt.Errorf("не удалось скопировать %s -> %s: %v:", oldPath, newPath, err)
				}
				if err := os.Remove(oldPath); err != nil {
					return fmt.Errorf("не удалось удалить исходный файл %s: %v", &oldPath, err)
				}
			}
			fmt.Printf("Файл %s перемещен в %s\n", fileName, targetDir)
		}
	}
	return nil
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	return destFile.Sync()
}
