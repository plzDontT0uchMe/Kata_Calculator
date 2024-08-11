package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dict = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}

func getRoman(value int) string {
	for k, v := range dict {
		if v == value {
			return k
		}
	}
	return ""
}

func toRoman(value int) string {
	var (
		zn1 string
		zn2 string
	)
	if value >= 1 && value <= 19 || value >= 40 && value <= 49 || value >= 90 && value <= 100 {
		zn1 = getRoman((value / 10) * 10)
		zn2 = getRoman(value % 10)
		return zn1 + zn2
	} else if value >= 20 && value <= 39 {
		count := value / 10
		for i := 0; i < count; i++ {
			zn1 += getRoman(10)
		}
		zn2 = getRoman(value % 10)
		return zn1 + zn2
	} else if value >= 50 && value <= 89 {
		zn1 = getRoman(50)
		zn2 = toRoman(value - 50)
		return zn1 + zn2
	}
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	splitText := strings.Split(text, " ")
	if len(splitText) != 3 {
		panic("строка не является математической операцией")
	}

	v1, err1 := strconv.Atoi(splitText[0])
	v2, err2 := strconv.Atoi(splitText[2])

	isRoman := false

	if err1 != nil && err2 != nil {
		isRoman = true
		v1 = dict[splitText[0]]
		v2 = dict[splitText[2]]
	}

	if err1 == nil && err2 != nil || err1 != nil && err2 == nil {
		panic("используются одновременно разные системы счисления")
	}

	if v1 > 10 || v2 > 10 || v1 <= 0 || v2 <= 0 {
		panic("неверный формат числа")
	}

	switch splitText[1] {
	case "+":
		v := v1 + v2
		if isRoman {
			fmt.Println(toRoman(v))
		} else {
			fmt.Println(v)
		}
	case "-":
		v := v1 - v2
		if isRoman {
			if v == 0 {
				panic("в римской системе нет 0")
			}
			if v < 0 {
				panic("в римской системе нет отрицательных значений")
			}
			fmt.Println(toRoman(v))
		} else {
			fmt.Println(v)
		}
	case "*":
		v := v1 * v2
		if isRoman {
			fmt.Println(toRoman(v))
		} else {
			fmt.Println(v)
		}
	case "/":
		v := v1 / v2
		if isRoman {
			if v == 0 {
				panic("в римской системе нет 0")
			}
			fmt.Println(toRoman(v))
		} else {
			fmt.Println(v)
		}
	default:
		panic("неизвестная операция")
	}
}
