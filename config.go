package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var Config LotteryConfig

type LotteryConfig struct {
	WebHookURL string         `json:"web_hook_url"` // Lark Robot WebHookURL, see https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
	DaysOfWeek []time.Weekday `json:"days_of_week"` // notify on every these days of week, default: {time.Tuesday, time.Thursday, time.Sunday}
	PredictNum int            `json:"predict_num"`  // Predict by recent predict_num's lottery results, default: 30
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

	if Config.WebHookURL == "" {
		log.Fatal("no web_hook_url provided")
	}

	if Config.PredictNum == 0 {
		Config.PredictNum = 30
	}
	if len(Config.DaysOfWeek) == 0 {
		Config.DaysOfWeek = []time.Weekday{time.Tuesday, time.Thursday, time.Sunday}
	}
}
