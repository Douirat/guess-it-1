package main

import (
	"fmt"
	stat "logic/Logic"
	"math"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("File name not found!")
		return
	} else if args[0] != "data.txt" {
		fmt.Println("Wrong file name!")
		return
	}
	statistic := stat.NewStatistic(args[0])
	fmt.Println("Average:", int(math.Round(statistic.Average)))
	fmt.Println("Median:", int(math.Round(statistic.Median)))
	fmt.Println("Variance:", int(math.Round(statistic.Variance)))
	fmt.Println("Standard Deviation:", int(math.Round(statistic.StandardDeviation)))
	
}