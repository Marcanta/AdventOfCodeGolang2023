package main

import "testing"

type TestArgs struct {
	time, distance int
}

func TestCalculateDistance(t *testing.T) {
	totalTime := 7
	holdingTimes := map[int]int{1: 6, 2: 10, 3: 12, 4: 12, 5: 10, 6: 6, 7: 0}
	for k, v := range holdingTimes {
		if calculateValue := calculateDistance(totalTime, k); v != calculateValue {
			t.Errorf("Output %v not equal to expected %v", calculateValue, v)
		}
	}
}

func TestFindNumberOfWinningWay(t *testing.T) {
	testArgs := map[TestArgs]int{
		{time: 7, distance: 9}:    4,
		{time: 15, distance: 40}:  8,
		{time: 30, distance: 200}: 9,
	}

	for k, v := range testArgs {
		if numberOfWay := findNumberOfWinningWay(k.time, k.distance); numberOfWay != v {
			t.Errorf("Output %v not equal to expected %v", numberOfWay, v)
		}
	}
}
