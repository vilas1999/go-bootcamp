package numberfiltering

import (
	"reflect"
	"testing"
)

// TestEvenNumbers tests the EvenNumbers function
func TestEvenNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}
	got := EvenNumbers(numbers)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`EvenNumbers(%v) = %v, want %v`, numbers, got, want)
	}
}

// TestEvenNumbers tests the EvenNumbers function
func TestOddNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 3, 5, 7, 9}
	got := OddNumbers(numbers)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`OddNumbers(%v) = %v, want %v`, numbers, got, want)
	}
}

func TestPrimeNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 7}
	got := getPrimeNumbers(numbers)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`PrimeNumbers(%v) = %v, want %v`, numbers, got, want)
	}
}

func TestOddPrimeNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{3, 5, 7}
	got := getOddPrimeNumbers(numbers)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`OddPrimeNumbers(%v) = %v, want %v`, numbers, got, want)
	}
}
