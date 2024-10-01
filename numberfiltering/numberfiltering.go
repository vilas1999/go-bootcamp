package numberfiltering

// Go function to return even numbers
func EvenNumbers(numbers []int) []int {
	evenNumbers := []int{}
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return evenNumbers
}

func OddNumbers(numbers []int) []int {
	oddNumbers := []int{}
	for _, number := range numbers {
		if number%2 == 1 {
			oddNumbers = append(oddNumbers, number)
		}
	}
	return oddNumbers
}

func getPrimeNumbers(numbers []int) []int {
	primeNumbers := []int{}
	for _, number := range numbers {
		if isPrimeNumber(number) {
			primeNumbers = append(primeNumbers, number)
		}
	}
	return primeNumbers
}

func getOddPrimeNumbers(number []int) []int {
	primeNumbers := getPrimeNumbers(number)
	return OddNumbers(primeNumbers)
}

func isPrimeNumber(number int) bool {
	if number == 1 || number == 0 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
