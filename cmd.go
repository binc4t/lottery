package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func CMD() {
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
		go Cycle()
		<-make(chan interface{})
	},
}

var notifyCmd = &cobra.Command{
	Use:  "notify",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		CalAndNotify()
	},
}
