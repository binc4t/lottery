package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Unit struct {
	id   int
	nums []int
}

type ball struct {
	id    int
	count int
}

type balls []ball

func (b balls) Len() int {
	return len(b)
}

func (b balls) Less(i, j int) bool {
	return b[i].count < b[j].count
}

func (b balls) Swap(i, j int) {
	tmp := b[i]
	b[i] = b[j]
	b[j] = tmp
}

func PredictByMode(units []Unit, num int) []int {
	redCount := make(balls, 33)
	for i := 0; i < 33; i++ {
		redCount[i].id = i + 1
	}
	blueCount := make(balls, 16)
	for i := 0; i < 16; i++ {
		blueCount[i].id = i + 1
	}

	for i := len(units) - 1; i >= len(units)-num; i-- {
		u := units[i]
		for i := 0; i < 6; i++ {
			redCount[u.nums[i]-1].count++
		}
		blueCount[u.nums[6]-1].count++
	}
	sort.Sort(redCount)
	sort.Sort(blueCount)

	retRed := make([]int, 6)
	for i, v := range redCount[33-6:] {
		retRed[i] = v.id
	}
	sort.Ints(retRed)
	return append(retRed, blueCount[16-1].id)
}

func GetLotteryData(num int) ([]Unit, error) {
	if num <= 0 {
		return nil, errors.New("invalid num")
	}
	url := "http://datachart.500.com/ssq/zoushi/newinc/jbzs_redblue.php?expect=" + strconv.Itoa(num)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`<td align="center">[0-9]+.*<\/td>[\w\W]*?chartBall02[\w\W]*?<\/tr>`)
	rs := reg.FindAllString(string(body), -1)

	idStr := `<td align="center">`
	redStr := `<td class="chartBall01">`
	blueStr := `<td class="chartBall02">`
	endStr := `</td>`
	ret := make([]Unit, len(rs))
	for i, s := range rs {
		ret[i] = Unit{0, make([]int, 7)}
		i1 := strings.Index(s, idStr)
		i2 := strings.Index(s[i1:], endStr)
		id := strings.TrimSpace(s[i1+len(idStr) : i1+i2])
		ret[i].id, _ = strconv.Atoi(id)
		s = s[i1+i2:]

		for j := 0; j < 6; j++ {
			i1 := strings.Index(s, redStr)
			i2 := strings.Index(s[i1:], endStr)
			red := strings.TrimSpace(s[i1+len(redStr) : i1+i2])
			ret[i].nums[j], _ = strconv.Atoi(red)
			s = s[i1+i2:]
		}

		i1 = strings.Index(s, blueStr)
		i2 = strings.Index(s[i1:], endStr)
		blue := strings.TrimSpace(s[i1+len(blueStr) : i1+i2])
		ret[i].nums[6], _ = strconv.Atoi(blue)
	}
	return ret, nil
}

func main() {
	data, err := GetLotteryData(30)
	if err != nil {
		log.Fatal(err)
	}
	ret := PredictByMode(data, 10)
	fmt.Println(ret)
}
