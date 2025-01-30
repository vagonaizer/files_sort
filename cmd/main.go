package main

import (
	"encoding/json"
	"files_sort/internal"
	"fmt"
	"os"
)

func main() {
	fmt.Println("\n")
	internal.JsonHandler(internal.ReadPath())

	configData, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Ошибка при чтении config.json", err)
		return
	}

	var cfg internal.Config
	err = json.Unmarshal(configData, &cfg)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	fmt.Println("Папка для проверки:", cfg.SourcePath)

	filesByExt, err := internal.GetFilesByExtension(cfg.SourcePath)
	if err != nil {
		fmt.Println("Ошибка сканирования файлов:", err)
		return
	}

	for ext, files := range filesByExt {
		fmt.Printf("Расширение: %s, количество файлов: %d\n", ext, len(files))
		for _, f := range files {
			fmt.Printf("  %s\n", f)
		}
		fmt.Println()
	}

}
