package conf

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	StartUrls      []string `json:"start_urls"`
	Keywords       []string `json:"keywords"`
	FilterKeywords []string `json:"filter_keywords"`
	MininumChars   int      `json:"mininum_chars"` // 介绍详情至少要多少个文字
}

func Load(confPath string) (*Config, error) {
	jsonFile, err := os.Open(confPath)
	if err != nil {
		log.Println("Error opening json file:", err)
		return nil, err
	}

	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	var config Config
	for {
		err := decoder.Decode(&config)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error decoding json:", err)
			return nil, err
		}

	}
	return &config, err
}
