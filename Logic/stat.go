package logic

import (
	// "fmt"
	"os"
	"sort"
	"math"
)

// Classify my program as an object population:
type Statistic struct {
Mined [] byte
Data []float64
Average float64
Median float64
Variance float64
StandardDeviation float64
}

// Instantiate a new Population object:
func NewStatistic(fileName string) *Statistic {
stat := new(Statistic)
data, err := os.ReadFile(fileName)
if err != nil {
	panic(err)
}
stat.Mined = data
stat.GetData()
stat.Average = stat.CalcAverage(stat.Data)
stat.GetMedian()
stat.GetVariance()
stat.StandardDeviation = math.Sqrt(stat.Variance)
return stat
}

// Convert the data from bytes to floats:
func (stat *Statistic) GetData() {
data := []float64{}
x := 0.0
for _, digit := range stat.Mined {
	if digit == 10 {
		data = append(data, x)
		x = - 1.0
	}
	if digit >= 48 && digit <= 57 {
		if x == -1.0 {
			x = 0.0
		}
		x = x * 10 +  float64(digit-48)
	} else {
		continue
	}
}
if x >= 0.0 {
	data = append(data, x)
}
stat.Data = data
}

// calculate the average:
func (stat *Statistic) CalcAverage(digs []float64) float64 {
	x := 0.0
	for _, v := range digs {
		x += v
	}
	return x / float64(len(stat.Data))
}

// Determine the median:
func (stat *Statistic) GetMedian() {
	sort.Float64s(stat.Data)
	if len(stat.Data) % 2 == 0 {
		stat.Median = (stat.Data[(len(stat.Data)/2)] + stat.Data[(len(stat.Data)/2) - 1]) / 2
	} else {
		stat.Median = stat.Data[(len(stat.Data)/2)]
	}
}

// Calculate the variance of these numbers:
func (stat *Statistic) GetVariance() {
	// Keep track of the squered diffrences:
	SqrDif := []float64{}
	for _, dig := range stat.Data {
		sq := dig - stat.Average
		sq = sq * sq
		SqrDif = append(SqrDif, (sq))
	}
	stat.Variance = stat.CalcAverage(SqrDif)
}

