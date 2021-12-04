package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func lenAndUpper1(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	defer lenAndUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total = total + number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink1(age int) bool {
	koreanAge := age + 2
	switch {
	case koreanAge < 18:
		return false
	case koreanAge > 18:
		return true
	default:
		return false
	}
}

func main() {
	totalLength, upperName := lenAndUpper("hyunseok")
	fmt.Println(totalLength, upperName)

	repeatMe("nico", "lynn", "dal", "mal")

	len1, up := lenAndUpper1("honggildong")
	println(len1, up)

	result := superAdd(1, 2, 3, 4, 5)
	println(result)

	canIDrinkResult := canIDrink(17)
	println(canIDrinkResult)

	canIDrinkResult1 := canIDrink1(17)
	println(canIDrinkResult1)

	var abcd []string
	abcd = []string{"asdf", "sdfg"}
	fmt.Println(len(abcd))
}
