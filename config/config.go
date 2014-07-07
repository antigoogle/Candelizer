package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BFX map[string]Market
}

type Market struct {
	CandleWidthMins int
	KeepCandleMins  int
	Indicators      []map[string]interface{}
}

func ReadConfig() Config {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		fmt.Printf("%s", err)
	}
	config := Config{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Printf("%s", err)
		return Config{}
	}
	return config

}
