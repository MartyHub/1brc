package main

import "fmt"

type station struct {
	city     string
	count    int
	sum      int
	min, max int
}

func (stn *station) add(val int) {
	stn.count++
	stn.sum += val
	stn.min = min(stn.min, val)
	stn.max = max(stn.max, val)
}

func (stn *station) avg() float64 {
	return float64(stn.sum) / float64(stn.count*10)
}

func (stn *station) merge(other *station) {
	stn.count += other.count
	stn.sum += other.sum
	stn.min = min(stn.min, other.min)
	stn.max = max(stn.max, other.max)
}

func (stn *station) String() string {
	return fmt.Sprintf("%s=%.1f/%.1f/%.1f", stn.city, round(stn.min), stn.avg(), round(stn.max))
}

func round(i int) float64 {
	return float64(i) / 10.0
}
