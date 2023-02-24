package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func cmd() {
	rootCmd.AddCommand(dataCmd)
	rootCmd.AddCommand(daemonCmd)
	rootCmd.AddCommand(notifyCmd)
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:  "predict",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		num := 30
		var err error
		if len(args) > 0 {
			num, err = strconv.Atoi(args[0])
			if err != nil || num > 100 || num < 0 {
				log.Fatal("invalid num")
			}
		}
		data, err := GetLotteryData(100)
		if err != nil {
			log.Fatal(err)
		}
		ret := PredictByMode(data, num)
		fmt.Println(ret)
	},
}

var dataCmd = &cobra.Command{
	Use:  "data",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		num := 30
		var err error
		if len(args) > 0 {
			num, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("invalid num")
			}
			if num != 30 && num != 50 && num != 100 {
				log.Fatal("invalid num")
			}
		}
		data, err := GetLotteryData(num)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range data {
			fmt.Printf("%v: %v\n", v.id, v.nums)
		}
	},
}

var daemonCmd = &cobra.Command{
	Use:  "daemon",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("please provide lark webHookURL")
		}
		webHookURL := args[0]
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case now := <-ticker.C:
				// check if today is Tuesday, Thursday, or Sunday and the current time is 18:00
				if now.Weekday() == time.Tuesday || now.Weekday() == time.Thursday || now.Weekday() == time.Sunday {
					if now.Hour() == 18 && now.Minute() == 0 && now.Second() == 0 {
						data, err := GetLotteryData(100)
						if err != nil {
							log.Fatal(err)
						}
						ret := PredictByMode(data, 30)
						retStr := strconv.Itoa(ret[0])
						for i := 1; i < len(ret); i++ {
							retStr = retStr + " " + strconv.Itoa(ret[i])
						}
						NotifyLark(webHookURL, "number today", retStr)
					}
				}
			}
		}
	},
}

var notifyCmd = &cobra.Command{
	Use:  "notify",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("please provide lark webHookURL")
		}
		webHookURL := args[0]
		data, err := GetLotteryData(100)
		if err != nil {
			log.Fatal(err)
		}
		ret := PredictByMode(data, 30)
		retStr := strconv.Itoa(ret[0])
		for i := 1; i < len(ret); i++ {
			retStr = retStr + " " + strconv.Itoa(ret[i])
		}
		NotifyLark(webHookURL, "number", retStr)
	},
}
