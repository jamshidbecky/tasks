package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// ********************************************************

// func main() {
// 	var number int
// 	var qoldiq, revNumber string
// 	fmt.Scanln(&number)

// 	for number > 0 {
// 		count := 0
// 		qoldiq1 := number % 10
// 		number = number / 10
// 		qoldiq = strconv.Itoa(qoldiq1)
// 		revNumber += qoldiq
// 		count += 1
// 	}

// 	revNumber1, _ := strconv.Atoi(revNumber)

// 	fmt.Println(revNumber1)
	
// }

// ********************************************************

// func main() {
// 	var str1, str2 string

// 	fmt.Scanln(&str1)
// 	fmt.Scanln(&str2)

// 	res := str1[:len(str1) / 2] + str2 + str1[:len(str1) / 2]

// 	fmt.Println(res)

// }

// ********************************************************

// func main()  {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Ma'lumotni kiriting: ")
// 	data, _ := reader.ReadString('\n')
// 	data = strings.TrimSpace(data)
// 	dataLength := len(data)

// 	reversed := reverseString(data)

// 	if dataLength % 4 == 0 {
// 		data = data[:dataLength / 2] + reversed[:dataLength / 2]
// 		fmt.Println(data)
// 	}
// }

// func reverseString(input string) string {
// 	runes := []rune(input)
// 	length := len(runes)
// 	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// ********************************************************

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Ma'lumotni kiriting:")
// 	data, _ := reader.ReadString('\n')
// 	data = strings.TrimSpace(data)

// 	runeDate := []rune(data)
	
// 	for i := 0; i < len(runeDate); i++ {
// 		if i == 0 || i == len(runeDate) - 1 {
// 			runeDate[i] = runeDate[i] - 32
// 		}
// 	}

// 	fmt.Println(string(runeDate))
// }

// ********************************************************

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Ma'lumotni kiriting:")
// 	data, _ := reader.ReadString('\n')
// 	data = strings.TrimSpace(data)
// 	count := 0
// 	for _, val := range data {
// 		if val == 'a' || val == 'A' {
// 			count  += 1
// 		}
// 	}
// 	fmt.Println(count)
// }

// ********************************************************

// func main()  {
// 	var arr []string
// 	var word string

// 	inputReader := bufio.NewReader(os.Stdin)
// 	matn, _ := inputReader.ReadString('\n')
// 	matn = strings.TrimSpace(matn)

// 	for _, val := range matn {
// 		word += string(val)

// 		if string(val) == " " || string(val) == "." || string(val) == "!" || string(val) == "?" {
// 			arr = append(arr, word)
// 			word = ""
// 		}
// 	}
// 	fmt.Println(len(arr))
// }

// ********************************************************

// func main()  {
// 	inputReader := bufio.NewReader(os.Stdin)
// 	matn, _ := inputReader.ReadString('\n')
// 	matn = strings.TrimSpace(matn)

// 	runeMatn := []rune(matn)

// 	for i := 0; i < len(runeMatn); i++ {
// 		if i % 2 == 0 {
// 			if runeMatn[i] >= 97 && runeMatn[i] <= 122 {
// 				runeMatn[i] = runeMatn[i] - 32
// 			} 
// 		} else {
// 			if runeMatn[i] >= 65 && runeMatn[i] <= 90 {
// 				runeMatn[i] = runeMatn[i] + 32
// 			}
// 		} 
// 	}
	
// 	fmt.Println(string(runeMatn))
// }

// ********************************************************

// func main()  {
// 	inputReader := bufio.NewReader(os.Stdin)
// 	matn, _ := inputReader.ReadString('\n')
// 	matn = strings.TrimSpace(matn)
	
// 	var num int
// 	fmt.Scanln(&num)

// 	res := matn[:num] + matn[(num + 1):]

// 	fmt.Println(res)
// }

// ********************************************************

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	matn, _ := inputReader.ReadString('\n')
	matn = strings.TrimSpace(matn)

	runeMatn := []rune(matn)
	fmt.Println(runeMatn)

	for i := 0; i < len(runeMatn); i++ {
		if runeMatn[i] >= 48 && runeMatn[i] <= 57 {
			fmt.Printf("%v-indexda raqam bor\n", i)
		} else {
			fmt.Println("textda raqam qatnashmagan")
		}
	} 
}