package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var romanMap = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var expressionSlice []string

func main() {
	var input string
	fmt.Println("Калькулятор принимает римские и арабские числа от 1 до 10")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	expression := strings.ReplaceAll(input, " ", "")
	expression = strings.ToUpper(expression)
	calculation(expression)
}

func calculation(expression string) {
	var operator string
	var stringCount int
	arabicNumbers := make([]int, 0)
	romanNumbers := make([]string, 0)
	romansToInt := make([]int, 0)
	for key := range operators {
		for _, value := range expression {
			if key == string(value) {
				operator += key
			}
		}
	}
	switch {
	case utf8.RuneCountInString(operator) > 1:
		panic("Может быть только 1 оператор")
	case utf8.RuneCountInString(operator) < 1:
		panic("Выражение должно содержать один из операторов: +, -, *, / ")
	}
	expressionSlice = strings.Split(expression, operator)
	for _, el := range expressionSlice {
		number, err := strconv.Atoi(el)
		if err == nil {
			arabicNumbers = append(arabicNumbers, number)
		} else {
			stringCount++
			romanNumbers = append(romanNumbers, el)
		}
	}
	switch {
	case stringCount == 1:
		panic("Арабские и римские цифры смешивать нельзя")
	case stringCount == 0:
		// Проверить сколько чисел. Проверить их значение от 1 до 10. Если все ок, то задать значение указателей и завершить
		if arabicNumbers[0] <= 10 && arabicNumbers[0] > 0 && arabicNumbers[1] > 0 && arabicNumbers[1] <= 10 {
			a = &arabicNumbers[0]
			b = &arabicNumbers[1]
			val, _ := operators[operator]
			fmt.Println(val())
			fmt.Println(romansToInt)
		} else {
			panic("Введите число от 1 до 10")
		}

	case stringCount == 2:
		for _, elem := range romanNumbers {
			if val, ok := romanMap[elem]; ok && val >= 1 && val <= 10 {
				romansToInt = append(romansToInt, val)
			} else {
				panic("Римские цифры введены некорректно")
			}
		}
		a = &romansToInt[0]
		b = &romansToInt[1]
		val, _ := operators[operator]
		number := val()
		if number > 0 {
			romanSolution := integerToRoman(number)
			fmt.Println(romanSolution)
		} else {
			fmt.Println("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
		}
	case stringCount > 2:
		panic("Как Вы это сделали?")
	}
}

func integerToRoman(number int) string {
	romanConv := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range romanConv {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
