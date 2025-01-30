package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonHandler(cfg Config) {
	jsonData, err := json.MarshalIndent(cfg, "", "   ")
	if err != nil {
		fmt.Println("Ошибка при сериализации в JSON: ", err)
		return
	}

	configFilename := "config.json"
	err = os.WriteFile(configFilename, jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл: ", err)
		return
	}

	fmt.Printf("Настройки успешно сохранены в файл %s\n", configFilename)
}
