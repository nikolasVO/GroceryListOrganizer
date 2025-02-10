package service

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// SynonymsMap – глобальный словарь, в котором каждому синониму сопоставляется каноническое название.
var SynonymsMap map[string]string

// LoadSynonyms загружает синонимы из файла по указанному пути.
// Формат файла: каждая строка имеет вид: <каноническое название>: синоним1, синоним2, синоним3,...
func LoadSynonyms(filePath string) error {
	SynonymsMap = make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			// Пропускаем пустые строки и комментарии
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			log.Printf("Неверный формат строки: %s", line)
			continue
		}

		canonical := strings.TrimSpace(parts[0])
		synonymsList := strings.Split(parts[1], ",")
		for _, syn := range synonymsList {
			syn = strings.TrimSpace(strings.ToLower(syn))
			if syn != "" {
				SynonymsMap[syn] = canonical
			}
		}
	}
	return scanner.Err()
}
