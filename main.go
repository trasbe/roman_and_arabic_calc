package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rome_numbers_key = map[string]int{ //Ключи представляют тип string, значения - тип int
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
}
var rome_numbers_value = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	11:  "XI",
	12:  "XII",
	13:  "XIII",
	14:  "XIV",
	15:  "XV",
	16:  "XVI",
	17:  "XVII",
	18:  "XVIII",
	19:  "XIX",
	20:  "XX",
	21:  "XXI",
	22:  "XXII",
	23:  "XXIII",
	24:  "XXIV",
	25:  "XXV",
	26:  "XXVI",
	27:  "XXVII",
	28:  "XXVIII",
	29:  "XXIX",
	30:  "XXX",
	31:  "XXXI",
	32:  "XXXII",
	33:  "XXXIII",
	34:  "XXXIV",
	35:  "XXXV",
	36:  "XXXVI",
	37:  "XXXVII",
	38:  "XXXVIII",
	39:  "XXXIX",
	40:  "XL",
	41:  "XLI",
	42:  "XLII",
	43:  "XLIII",
	44:  "XLIV",
	45:  "XLV",
	46:  "XLVI",
	47:  "XLVII",
	48:  "XLVIII",
	49:  "XLIX",
	50:  "L",
	51:  "LI",
	52:  "LII",
	53:  "LIII",
	54:  "LIV",
	55:  "LV",
	56:  "LVI",
	57:  "LVII",
	58:  "LVIII",
	59:  "LIX",
	60:  "LX",
	61:  "LXI",
	62:  "LXII",
	63:  "LXIII",
	64:  "LXIV",
	65:  "LXV",
	66:  "LXVI",
	67:  "LXVII",
	68:  "LXVIII",
	69:  "LXIX",
	70:  "LXX",
	71:  "LXXI",
	72:  "LXXII",
	73:  "LXXIII",
	74:  "LXXIV",
	75:  "LXXV",
	76:  "LXXVI",
	77:  "LXXVII",
	78:  "LXXVIII",
	79:  "LXXIX",
	80:  "LXXX",
	81:  "LXXXI",
	82:  "LXXXII",
	83:  "LXXXIII",
	84:  "LXXXIV",
	85:  "LXXXV",
	86:  "LXXXVI",
	87:  "LXXXVII",
	88:  "LXXXVIII",
	89:  "LXXXIX",
	90:  "XC",
	91:  "XCI",
	92:  "XCII",
	93:  "XCIII",
	94:  "XCIV",
	95:  "XCV",
	96:  "XCVI",
	97:  "XCVII",
	98:  "XCVIII",
	99:  "XCIX",
	100: "C",
}

func add(x float32, y float32) float32 {
	var res float32 = x + y
	return res
}

func substraction(x float32, y float32) float32 {
	var res float32 = x - y
	return res
}

func divison(x float32, y float32) float32 {
	var res float32 = x / y
	return res
}
func multiplication(x float32, y float32) float32 {
	var res float32 = x * y
	return res
}
func run_operate(x float32, y float32, z string) (float32, error) {
	var res float32
	var err error
	if z == "+" {
		res = add(x, y)
	} else if z == "-" {
		res = substraction(x, y)
	} else if z == "/" {
		res = divison(x, y)
	} else if z == "*" {
		res = multiplication(x, y)
	} else {
		return 0, err
	}
	return res, nil
}

func check_romanian_symbols(x string, y string) bool {
	var matched, matched2 bool
	matched, _ = regexp.MatchString(`^(I|II|III|IV|V|VI|VII|VIII|IX|X)$`, x)
	matched2, _ = regexp.MatchString(`^(I|II|III|IV|V|VI|VII|VIII|IX|X)$`, y)
	if matched && matched2 {
		//println("check_romanian_symbols_approved")
		return true
	}
	return false
}

func check_arabic_numbers(x float32, y float32) bool {

	if (x >= 1 && x <= 10) && (y >= 1 && y <= 10) {
		//println("check_arabic_range_approved") //cool
		return true
	}
	return false
}

func check_arabic_numbers_regexp(x string, y string) bool {
	var matched, matched2 bool
	matched, _ = regexp.MatchString(`^(1|2|3|4|5|6|7|8|9|10)$`, x)
	matched2, _ = regexp.MatchString(`^(1|2|3|4|5|6|7|8|9|10)$`, y)
	if matched && matched2 {
		//println("check_arabic_numbers_approved")
		return true
	}
	return false
}

func check_operation(x string) bool {
	if x == "*" || x == "/" || x == "+" || x == "-" {
		return true
	}
	return false
}

func main() {
	var err error
	var result float32
	var result_roman string
	var string1, string2, operation string
	fmt.Println("Введите выражение арабскими цифрами или римскими от 1 до 10")
	if err != nil {
		log.Fatal(err)
	}
	r1 := bufio.NewReader(os.Stdin)
	input, err := r1.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(input)
	operation = words[1]
	string1 = words[0]
	string2 = words[2]
	if len(words) > 3 {
		log.Fatal("Введите выражение в формате \"число операция число\" ")
	}
	/*fmt.Println("Введите выражение арабскими цифрами или римскими от 1 до 10")
	_, err = fmt.Scan(&string3)
	if err != nil {
		log.Fatal(err)
	}*/
	//i_scan_numbered, err = fmt.Scan(&string1, &operation, &string2)
	//_, err = fmt.Scan(&string1, &operation, &string2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	if string1 == "" || string2 == "" || operation == "" {
		log.Fatal("Введены не все данные")
	}
	//println("operation = ;", operation)
	//println("check_operation = ;", check_operation(operation))
	if check_operation(operation) == false {
		log.Fatal("Введите правильный математический оператор")
	}
	//fmt.Println("Entered numbers:  ", string1, operation, string2, i_scan_numbered)
	int1, _ := strconv.Atoi(string1)
	/*if err != nil {
		log.Fatal(err)
	}*/
	int2, _ := strconv.Atoi(string2)
	/*if err != nil { here's err when rome numbers uses .Atoi
		log.Fatal(err)
	}*/
	num1 := float32(int1)
	num2 := float32(int2)
	/*if (check_romanian_symbols(string1, string2)) == false && check_arabic_numbers(num1, num2) == false {
		log.Fatalf("check romanian symbols and check arabic numbers not approved\n Entered expression: %s %s %s", string1, operation, string2)
	}*/
	if (check_romanian_symbols(string1, string2)) == false && check_arabic_numbers_regexp(string1, string2) == false {
		log.Fatalf("Проверьте правильность ввода данных")
	}
	if check_romanian_symbols(string1, string2) {
		int1 = rome_numbers_key[string1]
		int2 = rome_numbers_key[string2]
		//println("int1 = ", &int1, "; int2 = ", &int2)
		result, err = run_operate(float32(int1), float32(int2), operation)
		if err != nil {
			log.Fatal(err)
		}
		//result transfer to romanian symbols
		int1 = int(result) //result in int type we need to change it to romanian
		result_roman = rome_numbers_value[int1]
		if result < 0 { //result can't be <0
			log.Fatal("Римские числа не могут быть отрицательными")
		}
		fmt.Printf("Результат :\n%s\n", result_roman)
	} else if check_arabic_numbers(num1, num2) {
		//println("num1 = ", &num1, "; num2 = ", &num2)
		result, err = run_operate(num1, num2, operation)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Результат :\n%.f\n", result)
	}
}
