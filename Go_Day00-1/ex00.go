package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type resNumbers struct {
	median float64
	mean   float64
	mode   int64
	sd     float64
}

func CalcMedian(arrayNum []int64) float64 {
	sort.Slice(arrayNum, func(i, j int) bool { return arrayNum[i] < arrayNum[j] })
	mNumber := len(arrayNum) / 2
	if !(len(arrayNum)%2 == 0) {
		return float64(arrayNum[mNumber])
	} else {
		return (float64(arrayNum[mNumber-1]) + float64(arrayNum[mNumber])) / 2
	}
}
func CalcMean(arrayNum []int64) float64 {
	var total int64
	for _, v := range arrayNum {
		total += v
	}
	return float64(total) / float64(len(arrayNum))
}
func CalcMode(arrayNum []int64) int64 {
	var mode int64
	countMap := make(map[int64]int)
	for _, value := range arrayNum {
		countMap[value]++
	}
	max := 0
	for _, key := range arrayNum {
		freq := countMap[key]
		if freq > max {
			mode = key
			max = freq
		}
	}
	return mode
}

func CalcSd(arrayNum []int64) float64 {
	var total int64
	for _, v := range arrayNum {
		total += v
	}
	average := float64(total) / float64(len(arrayNum))
	toSqr := 0.0
	for _, v := range arrayNum {
		toSqr = toSqr + math.Pow((float64(v)-average), 2)
	}
	return math.Sqrt((toSqr) / float64(len(arrayNum)))
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	ary := []int64{}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVERED: %v\n", err)
		}
	}()
	for in.Scan() {
		line := in.Text()
		if line == "" {
			os.Exit(1)
		}
		x, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		if x > 100000 || x < (-100000) {
			os.Exit(3)
		}
		ary = append(ary, x)
	}
	var Mean, Median, Mode, SD bool
	flag.BoolVar(&Median, "Median", false, "tis is median")
	flag.BoolVar(&Mean, "Mean", false, "tis is mean")
	flag.BoolVar(&Mode, "Mode", false, "tis is mode")
	flag.BoolVar(&SD, "SD", false, "tis is sd")
	flag.Parse()
	res := resNumbers{}
	res.mean = CalcMean(ary)
	res.median = CalcMedian(ary)
	res.mode = CalcMode(ary)
	res.sd = CalcSd(ary)

	if Mean {
		fmt.Printf("Mean: %.2f\n", res.mean)
	}
	if Median {
		fmt.Printf("Median: %.2f\n", res.median)
	}
	if Mode {
		fmt.Printf("Mode: %d\n", res.mode)
	}
	if SD {
		fmt.Printf("SD: %.2f\n", res.sd)
	}
	if !Mean && !Mode && !Median && !SD {
		fmt.Printf("Mean: %.2f\nMedian: %.2f\nMode: %d\nSD: %.2f\n", res.mean, res.median, res.mode, res.sd)
	}
}
