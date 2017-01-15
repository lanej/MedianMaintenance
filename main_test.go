package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseInitial(t *testing.T) {
	input := []int{}
	expected := float64(0)
	median := NewMedian(input)

	actual := median.Median()

	assert.Equal(t, expected, actual)
}

func TestCaseOne(t *testing.T) {
	input := []int{1}
	expected := float64(1)
	median := NewMedian(input)

	actual := median.Median()

	assert.Equal(t, expected, actual)
}

func TestCaseOneTwo(t *testing.T) {
	input := []int{1, 2}
	expected := float64(1)
	median := NewMedian(input)

	actual := median.Median()

	assert.Equal(t, expected, actual)
}

func TestCaseTwoOne(t *testing.T) {
	input := []int{2, 1}
	expected := float64(1)
	median := NewMedian(input)

	actual := median.Median()

	assert.Equal(t, expected, actual)
}

func TestCaseThreeTwoOne(t *testing.T) {
	input := []int{3, 2, 1}
	expected := float64(2)
	median := NewMedian(input)

	actual := median.Median()

	assert.Equal(t, expected, actual)
}

func TestCase1(t *testing.T) {
	input := []int{1, 3, 3, 6, 7, 8, 9}
	expected := float64(6)
	median := NewMedian(input)

	actual := median.Median()

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{6331, 2793, 1640, 9290, 225, 625, 6195, 2303, 5685, 1354}
	expected := float64(9335)
	actual := sumOver(input)

	assert.Equal(t, expected, actual)
}

func TestHomework(t *testing.T) {
	input := load("hw.txt")
	expected := float64(1213)
	actual := sumOver(input)

	assert.Equal(t, expected, actual)
}

func sumOver(input []int) float64 {
	var medianSum float64
	median := NewMedian([]int{})
	for _, j := range input {
		median.Add(j)
		// fmt.Printf("%v: Median is %v\n", j, median.Median())
		medianSum += median.Median()
	}

	// fmt.Printf("True sum: %v\n", medianSum)
	actual := math.Mod(medianSum, float64(10000))

	return actual
}
