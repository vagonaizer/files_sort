package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	SourcePath string `json: "source_path"`
	ImagesPath string `json:"images_path"`
	VideosPath string `json:"videos_path"`
	OtherPath  string `json:"other_path"`
}

func ReadPath() Config {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("source path: ")
	sourcePath, _ := reader.ReadString('\n')
	sourcePath = strings.TrimSpace(sourcePath)

	fmt.Println("folder for images: ")
	imagesPath, _ := reader.ReadString('\n')
	imagesPath = strings.TrimSpace(imagesPath)

	fmt.Println("folder for videos: ")
	videosPath, _ := reader.ReadString('\n')
	videosPath = strings.TrimSpace(videosPath)

	fmt.Println("folder for other: ")
	othersPath, _ := reader.ReadString('\n')
	othersPath = strings.TrimSpace(othersPath)

	cfg := Config{
		SourcePath: sourcePath,
		ImagesPath: imagesPath,
		VideosPath: videosPath,
		OtherPath:  othersPath,
	}
	return cfg
}
