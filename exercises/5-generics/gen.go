package main

import "fmt"

type Number interface {
  int64 | float64
}

func main() {

	// Init
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"firts":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
    // Non-generic
    SumInts(ints),
    SumFloats(floats),
	)

	fmt.Printf("Generic Sums: %v and %v\n",
    // Generic
		SimIntsOrFloats[string, int64](ints),
		SimIntsOrFloats[string, float64](floats),
	)
  
  fmt.Printf("Generic Sums2: %v and %v\n",
    // Generic
		SimIntsOrFloats(ints),
		SimIntsOrFloats(floats),
	)

  fmt.Printf("Generic Sums3: %v and %v\n",
    // Generic
		SumNumbers(ints),
		SumNumbers(floats),
	)

}

// Non-genetic sums
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Generic sums
func SimIntsOrFloats[K comparable, V int64 | float64] (m map[K]V) V {
  var s V
  for _, v := range m{
    s += v
  }
  return s
}
// K especcify comparable param ( == or =!)
// V will be int64 or float64

// Using an interface
func SumNumbers[K comparable, V Number](m map[K]V) V {
  var s V
  for _, v := range m{
    s += v
  }
  return s
}

