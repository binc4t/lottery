package main

import (
	"context"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/lark"
	"log"
	"strconv"
	"strings"
	"time"
)

func notifyLark(webHookURL, subject, message string) {
	larkWebhookSvc := lark.NewWebhookService(webHookURL)

	notifier := notify.New()
	notifier.UseServices(larkWebhookSvc)

	if err := notifier.Send(context.Background(), subject, message); err != nil {
		log.Fatalf("notifier.Send() failed: %s", err.Error())
	}

	log.Println("notification sent")
}

func CalAndNotify() {
	data, err := GetLotteryData(100)
	if err != nil {
		log.Fatal(err)
	}
	ret := PredictByMode(data, Config.PredictNum)
	retStr := strconv.Itoa(ret[0])
	for i := 1; i < len(ret); i++ {
		retStr = retStr + " " + strconv.Itoa(ret[i])
	}
	notifyLark(Config.WebHookURL, "今日财富密码", retStr)
}

func Cycle() {
	ticker := time.NewTicker(time.Second)
	h, m := getHourMinuteFromTime(Config.Time)
	for {
		select {
		case now := <-ticker.C:
			for day := range Config.DaysOfWeek {
				if int(now.Weekday()) == day {
					if now.Hour() == h && now.Minute() == m && now.Second() == 0 {
						CalAndNotify()
					}
				}
			}
		}
	}
}

func getHourMinuteFromTime(t string) (hour, minute int) {
	seps := strings.Split(t, ":")
	var err error
	if len(seps) != 2 {
		hour, err = strconv.Atoi(seps[0])
		if err != nil {
			log.Fatal(err)
		}
		minute, err = strconv.Atoi(seps[0])
		if err != nil {
			log.Fatal(err)
		}
	}
}
