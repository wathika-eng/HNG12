package main

import (
	"fmt"
	"io"
	"net/http"
)

// parse the number and check properties as defined in the Number struct
func parseNum(num *numClass) {
	if isPrime(num.Number) {
		num.IsPrime = true
	}
	if isPerfect(num.Number) {
		num.IsPerfect = true
	}
	if isArmstrong(num.Number) {
		num.Properties = append(num.Properties, "armstrong")
	}
	if isOdd(num.Number) {
		num.Properties = append(num.Properties, "odd")
	} else {
		num.Properties = append(num.Properties, "even")
	}
	num.DigitSum = sumDigits(num.Number)
	apiResp, err := funFact(num.Number)
	if err != nil {
		num.FunFact = string(err.Error())
	}
	num.FunFact = apiResp
}

// get a random fun fact from an external API
func funFact(n int) (string, error) {
	// format the url to include the number as a query
	url := fmt.Sprintf("http://numbersapi.com/%v", n)
	resp, err := http.Get(url)
	// return an error if
	if err != nil || resp.StatusCode != 200 {
		return "", fmt.Errorf("error while fetching fun fact from API: %v", err)
	}
	// close the response body to save on resources
	defer resp.Body.Close()
	// will return byte data response
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading response: %v", err)
	}
	return string(response), nil
}
