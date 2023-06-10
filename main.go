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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ma'lumotni kiriting:")
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	count := 0
	for _, val := range data {
		if val == 'a' || val == 'A' {
			count  += 1
		}
	}
	fmt.Println(count)
}