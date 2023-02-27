package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var Config LotteryConfig

type LotteryConfig struct {
	WebHookURL string         `json:"web_hook_url"`
	DaysOfWeek []time.Weekday `json:"days_of_week"`
	PredictNum int            `json:"predict_num"`
}

func InitConfig() {
	b, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	Config = LotteryConfig{}
	err = json.Unmarshal(b, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
