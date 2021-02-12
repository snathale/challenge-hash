package main

type Config struct {
	DbPort     int    `json:"port"`
	DbPassword string `json:"password"`
	DbDatabase string `json:"database"`
	DbHost     string `json:"host"`
	LogLevel   int    `json:"logLevel"`
}
