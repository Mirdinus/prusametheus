package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Printer struct {
		IP       string `json:"ip"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"printer"`
	Server struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	} `json:"server"`
	FetchIntervalS int `json:"fetch_interval_s"`
}

var AppConfig Config

func LoadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
}
